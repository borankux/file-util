package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	//fmt.Printf("run with:%d \n", n)
}

func main() {
	t := time.Now()
	defer ants.Release()
	runTimes := 100000
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(100, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	//ants.WithNonblocking(true)
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()

	fmt.Printf("running goroutines:%d\n", p.Running())
	fmt.Printf("finished all tasks, result is %d \n", sum)
	fmt.Printf("spent time:%v\n", time.Since(t))
}
