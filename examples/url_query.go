package examples

import (
	"fmt"
	"io/ioutil"
	"log"
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
		log.Println(r.url, r.contentLen, r.err)
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
