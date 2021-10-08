package done

import "fmt"

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, in chan int, done chan bool) {
	for val := range in {
		fmt.Printf("%d----%c\n", id, val)
		done <- true
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool, 2),
	}
	go doWork(id, w.in, w.done)
	return w
}

func ChanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}

}
