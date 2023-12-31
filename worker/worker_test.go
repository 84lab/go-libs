package worker

import (
	"context"
	"sync"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestWorkerWithDefaultOptions(t *testing.T) {
	counter := 0
	worker := NewWorkerBuilder("test", func(context.Context) error {
		counter++
		return nil
	}).WithOptions(DefaultOptions(100 * time.Millisecond)).Build()

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	worker.Start(ctx, wg)

	time.Sleep(350 * time.Millisecond)
	cancel()
	wg.Wait()

	assert.Equal(t, 4, counter, "Should execute 4 times - 1st immidietly, and 3 after")
}

func TestWorkerStartsConsequently(t *testing.T) {
	counter := 0
	options := DefaultOptions(100 * time.Millisecond)
	options.RunConsequently = true

	worker := NewWorkerBuilder("test", func(context.Context) error {
		time.Sleep(100 * time.Millisecond)
		counter++
		return nil
	}).WithOptions(options).Build()

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	worker.Start(ctx, wg)

	time.Sleep(350 * time.Millisecond)
	cancel()
	wg.Wait()

	assert.Equal(t, 3, counter, "Should execute 3 times - 1st immidietly, and 2 after with delay between runs")
}
