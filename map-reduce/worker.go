package map_reduce

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	api2 "mit-6.824/map-reduce/proto"
	"net"
	"os"
	"strings"
	"time"
	"unicode"
)

type Worker struct {
	Id                 string
	Address            string
	CoordinatorAddress string
	client             api2.CoordinatorServerClient
	server             api2.WorkerServerServer
}

func (w *Worker) Start() {
	log.Println("begin to start worker")

	// connect to coordinator
	conn, err := grpc.Dial(w.CoordinatorAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect to coordinator failed: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("close connect to coordinator failed")
		}
	}()

	c := api2.NewCoordinatorServerClient(conn)
	w.client = c

	w.register()
	// 启动rpc server
	listener, err := net.Listen("tcp", w.Address)
	if err != nil {
		log.Fatalf("worker start failed,error = %s", err)
	}
	grpcServer := grpc.NewServer()
	api2.RegisterWorkerServerServer(grpcServer, w)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (w *Worker) register() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &api2.RegisterWorkerRequest{
		Address: w.Address,
	}
	resp, err := w.client.RegisterWorker(ctx, req)
	if err != nil {
		log.Fatalf("register failed,error = %s", err)
	}
	w.Id = resp.Id
}

func (w *Worker) CreateMapTask(ctx context.Context, req *api2.CreateMapTaskRequest) (*api2.CreateMapTaskResponse, error) {
	log.Println("get map task")
	go func() {
		w.executeMapTask(req)
	}()
	return &api2.CreateMapTaskResponse{
		Id:     req.Id,
		TaskId: req.TaskId,
		Result: true,
	}, nil
}

func (w *Worker) executeMapTask(req *api2.CreateMapTaskRequest) error {
	log.Printf("begin to execute map task for %s map-reduce task,id = %d", req.TaskId, req.Id)

	ff := func(r rune) bool { return !unicode.IsLetter(r) }
	files := make(map[int]*os.File)
	for _, input := range req.Inputs {
		// 打开文件
		f, err := os.Open(input)
		if err != nil {
			fmt.Printf("open file failed,err = %s\n", err)
			return err
		}

		var keyValues []KeyValuePair
		buf := bufio.NewReader(f)

		for {
			lineBytes, _, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			line := string(lineBytes)
			if len(strings.Trim(line, "")) == 0 {
				continue
			}
			words := strings.FieldsFunc(line, ff)
			for _, word := range words {
				pair := KeyValuePair{
					Key:   word,
					Value: 1,
				}
				keyValues = append(keyValues, pair)
			}
		}

		for _, v := range keyValues {
			hashVal := int(HashFunction(v.Key) % req.GetReduceSize())

			f, ok := files[hashVal]
			if !ok {
				file, err := w.createFile(hashVal)
				if err != nil {
					fmt.Println("create file failed")
					break
				}
				files[hashVal] = file
				f = file
			}

			_, err := fmt.Fprintln(f, v.String())
			if err != nil {
				fmt.Println("append to file failed")
				f.Close()
				return err
			}
		}
	}

	outputs := make(map[uint32]string)
	for k, v := range files {
		outputs[uint32(k)] = v.Name()
	}
	reportProgressRequest := &api2.ReportMapTaskProgressRequest{
		TaskId:  req.TaskId,
		Id:      req.Id,
		State:   1,
		Outputs: outputs,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := w.client.ReportMapTaskProgress(ctx, reportProgressRequest)
	if err != nil {
		log.Println("report map task state failed")
		return err
	}
	log.Printf("report map task state %v", resp.Result)
	return nil
}

func (w *Worker) CreateReduceTask(ctx context.Context, req *api2.CreateReduceTaskRequest) (*api2.CreateReduceTaskResponse, error) {
	log.Println("begin to process reduce task")
	go func() {
		w.executeReduceTask(req)
	}()
	resp := &api2.CreateReduceTaskResponse{}

	return resp, nil
}

func (w *Worker) executeReduceTask(req *api2.CreateReduceTaskRequest) error {
	files := req.Inputs
	fmt.Println(files)

	wordMap := map[string]int{}
	ff := func(r rune) bool { return !unicode.IsLetter(r) }
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("open file %s failed\n", file)
			f.Close()
			continue
		}

		buf := bufio.NewReader(f)
		for {
			lineBytes, _, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			line := string(lineBytes)
			if len(strings.Trim(line, "")) == 0 {
				continue
			}
			words := strings.FieldsFunc(line, ff)
			for _, word := range words {
				wordMap[word] += 1
			}
		}
	}

	// reduce 结果写入到文件中
	fileName := fmt.Sprintf("reduce-%d.txt", req.ReduceRegionId)
	writeFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create reduce file failed")
		return err
	}

	for k, v := range wordMap {
		_, err := fmt.Fprintf(writeFile, "%v,%v\n", k, v)
		if err != nil {
			fmt.Println("append to file failed")
			writeFile.Close()
			return err
		}
	}

	outputs := api2.ReportReduceTaskProgressRequest{
		TaskId: req.TaskId,
		Id:     req.Id,
		State:  3,
		Output: fileName,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := w.client.ReportReduceTaskProgress(ctx, &outputs)
	if err != nil {
		log.Printf("report reduce task progress failed,error = %s", err)
		return err
	}

	log.Println(resp)

	return nil
}

func (w *Worker) createFile(index int) (*os.File, error) {
	intermediateFileName := fmt.Sprintf("intermediate-%s-%d.txt", w.Id, index)
	return os.Create(intermediateFileName)
}

type KeyValuePair struct {
	Key   string
	Value int
}

func (kvp *KeyValuePair) String() string {
	return fmt.Sprintf("%v,%v", kvp.Key, kvp.Value)
}
