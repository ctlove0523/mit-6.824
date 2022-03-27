package map_reduce

import (
	"github.com/go-basic/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	api2 "mit-6.824/map-reduce/proto"
	"net"
	"time"
)

type workerInfo struct {
	id      string
	address string
}

func NewCoordinator(address string) *Coordinator {
	return &Coordinator{
		address: address,
		workers: []*workerInfo{},
		tasks:   map[string]*Task{},
	}
}

type Coordinator struct {
	address string
	workers []*workerInfo
	tasks   map[string]*Task
}

func (c *Coordinator) Start() error {
	log.Println("api begin to start")
	listener, err := net.Listen("tcp", "localhost:4789")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api2.RegisterCoordinatorServerServer(s, c)
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
func (c *Coordinator) ReportMapTaskProgress(ctx context.Context, req *api2.ReportMapTaskProgressRequest) (*api2.ReportMapTaskProgressResponse, error) {
	log.Println("begin to process map task progress report")
	log.Println(req.Output)
	task := c.tasks[req.TaskId]
	task.mapTask.subTasks[req.Id].state = Finished
	task.mapTask.UpdateSubTask(int(req.Id), Finished, req.Output)

	go func() {
		if task.mapTask.IsFinished() {
			log.Println("map task finished")
			err := c.executeReduceTask(req.TaskId)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	return &api2.ReportMapTaskProgressResponse{Result: true}, nil
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
		req := &api2.CreateReduceTaskRequest{
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
		client := api2.NewWorkerServerClient(conn)

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
func (c *Coordinator) ReportReduceTaskProgress(ctx context.Context, req *api2.ReportReduceTaskProgressRequest) (*api2.ReportReduceTaskProgressResponse, error) {
	log.Println("begin to process map task progress report")
	log.Println(req.Output)
	task := c.tasks[req.TaskId]
	task.reduceTask.subTasks[req.Id].state = Finished
	task.reduceTask.subTasks[req.Id].Output = req.Output

	resp := &api2.ReportReduceTaskProgressResponse{
		Result: true,
	}

	return resp, nil
}
func (c *Coordinator) RegisterWorker(ctx context.Context, req *api2.RegisterWorkerRequest) (*api2.RegisterWorkerResponse, error) {
	log.Println("begin to process worker register")

	worker := &workerInfo{
		id:      newWorkerId(),
		address: req.GetAddress(),
	}
	c.workers = append(c.workers, worker)

	resp := &api2.RegisterWorkerResponse{
		Id:                 worker.id,
		CoordinatorAddress: c.address,
	}

	return resp, nil

}

func (c *Coordinator) CreateTask(ctx context.Context, req *api2.CreateMapReduceTaskRequest) (*api2.CreateMapReduceTaskResponse, error) {
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
			mapReq := &api2.CreateMapTaskRequest{
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
			client := api2.NewWorkerServerClient(conn)

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

	resp := &api2.CreateMapReduceTaskResponse{
		Id: task.Id,
	}

	return resp, nil
}

func newWorkerId() string {
	return uuid.New()
}
