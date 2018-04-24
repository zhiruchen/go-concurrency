package examples

import "fmt"

func doJobs() {
	jobs := make(chan int, 10)
	done := make(chan struct{})

	go func() {
		for {
			select {
				case job, more := <- jobs:
					if more {
						fmt.Println("received job: ", job)
					} else {
						done <- struct{}{}
						return
					}
			}
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println("sending job: ", i)
		jobs <- i
	}
	close(jobs)

	<-done
}
