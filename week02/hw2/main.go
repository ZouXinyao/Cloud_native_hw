package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
		fmt.Printf("producer %s \n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer %s \n", data)
	}
	wg.Done()
}

func main() {

	chanSteam := make(chan string, 10)

	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanSteam)
	}

	for i := 0; i < 2; i++ {
		wgCs.Add(1)
		go consumer(wgCs, chanSteam)
	}

	wgPd.Wait()
	close(chanSteam)
	wgCs.Wait()
}
