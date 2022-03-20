package map_reduce

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"
	"unicode"
)

type Worker struct {
	Id                 int
	Address            string
	CoordinatorAddress string
}

func NewWorker(id int, address, CoordinatorAddress string) *Worker {
	w := &Worker{
		Id:                 id,
		Address:            address,
		CoordinatorAddress: CoordinatorAddress,
	}

	w.start()
	w.register()

	return w
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

type KeyValuePair struct {
	Key   string
	Value int
}

func (kvp *KeyValuePair) String() string {
	return fmt.Sprintf("%v,%v", kvp.Key, kvp.Value)
}

func (w *Worker) reportMapTaskState(req *AssignMapTaskRequest, outputs []MapTaskOutput) {
	client, err := rpc.DialHTTP("tcp", w.CoordinatorAddress)
	if err != nil {
		fmt.Printf("connect to coordinator failed,err = %s\n", err)
		return
	}
	reportReq := ReportMapTaskProgressRequest{
		WorkerId:   w.Id,
		TaskId:     req.TaskId,
		TaskStatus: 1,
		Outputs:    outputs,
	}
	resp := RpcResponse{}
	err = client.Call("Coordinator.ReportMapTaskProgress", &reportReq, &resp)
	if err != nil {
		fmt.Println("report map task state failed")
	}

}

func (w *Worker) mapFunction(req *AssignMapTaskRequest) error {
	fmt.Printf("begin to execute map task,input files = %s\n", req.Files)

	ff := func(r rune) bool { return !unicode.IsLetter(r) }
	files := make(map[int]*os.File)
	for _, fileName := range req.Files {
		// 打开文件
		f, err := os.Open(fileName)
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
			hashVal := HashFunction(v.Key) % req.ReduceSize
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

	var fileNames []MapTaskOutput
	for k, v := range files {
		fileNames = append(fileNames, MapTaskOutput{
			FileName:       v.Name(),
			ReduceRegionNo: k,
		})
	}
	w.reportMapTaskState(req, fileNames)
	return nil
}

func (w *Worker) AssignMapTask(req *AssignMapTaskRequest, resp *AssignMapTaskResponse) error {
	fmt.Println("get map task")
	go func() {
		w.mapFunction(req)
	}()

	return nil
}

func (w *Worker) AssignReduceTask(req *AssignReduceTaskRequest, resp *AssignMapTaskResponse) error {
	fmt.Println("get reduce task")
	go func() {
		w.reduceFunction(req)
	}()

	return nil
}

func (w *Worker) reduceFunction(req *AssignReduceTaskRequest) error {
	files := req.FileNames
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
	writeFile, err := os.Create(fmt.Sprintf("reduce-%d.txt", req.ReduceNumber))
	if err != nil {
		fmt.Println("create reduce file failed")
		return err
	}

	for k, v := range wordMap {
		_, err := fmt.Fprintf(writeFile, "%v,%v\n", k, v)
		if err != nil {
			fmt.Println("append to file faield")
			writeFile.Close()
			return err
		}
	}

	return nil
}
func (w *Worker) createFile(index int) (*os.File, error) {
	intermediateFileName := fmt.Sprintf("intermediate-%d-%d.txt", w.Id, index)
	return os.Create(intermediateFileName)
}

func (w *Worker) AssignTask(res *AssignTaskRequest, resp *AssignTaskResponse) error {
	fmt.Println("get assigned task")
	return nil
}

func (w *Worker) start() error {
	err := rpc.Register(w)
	if err != nil {
		fmt.Println("workerInfo rpc register failed")
		return err
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", w.Address)
	if err != nil {
		fmt.Println("workerInfo listener failed")
		return err
	}

	go http.Serve(listener, nil)

	return nil
}

func (w *Worker) register() {
	client, err := rpc.DialHTTP("tcp", w.CoordinatorAddress)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	req := RegisterWorkerRequest{
		Id:      w.Id,
		Address: w.Address,
	}
	resp := &RegisterWorkerResponse{}
	err = client.Call("Coordinator.RegisterWorker", req, resp)
	if err != nil {
		fmt.Println("register failed")
		panic(err.Error())
	}

}
