package examples

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

type qResult struct {
	url        string
	contentLen int
	err        error
}

func QueryURLs(us []string) {
	runtime.GOMAXPROCS(8)
	rs := make(chan *qResult, len(us))

	for _, url := range us {
		go func(u string) {
			doQuery(u, rs)
		}(url)
	}

	for i := 1; i <= len(us); i++ {
		r := <-rs
		if r.err != nil {
			log.Println(r.url, r.err)
		} else {
			log.Println(r.url, r.contentLen)
		}
	}
	close(rs)
}

func doQuery(u string, rs chan<- *qResult) {
	c := make(chan *qResult)
	go queryURL(u, c)
	select {
	case v := <-c:
		rs <- v
	case <-time.After(3 * time.Second):
		rs <- &qResult{url: u, err: fmt.Errorf("req: %s timeout", u)}
	}
	close(c)
}

func queryURL(url string, c chan *qResult) {
	var r = &qResult{url: url}

	resp, err := http.Get(url)
	if err != nil {
		r.err = err
		c <- r
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.err = err
		c <- r
		return
	}

	r.contentLen = len(body)
	c <- r
	return
}

func simulateWaitGroup() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(8)
	rs := make(chan int, 9)

	for i := 1; i <= 9; i++ {
		go func(i int) {
			ch := make(chan int)
			var d time.Duration
			if i%2 == 0 {
				d = 3 * time.Second
			} else {
				d = time.Second
			}
			go timeConsumingFunc(ch, d)
			select {
			case v := <-ch:
				rs <- v
			case <-time.After(2 * time.Second):
				fmt.Println(i, "timeout")
				rs <- -1
			}
		}(i)
	}

	for j := 1; j <= 9; j++ {
		fmt.Println(<-rs)
	}

	close(rs)
}

func timeConsumingFunc(ch chan int, d time.Duration) {
	time.Sleep(d)
	ch <- rand.Int()
}
