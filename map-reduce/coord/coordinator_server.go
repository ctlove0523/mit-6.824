package main

import (
	"fmt"
	mr "map-reduce"
)

func main() {
	coordinator := mr.MakeCoordinator()
	if coordinator == nil {
		fmt.Println("create coordinator failed")
	}

	ch := make(chan struct{})

	<-ch
}
