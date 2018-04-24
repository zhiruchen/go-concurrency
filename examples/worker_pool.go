package examples

import (
	"fmt"
	"sync"
)

type task struct {
	err error
	f   func() error
}

func newTask(f func() error) *task {
	return &task{f: f}
}

func (t *task) run(wg *sync.WaitGroup) {
	t.err = t.f()
	wg.Done()
}

type WorkerPool struct {
	tasks       []*task
	concurrency int
	taskChan    chan *task
	wg          sync.WaitGroup
}

func NewPool(tasks []*task, concurrency int) *WorkerPool {
	return &WorkerPool{
		tasks:       tasks,
		concurrency: concurrency,
		taskChan:    make(chan *task),
	}
}

func (p *WorkerPool) work() {
	for task := range p.taskChan {
		task.run(&p.wg)
	}
}

func (p *WorkerPool) run() {
	for i := 0; i < len(p.tasks); i++ {
		go p.work()
	}

	p.wg.Add(len(p.tasks))
	for _, task := range p.tasks {
		p.taskChan <- task
	}

	close(p.taskChan)
	p.wg.Wait()
}

func RunTasks() {
	tasks := []*task{
		{f: func() error { fmt.Println("run task1"); return nil }},
		{f: func() error { fmt.Println("run task2"); return nil }},
		{f: func() error { fmt.Println("run task3"); return nil }},
		{f: func() error { fmt.Println("run task4"); return nil }},
		{f: func() error { fmt.Println("run task5"); return nil }},
	}

	p := NewPool(tasks, 10)
	p.run()

	for _, t := range p.tasks {
		if t.err != nil {
			fmt.Println(t.err)
		}
	}
}
