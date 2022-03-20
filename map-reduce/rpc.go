package map_reduce

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

type ReportWorkProgressRequest struct {
	WorkerId   int
	TaskId     string
	TaskStatus int
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
