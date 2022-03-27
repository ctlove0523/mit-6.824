package main

import (
	mr "mit-6.824/map-reduce"
)

func main() {
	c := mr.NewCoordinator("localhost:4789", "localhost:3789")

	c.Start()
}
