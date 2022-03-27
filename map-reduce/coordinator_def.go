package map_reduce

import (
	"github.com/go-basic/uuid"
)

func NewTask(name string, inputs []string, mapSize, reduceSize int32) *Task {
	task := &Task{
		Id:         uuid.New(),
		Name:       name,
		Inputs:     inputs,
		MapSize:    mapSize,
		ReduceSize: reduceSize,
		Outputs:    []string{},
		state:      WaitProcess,
	}

	return task
}

type Task struct {
	Id         string
	Name       string
	Inputs     []string // 已经分割好的文件列表
	MapSize    int32    // 当和worker的数量相等
	ReduceSize int32    // reduce后输出文件的数量，也是map任务结果的分区数量
	Outputs    []string
	state      TaskState
	mapTask    *MapTask
	reduceTask *ReduceTask
}

type MapTask struct {
	Id       string
	Inputs   []string // 执行一个Map任务的文件列表
	state    TaskState
	subTasks []*SubMapTask    // 每个worker执行的子任务
	Outputs  map[int][]string // key为按照ReduceSize分区后的值，取值范围[0,ReduceSize)
}

func (mt *MapTask) calculateState() TaskState {
	s := Finished
	for _, v := range mt.subTasks {
		if v.state < s {
			s = v.state
		}
		if s == Failed {
			break
		}
	}

	return s
}

func (mt *MapTask) UpdateSubTask(id int, state TaskState, results map[uint32]string) {
	st := mt.subTasks[id]
	st.state = state
	st.Outputs = results
	if st.Outputs == nil {
		st.Outputs = map[uint32]string{}
	}

	mt.state = mt.calculateState()

	for k, v := range results {
		outs, ok := mt.Outputs[int(k)]
		if !ok {
			outs = []string{}
		}
		outs = append(outs, v)
		mt.Outputs[int(k)] = outs
	}

}

func (mt *MapTask) IsFinished() bool {
	res := true
	for _, v := range mt.subTasks {
		res = res && v.IsFinished()
	}

	return res
}

type SubMapTask struct {
	Id       uint16
	WorkerId string
	state    TaskState
	Inputs   []string          // 一个worker执行map任务的输入
	Outputs  map[uint32]string // key为按照ReduceSize分区后的值，取值范围[0,ReduceSize)
}

func (smt *SubMapTask) IsFinished() bool {
	return smt.state == Finished
}

type ReduceTask struct {
	Id       string
	Inputs   map[int][]string // 执行一个Reduce任务的文件列表
	state    TaskState
	subTasks []*SubReduceTask // 每个worker执行的子任务
	Outputs  []string
}

func (rt *ReduceTask) calculateState() TaskState {
	s := Finished
	for _, v := range rt.subTasks {
		if v.state < s {
			s = v.state
		}
		if s == Failed {
			break
		}
	}

	return s
}

func (rt *ReduceTask) UpdateSubTask(id int, state TaskState, output string) {
	st := rt.subTasks[id]
	st.state = state
	st.Output = output
	rt.state = rt.calculateState()
}

func (rt *ReduceTask) IsFinished() bool {
	res := true
	for _, v := range rt.subTasks {
		res = res && v.IsFinished()
	}

	return res
}

type SubReduceTask struct {
	Id             uint
	WorkerId       string
	ReduceRegionId int // reduce 任务分区Id
	state          TaskState
	Inputs         []string // 一个worker执行reduce任务的输入
	Output         string
}

func (srt *SubReduceTask) IsFinished() bool {
	return srt.state == Finished
}

type TaskState int

const (
	_ TaskState = iota
	Failed
	WaitProcess
	Processing
	Finished
)

type TaskDetail struct {
	TaskId string `json:"task_id"`
	State  int    `json:"state"`
}
