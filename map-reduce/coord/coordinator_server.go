package main

import (
	"fmt"
	map_reduce "map-reduce"
)

func main() {
	coordinator := map_reduce.MakeCoordinator()
	if coordinator == nil {
		fmt.Println("create coordinator failed")
	}

	ch := make(chan struct{})

	<-ch
}
