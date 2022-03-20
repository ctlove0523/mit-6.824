package main

import (
	"fmt"
	mr "map-reduce"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8500")
	if err != nil {
		fmt.Println("connect to coordinator failed")
		return
	}
	req := mr.AddTaskRequest{
		FileNames:  []string{"pg-grimm.txt", "pg-huckleberry_finn.txt"},
		ReduceSize: 2,
	}
	resp := mr.AddTaskResponse{}
	err = client.Call("Coordinator.AddMapReduceTask", &req, &resp)
	if err != nil {
		fmt.Printf("user add map reduce failed,err = %s\n", err)
		return
	}

	fmt.Println(resp.TaskId)

}
