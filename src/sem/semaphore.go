package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
)
import "golang.org/x/sync/semaphore"

var (
	maxWorkers = runtime.GOMAXPROCS(0)                    //worker 数量
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量
	task       = make([]int, maxWorkers*4)                //任务数
)

func main() {
	ctx := context.Background()

	fmt.Println(maxWorkers)

	for i := range task {
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		go func() {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond)
			task[i] = i + 1
		}()
	}

	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败: %v", err)
	}

	fmt.Println(task)

}
