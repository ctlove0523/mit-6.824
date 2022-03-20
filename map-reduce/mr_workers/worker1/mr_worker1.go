package main

import (
	"fmt"
	mr "map-reduce"
)

func main() {
	w := mr.NewWorker(1, "localhost:8600", "localhost:8500")
	if w == nil {
		fmt.Println("start worker failed")
	}

	ch := make(chan struct{})
	<-ch
}
