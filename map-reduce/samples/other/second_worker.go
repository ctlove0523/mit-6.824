package main

import (
	mr "mit-6.824/map-reduce"
)

func main() {
	w := mr.Worker{
		Address:            "localhost:4791",
		CoordinatorAddress: "localhost:4789",
	}

	w.Start()
}
