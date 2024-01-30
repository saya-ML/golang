package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Done:", Run(2, 1_000))
}

func Run(workersCount, worksToDo int) uint32 {
	var counter atomic.Uint32
	g := Generator{
		workers: make(chan struct{}, workersCount),
	}
	for i := 0; i < worksToDo; i++ {
		go g.generateStruct(i, &counter)
	}
	g.wg.Wait()
	return counter.Load()
}

type Generator struct {
	workers chan struct{}
	wg      sync.WaitGroup
}

func (g *Generator) generateStruct(timeToSleep int, pCounter *atomic.Uint32) {
	g.workers <- struct{}{}
	g.wg.Add(1)
	go func(tts int) {
		time.Sleep(time.Duration(tts % 100))
		<-g.workers
		pCounter.Add(1)
		g.wg.Done()
	}(timeToSleep)
}
