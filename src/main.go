package main

import (
	"fmt"
	"github.com/pkwenda/asJob/src/download"
)

func main() {
	for i:=0; i<70; i++ {
		if err := download.Go(); err != nil {
			fmt.Println(err)
		}
	}




}
