package main

import (
	"fmt"
	mr "map-reduce"
)

func main() {
	w := mr.NewWorker(2, "localhost:8601", "localhost:8500")
	if w == nil {
		fmt.Println("start worker failed")
	}

	ch := make(chan struct{})
	<-ch
}
