package examples

import "testing"

func TestQueryURLs(t *testing.T) {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://zhiruchen.github.io",
		"https://github.com/zhiruchen",
		"https://zbg.herokuapp.com",
		"https://baidu.com",
		"https://stackoverflow.com/",
	}

	QueryURLs(urls)
}

func TestSimulateWaitGroup(t *testing.T) {
	simulateWaitGroup()
}
