package main

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := NewGorunPool(5)
	for i := 0; i < 10000; i++ {
		c := i
		job := func() error {
			time.Sleep(1 * time.Second)
			t.Error("do thread!", c)
			return nil
		}
		task := NewGorunTask(job, nil)
		pool.Execute(task)
	}
	time.Sleep(3 * time.Second)
}
