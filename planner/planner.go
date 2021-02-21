package planner

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mebyus/ffd/cmn"
	"github.com/mebyus/ffd/setting"
)

type pool struct {
	max int
	cur int
}

type Result struct {
	Content io.ReadCloser
	Err     error
}

type Task struct {
	Label       string
	URL         string
	Result      *Result
	Destination chan<- *Result
}

var (
	Tasks  = make(chan *Task, 10)
	Client = &http.Client{
		Timeout: setting.ClientTimeout,
	}
)

func Planner() {
	pools := map[string]*pool{
		"SB": {
			max: 1,
			cur: 0,
		},
		"RR": {
			max: 5,
			cur: 0,
		},
	}
	pending := map[string]*[]*Task{}
	for key := range pools {
		pending[key] = &[]*Task{}
	}
	toGatherer := make(chan *Task, 10)
	notifications := make(chan string, 10)
	go gatherer(toGatherer, notifications)
	for {
		select {
		case task := <-Tasks:
			p, allowed := pools[task.Label]
			if allowed {
				if p.cur < p.max {
					go worker(task, toGatherer, Client)
					p.cur++
					fmt.Printf("%s: %d # %s\n", task.Label, p.cur, task.URL)
				} else {
					queue := pending[task.Label]
					*queue = append(*queue, task)
				}
			} else {
				task.Destination <- &Result{
					Err: fmt.Errorf("unknown worker pool label: %s", task.Label),
				}
			}
		case notification := <-notifications:
			p, allowed := pools[notification]
			if allowed {
				p.cur--
				queue := pending[notification]
				if len(*queue) > 0 {
					task := (*queue)[0]
					go worker(task, toGatherer, Client)
					p.cur++
					fmt.Printf("%s: %d # %s\n", task.Label, p.cur, task.URL)
					*queue = (*queue)[1:]
				}
			} else {
				fmt.Printf("unknown notification label: %s\n", notification)
			}
		}
	}
}

func worker(task *Task, toGatherer chan<- *Task, client *http.Client) {
	start := time.Now()
	body, err := cmn.GetBody(task.URL, client)
	fmt.Printf("%v spent for %s\n", time.Since(start), task.URL)
	task.Result = &Result{
		Content: body,
		Err:     err,
	}
	toGatherer <- task
}

func gatherer(tasks <-chan *Task, notifications chan<- string) {
	for task := range tasks {
		notifications <- task.Label
		task.Destination <- task.Result
	}
}
