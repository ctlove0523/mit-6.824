package map_reduce

import "sync"

type RpcRequest struct {
}

type RpcResponse struct {
}

type RegisterWorkerRequest struct {
	Id      int
	Address string
}

type RegisterWorkerResponse struct {
	Id                 int
	CoordinatorAddress string
}

type ReportMapTaskProgressRequest struct {
	WorkerId   int
	TaskId     string
	TaskStatus int
	Outputs    []MapTaskOutput
}

// ReportReduceTaskProgressRequest report reduce task progress
type ReportReduceTaskProgressRequest struct {
	WorkerId   int
	TaskId     string
	TaskStatus int
	Outputs    []ReduceTaskOutput
}

// AddTaskRequest 创建map-reduce任务请求
type AddTaskRequest struct {
	FileNames  []string
	ReduceSize int
}

// AddTaskResponse 创建map-reduce任务响应
type AddTaskResponse struct {
	TaskId string
}

// workerInfo 用于表示已经注册的worker
type workerInfo struct {
	Id      int    // 唯一表示一个执行任务的worker
	Address string // 执行任务的worker的地址，用于RPC调用
}

type Task struct {
	Id           string
	mapStatus    int
	mapLock      sync.Mutex
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

type AssignMapTaskRequest struct {
	TaskId     string
	MapTaskId  int
	Files      []string
	ReduceSize int
}

type AssignMapTaskResponse struct {
	Result bool
}

type AssignReduceTaskRequest struct {
	TaskId       string
	FileNames    []string
	ReduceNumber int
}

type AssignReduceTaskResponse struct {
	Result bool
}

type AssignTaskRequest struct {
	TaskId   int
	FileName string
}

type AssignTaskResponse struct {
	Result bool
}
