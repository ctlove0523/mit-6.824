package map_reduce

import (
	"fmt"
	"sync"
)
import "net"
import "net/rpc"
import "net/http"

// workerInfo 用于表示已经注册的worker
type workerInfo struct {
	Id      int    // 唯一表示一个执行任务的worker
	Address string // 执行任务的worker的地址，用于RPC调用
}

type Task struct {
	Id        string
	mapStatus int
	mapLock   sync.Mutex

	reduceStatus int
	reduceLock   sync.Mutex
}

type MapTask struct {
	TaskId   string
	WorkerId int
	Outputs  []MapTaskOutput
}

type MapTaskOutput struct {
	ReduceRegionNo int
	FileName       string
}

type ReduceTask struct {
	TaskId         string
	WorkerId       int
	ReduceRegionNo int
	Outputs        []ReduceTaskOutput
}

type ReduceTaskOutput struct {
	ReduceRegionNo int
	FileName       string
}

type Coordinator struct {
	Address     string
	Workers     []workerInfo
	Tasks       map[string]*Task
	MapTasks    map[string]map[int]*MapTask    // value为每个worker处理的map任务
	ReduceTasks map[string]map[int]*ReduceTask // value为每个worker处理的reduce任务
	ReduceSize  int
}

func (c *Coordinator) ReportMapTaskProgress(req *ReportMapTaskProgressRequest, resp *RpcResponse) error {
	fmt.Printf("worker %d begin report map task progress, map task in %d\n", req.WorkerId, req.TaskStatus)
	task := c.Tasks[req.TaskId]
	task.mapLock.Lock()
	task.mapStatus = task.mapStatus - req.TaskStatus
	mapTasks := c.MapTasks[req.TaskId]
	mapTask := mapTasks[req.WorkerId]
	mapTask.Outputs = append(mapTask.Outputs, req.Outputs...)

	task.mapLock.Unlock()

	if task.mapStatus == 0 {
		fmt.Println("map task finished,begin to execute reduce task")
		go func() {
			mapTaskMap := c.MapTasks[req.TaskId]
			var mapTaskOutputs []MapTaskOutput
			for _, v := range mapTaskMap {
				mapTaskOutputs = append(mapTaskOutputs, v.Outputs...)
			}

			var requests []AssignReduceTaskRequest
			for i := 0; i < c.ReduceSize; i++ {
				var names []string
				for _, out := range mapTaskOutputs {
					if out.ReduceRegionNo == i {
						names = append(names, out.FileName)
					}
				}
				request := AssignReduceTaskRequest{
					TaskId:       req.TaskId,
					FileNames:    names,
					ReduceNumber: i,
				}

				requests = append(requests, request)
			}

			for _, v := range requests {
				w := c.Workers[v.ReduceNumber%len(c.Workers)]
				c.AssignReduceTask(w, v)
			}
		}()
	}
	*resp = RpcResponse{}
	return nil
}

// RegisterWorker worker向Coordinator注册
func (c *Coordinator) RegisterWorker(req RegisterWorkerRequest, resp *RegisterWorkerResponse) error {
	fmt.Println("register workerInfo")
	w := workerInfo{
		Id:      req.Id,
		Address: req.Address,
	}
	c.Workers = append(c.Workers, w)

	*resp = RegisterWorkerResponse{
		Id:                 req.Id,
		CoordinatorAddress: c.Address,
	}
	return nil
}

func (c *Coordinator) AddMapReduceTask(req *AddTaskRequest, resp *AddTaskResponse) error {
	fmt.Println("begin map reduce task")
	t := &Task{
		Id:        NewTaskId(),
		mapStatus: len(c.Workers),
		mapLock:   sync.Mutex{},
	}
	c.Tasks[t.Id] = t
	fileNames := req.FileNames

	inputs := map[int][]string{}
	for i := 0; i < len(fileNames); i++ {
		index := i % len(c.Workers)
		files, ok := inputs[index]
		if !ok {
			files = []string{fileNames[i]}
		} else {
			files = append(files, fileNames[i])
		}
		inputs[index] = files
	}

	for k, v := range inputs {
		mapRequest := AssignMapTaskRequest{
			TaskId:     t.Id,
			MapTaskId:  k,
			Files:      v,
			ReduceSize: c.ReduceSize,
		}
		mapTask := &MapTask{
			TaskId:   t.Id,
			WorkerId: c.Workers[k].Id,
			Outputs:  []MapTaskOutput{},
		}
		fmt.Printf("map task； worker id = %d\n", mapTask.WorkerId)
		workerMapTasks, ok := c.MapTasks[t.Id]
		if !ok {
			workerMapTasks = make(map[int]*MapTask)
		}
		workerMapTasks[c.Workers[k].Id] = mapTask
		c.MapTasks[t.Id] = workerMapTasks
		_ = c.AssignMapTask(c.Workers[k], mapRequest)
	}

	*resp = AddTaskResponse{
		TaskId: t.Id,
	}

	return nil
}

func (c *Coordinator) AssignMapTask(w workerInfo, task AssignMapTaskRequest) error {
	fmt.Println("begin assign map task to workerInfo")
	client, err := rpc.DialHTTP("tcp", w.Address)
	if err != nil {
		fmt.Println("create mr_workers connect to workerInfo failed")
		return err
	}
	resp := AssignMapTaskResponse{}
	err = client.Call("Worker.AssignMapTask", &task, &resp)
	if err != nil {
		fmt.Printf("assign map task to work failed,err = %s\n", err)
		return err
	}

	return nil
}

func (c *Coordinator) AssignReduceTask(w workerInfo, req AssignReduceTaskRequest) error {
	fmt.Println("begin to assign reduce task to worker")
	client, err := rpc.DialHTTP("tcp", w.Address)
	if err != nil {
		fmt.Println("create mr_workers connect to worker failed")
		return err
	}
	resp := AssignReduceTaskResponse{}
	err = client.Call("Worker.AssignReduceTask", &req, &resp)
	if err != nil {
		fmt.Printf("assign map task to work failed,err = %s\n", err)
		return err
	}

	return nil
}

// server start a thread that listens for RPCs from workerInfo.go
func (c *Coordinator) server() {
	err := rpc.Register(c)
	if err != nil {
		fmt.Println("register coordinator failed")
		return
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", c.Address)
	if err != nil {
		fmt.Printf("coordinator listern tcp failed %s\n", err)
		return
	}
	go func() {
		err := http.Serve(listener, nil)
		if err != nil {
			fmt.Println("coordinator http serve failed")
		}
	}()

	fmt.Println("coordinator start success")
}

// MakeCoordinator create a Coordinator.
func MakeCoordinator() *Coordinator {
	c := Coordinator{
		Address:     "localhost:8500",
		Workers:     []workerInfo{},
		Tasks:       map[string]*Task{},
		MapTasks:    map[string]map[int]*MapTask{},
		ReduceTasks: map[string]map[int]*ReduceTask{},
		ReduceSize:  2,
	}

	c.server()
	return &c
}
