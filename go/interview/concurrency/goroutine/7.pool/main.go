package main

import (
	"sync"
	"time"
)

/**
* 一个固定大小的goroutine 协程池
* 拦截和排队过量的并发协程数量
 */
type GorunTask struct {
	Id       int64
	Name     string
	Status   string
	Run      func() error
	Callback func()
	Err      error
}

func NewGorunTaskWithName(name string, run func() error, callback func()) *GorunTask {
	task := NewGorunTask(run, callback)
	task.Name = name
	return task
}

func NewGorunTask(run func() error, callback func()) *GorunTask {
	t0 := time.Now()
	task := GorunTask{
		Id:       t0.UnixNano(),
		Run:      run,
		Callback: callback,
		Status:   "inited",
	}
	return &task
}

type GorunPool struct {
	Size    int
	Ticket  int
	Stoped  bool
	ResChan chan *GorunTask
	lock    sync.RWMutex
	Tasks   []*GorunTask
	Once    sync.Once
}

func NewGorunPool(size int) *GorunPool {
	pool := GorunPool{
		Size:    size,
		Ticket:  0,
		Stoped:  false,
		ResChan: make(chan *GorunTask),
		Tasks:   []*GorunTask{},
	}
	go pool.Once.Do(func() {
		for {
			if pool.Stoped {
				break
			}
			tk := <-pool.ResChan
			if tk.Callback != nil {
				go tk.Callback()
			}
			pool.lock.Lock()
			pool.Ticket--
			pool.lock.Unlock()
		}
	})
	go pool.execqueue()
	return &pool
}

func (pool *GorunPool) call(task *GorunTask) {
	pool.Ticket++
	go func() {
		task.Status = "running"
		task.Err = task.Run()
		task.Status = "exected!"
		pool.ResChan <- task
	}()
}

func (pool *GorunPool) execqueue() {
	for {
		if pool.Stoped {
			break
		}
		if len(pool.Tasks) > 0 && pool.Size-pool.Ticket > 0 {
			pool.lock.Lock()
			task := pool.Tasks[0]
			pool.Tasks = pool.Tasks[1:len(pool.Tasks)]
			pool.lock.Unlock()
			pool.call(task)
		} else {
			time.Sleep(100 * time.Microsecond)
		}
	}
}

func (pool *GorunPool) Execute(task *GorunTask) {
	// 1.判断当前pool的状态
	pool.lock.Lock()
	if pool.Ticket >= pool.Size { //队列等待
		task.Status = "waitting"
		pool.Tasks = append(pool.Tasks, task)
	} else {
		pool.call(task)
	}
	pool.lock.Unlock()
}
