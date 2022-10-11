package gtool

import (
	"sync"
	"time"
)

func GoroutineExecutorService(services ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(services) + 1)
	for _, item := range services {
		go item()
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}
	wg.Wait()
}
