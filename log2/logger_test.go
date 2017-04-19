package log2

import (
	"sync"
	"testing"
)

func TestInfo(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			Info(WithFuncName(WithUserID(New(), "dummy-user-id"), "dummy-func-name"), "log string")
		}(&wg)
	}
	wg.Wait()
}
