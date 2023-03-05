package impl

import (
	"github.com/slarkdarr/Tugas-2-Kriptografi/internal"
	"sync"
)

type (
	circular struct {
		count int
		max   int
		data  []internal.Executor
		sync.Mutex
	}
)

func (c *circular) ResetCount() {
	c.Lock()
	defer c.Unlock()
	c.count = 0
}

func (c *circular) Execute(chunk []byte) []byte {
	c.Lock()
	defer func() {
		c.count++
		if c.count >= c.max {
			c.count = 0
		}
		c.Unlock()
	}()
	return c.data[c.count].Execute(chunk)
}

func NewCircularPermutation(count int) internal.GroupExecutor {
	data := make([]internal.Executor, count)
	for i := 0; i < count; i++ {
		data[i] = NewPermutation()
	}
	return &circular{
		count: 0,
		max:   count,
		data:  data,
	}
}

func NewCircularSubsitution(count int) internal.GroupExecutor {
	data := make([]internal.Executor, count)
	for i := 0; i < count; i++ {
		data[i] = NewSubstitution()
	}
	return &circular{
		count: 0,
		max:   count,
		data:  data,
	}
}
