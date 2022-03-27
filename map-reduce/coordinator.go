package map_reduce

import (
	"encoding/json"
	"errors"
	"github.com/go-basic/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	api "mit-6.824/map-reduce/proto"
	"net"
	"net/http"
	"time"
)

type workerInfo struct {
	id      string
	address string
}

func NewCoordinator(address string, httpAddress string) *Coordinator {
	return &Coordinator{
		address:     address,
		httpAddress: httpAddress,
		workers:     []*workerInfo{},
		tasks:       map[string]*Task{},
		stopCh:      make(chan struct{}),
	}
}

type Coordinator struct {
	address     string
	httpAddress string
	workers     []*workerInfo
	tasks       map[string]*Task

	stopCh chan struct{}
}

func (c *Coordinator) Start() error {
	log.Println("coordinator begin to start")

	// 启动grpc server
	go func() {
		log.Println("coordinator start grpc server")
		listener, err := net.Listen("tcp", "localhost:4789")
		if err != nil {
			log.Fatalf("coordinator grpc failed to listen: %v", err)
		}
		s := grpc.NewServer()
		api.RegisterCoordinatorServerServer(s, c)

		if err := s.Serve(listener); err != nil {
			log.Printf("coordinator grpc server failed to serve: %v", err)
		}
	}()

	// 启动http server
	go func() {
		router := httprouter.New()
		router.GET("/tasks/:id", c.queryTask)
		err := http.ListenAndServe(c.httpAddress, router)
		if err != nil {
			log.Printf("coordinator http server failed to serve: %v", err)
		}

	}()

	<-c.stopCh
	return nil
}
func (c *Coordinator) queryTask(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId := params.ByName("id")
	task, ok := c.tasks[taskId]
	if !ok {
		writer.WriteHeader(404)
		return
	}

	resp := TaskDetail{
		TaskId: taskId,
		State:  int(task.state),
	}
	body, err := json.Marshal(resp)
	if err != nil {
		log.Printf("json marshal failed")
		writer.WriteHeader(500)
	}
	writer.Write(body)
	writer.WriteHeader(200)
}

func (c *Coordinator) QueryTask(ctx context.Context, req *api.QueryTaskRequest) (*api.QueryTaskResponse, error) {
	task, ok := c.tasks[req.TaskId]
	if !ok {
		return nil, errors.New("task not found")
	}
	resp := &api.QueryTaskResponse{
		TaskId: task.Id,
		State:  int32(task.state),
	}

	return resp, nil
}

func (c *Coordinator) ReportMapTaskProgress(ctx context.Context, req *api.ReportMapTaskProgressRequest) (*api.ReportMapTaskProgressResponse, error) {
	log.Println("begin to process map task progress report")
	log.Println(req.Outputs)
	task := c.tasks[req.TaskId]
	task.mapTask.subTasks[req.Id].state = Finished
	task.mapTask.UpdateSubTask(int(req.Id), Finished, req.Outputs)

	go func() {
		if task.mapTask.IsFinished() {
			log.Println("map task finished")
			err := c.executeReduceTask(req.TaskId)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	return &api.ReportMapTaskProgressResponse{Result: true}, nil
}

func (c *Coordinator) executeReduceTask(taskId string) error {
	log.Println("begin to execute reduce task")
	task := c.tasks[taskId]
	mapOutputs := task.mapTask.Outputs
	log.Printf("map outputs is %v", mapOutputs)

	reduceTask := &ReduceTask{
		Id:       NewTaskId(),
		Inputs:   mapOutputs,
		state:    WaitProcess,
		subTasks: []*SubReduceTask{},
		Outputs:  []string{},
	}

	task.reduceTask = reduceTask

	for k, v := range reduceTask.Inputs {
		subTask := &SubReduceTask{
			Id:             uint(k),
			WorkerId:       c.workers[k%len(c.workers)].id,
			ReduceRegionId: k,
			state:          WaitProcess,
			Inputs:         v,
		}
		reduceTask.subTasks = append(reduceTask.subTasks, subTask)
	}

	for k, v := range reduceTask.Inputs {
		req := &api.CreateReduceTaskRequest{
			TaskId:         taskId,
			Id:             uint32(k),
			ReduceRegionId: uint32(k),
			Inputs:         v,
		}

		woker := c.workers[k%len(c.workers)]
		conn, err := grpc.Dial(woker.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := api.NewWorkerServerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := client.CreateReduceTask(ctx, req)
		if err != nil {
			log.Printf("create reduce task failed,error = %s", err)
			return err
		}

		log.Println(resp)
	}

	return nil
}
func (c *Coordinator) ReportReduceTaskProgress(ctx context.Context, req *api.ReportReduceTaskProgressRequest) (*api.ReportReduceTaskProgressResponse, error) {
	log.Println("begin to process map task progress report")
	log.Println(req.Output)
	task := c.tasks[req.TaskId]
	task.reduceTask.subTasks[req.Id].state = Finished
	task.reduceTask.subTasks[req.Id].Output = req.Output

	if task.reduceTask.IsFinished() {
		log.Println("reduce task finished")
	}
	resp := &api.ReportReduceTaskProgressResponse{
		Result: true,
	}

	return resp, nil
}
func (c *Coordinator) RegisterWorker(ctx context.Context, req *api.RegisterWorkerRequest) (*api.RegisterWorkerResponse, error) {
	log.Println("begin to process worker register")

	worker := &workerInfo{
		id:      newWorkerId(),
		address: req.GetAddress(),
	}
	c.workers = append(c.workers, worker)

	resp := &api.RegisterWorkerResponse{
		Id:                 worker.id,
		CoordinatorAddress: c.address,
	}

	return resp, nil

}

func (c *Coordinator) CreateTask(ctx context.Context, req *api.CreateMapReduceTaskRequest) (*api.CreateMapReduceTaskResponse, error) {
	log.Println("begin to create map reduce task")

	task := NewTask(req.GetName(), req.GetInputs(), req.GetMapSize(), req.GetReduceSize())

	c.tasks[task.Id] = task

	go func() {
		// 划分map任务的输入
		inputs := map[int][]string{}
		for i := 0; i < len(req.Inputs); i++ {
			index := i % len(c.workers)
			files, ok := inputs[index]
			if !ok {
				files = []string{req.Inputs[i]}
			} else {
				files = append(files, req.Inputs[i])
			}
			inputs[index] = files
		}

		// 任务分发给worker执行
		mapTask := &MapTask{
			Id:       NewTaskId(),
			Inputs:   req.Inputs,
			state:    WaitProcess,
			subTasks: make([]*SubMapTask, len(inputs)),
			Outputs:  map[int][]string{},
		}
		task.mapTask = mapTask

		for k, v := range inputs {
			mapReq := &api.CreateMapTaskRequest{
				TaskId:     task.Id,
				Id:         uint32(k),
				Inputs:     v,
				ReduceSize: 3,
			}

			subMapTask := &SubMapTask{
				Id:       uint16(k),
				WorkerId: c.workers[k].id,
				state:    WaitProcess,
				Inputs:   v,
				Outputs:  map[uint32]string{},
			}
			task.mapTask.subTasks[k] = subMapTask

			conn, err := grpc.Dial(c.workers[k].address, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			client := api.NewWorkerServerClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := client.CreateMapTask(ctx, mapReq)
			if err != nil {
				log.Fatalf("create map task failed")
			}

			subMapTask.state = Processing
			log.Println(resp.Result)

		}
	}()

	resp := &api.CreateMapReduceTaskResponse{
		Id: task.Id,
	}

	return resp, nil
}

func newWorkerId() string {
	return uuid.New()
}
