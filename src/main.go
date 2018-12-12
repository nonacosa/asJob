package main

import (
	"fmt"
	"github.com/pkwenda/asJob/src/download"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i:=0; i<4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err:= download.Spider(); err != nil {
				fmt.Print(err)
			}
		}(i)

	}
	wg.Wait()

}
