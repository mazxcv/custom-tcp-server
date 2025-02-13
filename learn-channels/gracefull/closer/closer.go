package closer

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

func (c *Closer) Add(f Func) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f)
}

type Func func(ctx context.Context) error

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		msgs     = make([]string, 0, len(c.funcs))
		complete = make(chan struct{}, 1)
	)

	go func() {
		for _, f := range c.funcs {
			if err := f(ctx); err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}
		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("context done: %w", ctx.Err())
	}

	if len(msgs) > 0 {
		return fmt.Errorf("%s", strings.Join(msgs, "\n"))
	}

	return nil

}
