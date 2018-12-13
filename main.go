package main

import (
	"github.com/pkwenda/asJob/download"
	"sync"
)

var wg sync.WaitGroup

var count int
func main() {
	download.Worker(2)
	//for i:=0; i<5005; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		count++
	//		fmt.Println(count)
	//	}(i)
	//
	//}
	//
	//wg.Wait()
}

