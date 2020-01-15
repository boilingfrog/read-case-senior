package main

import (
	"fmt"
	"sync"
)

type lock struct {
	c chan struct{}
}

func newLock() lock {
	var l lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l lock) lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l lock) unLock() {
	l.c <- struct{}{}
}

var conter int

func TryLock() {
	var l = newLock()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.lock() {
				fmt.Println("err lock")
				return
			}
			counter++
			fmt.Println("add", counter)
			l.unLock()
		}()
	}
	wg.Wait()
}
