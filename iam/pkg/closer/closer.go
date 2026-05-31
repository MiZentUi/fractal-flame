package closer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
)

var global = &closer{}

type closer struct {
	mtx   sync.Mutex
	once  sync.Once
	funcs []func(context.Context) error
}

func Add(fn func(context.Context) error) {
	global.mtx.Lock()
	defer global.mtx.Unlock()

	global.funcs = append(global.funcs, fn)
}

func CloseAll(ctx context.Context) error {
	var out error

	global.once.Do(func() {
		global.mtx.Lock()
		funcs := global.funcs
		global.mtx.Unlock()

		slog.Debug("Closing all dependencies...")

		for len(funcs) > 0 {
			select {
			case <-ctx.Done():
				out = ctx.Err()
				return
			default:
			}

			curr := funcs[len(funcs)-1]
			funcs = funcs[:len(funcs)-1]

			out = errors.Join(out, safety(ctx, curr))
		}
	})

	return out
}

func safety(ctx context.Context, fn func(ctx context.Context) error) (out error) {
	defer func() {
		r := recover()
		if r != nil {
			out = fmt.Errorf("panic: %v", r)
		}
	}()

	return fn(ctx)
}
