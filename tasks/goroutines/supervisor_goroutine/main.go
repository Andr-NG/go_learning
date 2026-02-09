package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main: starting")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Go(func() {
		supervisor(ctx)
	})

	// Simulate running system
	time.Sleep(3 * time.Second)

	fmt.Println("main: initiating shutdown")
	cancel() // ONLY place cancel is called

	wg.Wait()

}

type Worker struct {
	doneCh chan struct{}
	stopCh chan struct{}
}

func newWorker() *Worker {
	return &Worker{
		doneCh: make(chan struct{}),
		stopCh: make(chan struct{}),
	}
}

func (w *Worker) Close() {
	close(w.stopCh)
}
func (w *Worker) Wait() {
	<-w.doneCh
}

func (w *Worker) Run(s string) {
	defer close(w.doneCh)
	for {
		select {
		case <-w.stopCh:
			fmt.Println(s, "stopped")
			return
		default:
			fmt.Println(s, "performing work")
			time.Sleep(800 * time.Millisecond)
		}
	}
}

func supervisor(ctx context.Context) {

	ingestor := newWorker()
	processor := newWorker()
	collector := newWorker()

	go ingestor.Run("ingestor")
	go processor.Run("processor")
	go collector.Run("collector")

	// supervisor listens for context cancel
	<-ctx.Done()

	// supervisor closes ingestor channel to signal completion
	ingestor.Close()
	// supervisor reads from the ingestor done channel to ensure correct ordering.
	ingestor.Wait()

	processor.Close()
	processor.Wait()

	collector.Close()
	collector.Wait()

	fmt.Println("all done")
}
