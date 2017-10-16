package main

import (
	"github.com/zhiruchen/go-concurrency/lock"
	"github.com/zhiruchen/go-concurrency/sync"
)

func main() {
	// sync.Main()
	sync.ShareData()
	lock.Main()
	lock.Race()
	lock.RightRace()
}
