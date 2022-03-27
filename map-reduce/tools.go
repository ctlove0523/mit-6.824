package map_reduce

import (
	"fmt"
	"github.com/go-basic/uuid"
	"hash/fnv"
)

func HashFunction(key string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(key))
	if err != nil {
		fmt.Println("calculate hash Value failed")
		return 0
	}
	return h.Sum32() & 0x7fffffff
}

func NewTaskId() string {
	return uuid.New()
}

func NewWorkerId() string {
	return uuid.New()
}
