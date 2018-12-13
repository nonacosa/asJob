package main

import (
	"github.com/pkwenda/asJob/download"
	"sync"
)

var wg sync.WaitGroup

func main() {
	download.Worker(5)
}
