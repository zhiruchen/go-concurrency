package examples

import (
	"log"
)

/*Merge2Chan 合并俩channel
关闭nil channel 会panic
发送数据到nil channel 会一直阻塞
从nil channel 接受数据会也会一直阻塞
*/
func Merge2Chan() {
	a := asChan(1, 2, 3, 4)
	b := asChan(5, 6, 7, 8)

	for v := range merge2Chan(a, b) {
		log.Printf("received %d from c\n", v)
	}
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}

func merge2Chan(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
	Loop:
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					log.Println("a is done")
					a = nil
					continue Loop
				}

			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					log.Println("b is done")
					b = nil
					continue Loop
				}
			}
		}

		log.Println("closing c")
		close(c)
	}()

	return c
}
