package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mebyus/ffd/resource"
	"github.com/mebyus/ffd/track"
)

func unknown(args []string) error {
	return fmt.Errorf("unknown command")
}

func help(args []string) error {
	return nil
}

func download(args []string) (err error) {
	if len(args) == 0 {
		return nil
	} else if len(args) == 1 {
		err = resource.Download(args[0], false)
	} else {
		if args[0] == "-s" {
			err = resource.Download(args[1], true)
		} else {
			fmt.Printf("unknown flag \"%s\"\n", args[0])
			err = resource.Download(args[1], false)
		}
	}
	if err != nil {
		return
	}
	return
}

func add(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("\"add\" command: not enough arguments")
	}
	err := track.Add(args[0])
	if err != nil {
		return err
	}
	return nil
}

func check(args []string) error {
	err := track.Check()
	if err != nil {
		return err
	}
	return nil
}

func parse(args []string) (err error) {
	if len(args) < 2 {
		return fmt.Errorf("\"parse\" command: not enough arguments")
	}
	if len(args) == 2 {
		if !strings.HasPrefix(args[1], "--hostname=") {
			return fmt.Errorf("hostname not specified")
		}
		hostname := args[1][11:]
		if hostname == "" {
			return fmt.Errorf("hostname not specified")
		}
		err = resource.Parse(args[0], hostname, false)
	} else {
		separate := false
		if args[0] == "-s" {
			separate = true
		} else {
			fmt.Printf("unknown flag \"%s\"\n", args[0])
		}
		if !strings.HasPrefix(args[2], "--hostname=") {
			return fmt.Errorf("hostname not specified")
		}
		hostname := args[2][11:]
		if hostname == "" {
			return fmt.Errorf("hostname not specified")
		}
		err = resource.Parse(args[1], hostname, separate)
	}
	if err != nil {
		return err
	}
	return nil
}

func suppress(args []string) (err error) {
	if len(args) == 0 {
		return fmt.Errorf("\"suppress\" command: not enough arguments")
	} else if len(args) == 1 {
		err = track.Suppress(args[0], false)
	} else {
		if args[0] == "-r" {
			err = track.Suppress(args[0], true)
		} else {
			fmt.Printf("unknown flag \"%s\"\n", args[0])
		}
	}
	if err != nil {
		return
	}
	return
}

func list(args []string) (err error) {
	err = track.List()
	if err != nil {
		return err
	}
	return
}

func tasker(toDamper chan<- int) {

}

func damper(fromTasker <-chan int, toSpawner chan<- int) {

}

func gatherer(toDamper chan<- int, fromWorker <-chan int, toOrigin chan<- int) {

}

func spawner(fromDamper <-chan int, toGatherer chan<- int) {

}

func worker(toGatherer chan<- int) {
	time.Sleep(3 * time.Second)
}

func fake(args []string) error {
	poolSize := 10
	chTaskerDamper := make(chan int, 2)
	chDamperSpawner := make(chan int, 2)
	chGathererDamper := make(chan int, poolSize)
	chWorkerGatherer := make(chan int, poolSize)
	chFinish := make(chan int)
	go tasker(chTaskerDamper)
	go damper(chTaskerDamper, chDamperSpawner)
	go spawner(chDamperSpawner, chWorkerGatherer)
	go gatherer(chGathererDamper, chWorkerGatherer, chFinish)
	return nil
}

func choose(name string) (executor func(args []string) error) {
	switch name {
	case "download":
		executor = download
	case "parse":
		executor = parse
	case "fake":
		executor = fake
	case "help":
		executor = help
	case "add":
		executor = add
	case "check":
		executor = check
	case "suppress":
		executor = suppress
	case "list":
		executor = list
	default:
		executor = unknown
	}
	return
}

func main() {
	var (
		executor func(args []string) error
		args     []string
	)

	if len(os.Args) < 2 {
		args = nil
		executor = choose("help")
	} else if len(os.Args) == 2 {
		args = nil
		executor = choose(os.Args[1])
	} else {
		args = os.Args[2:]
		executor = choose(os.Args[1])
	}

	err := executor(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
