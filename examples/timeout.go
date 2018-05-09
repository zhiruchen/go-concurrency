package examples

import (
	"log"
	"sync"
	"time"

	req "github.com/parnurzeal/gorequest"
)

type reqResult struct {
	url  string
	resp req.Response
	body string
	errs []error
}

func visitURL(url string, resChan chan<- *reqResult) {
	resp, body, errs := req.New().Get(url).Timeout(4 * time.Second).End()
	resChan <- &reqResult{url: url, resp: resp, body: body, errs: errs}
}

func ForSelectVisitURL(urls []string) {
	resChans := getResults(urls)
Loop:
	for {
		select {
		case res, ok := <-resChans:
			if !ok {
				break Loop
			}
			log.Printf("get the result of %s, body: %d\n", res.url, len(res.body))
		default:
		}
	}

	log.Println("end visit")
}

func getResults(urls []string) <-chan *reqResult {
	var wg sync.WaitGroup
	resultChans := make(chan *reqResult, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			visitURL(url, resultChans)
			wg.Done()
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultChans)
	}()

	return resultChans
}
