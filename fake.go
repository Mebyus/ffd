package main

import (
	"time"

	"github.com/mebyus/ffd/cli"
)

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

func fake(command *cli.Command) error {
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
