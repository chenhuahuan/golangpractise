package main

import (
	"fmt"
	"sync"
)


//WaitGroupWrapper 并行管理器
type WaitGroupWrapper struct {
	num int64
	sync.WaitGroup
}

func main() {

	data  := make(map[interface{}] map[interface{}] interface{})

	fmt.Print(data[2])

	var mywait WaitGroupWrapper
	mywait.WaitGroup.Add(1)




}
