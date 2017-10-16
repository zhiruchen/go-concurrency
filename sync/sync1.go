package sync

import (
	"log"
	"time"
)

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})

	go func() {
		time.Sleep(delay)
		log.Println(text)
		close(ch)
	}()

	return ch
}

func Main() {
	wait := Publish("channel let goroutines communicate.", 5*time.Second)
	log.Printf("waiting for news\n")
	<-wait
	log.Printf("the news out, time to leave.\n")
}

func ShareData() {
	ch := make(chan int)

	go func() {
		n := 0
		n++
		ch <- n
	}()

	n := <-ch
	n++
	log.Printf("n:  %d\n", n)
}
