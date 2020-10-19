package worker

type ThreadPool struct {
	f        chan func()
	finished chan struct{}
}

func New(n int) *ThreadPool {

	fDone := make(chan struct{})

	threadPool := &ThreadPool{
		f:        make(chan func()),
		finished: make(chan struct{}),
	}

	for i := 0; i < n; i++ {
		go func() {
			for f := range threadPool.f {
				f()
			}
			fDone <- struct{}{}
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			_ = <-fDone
		}
		threadPool.finished <- struct{}{}

	}()
	return threadPool

}

func (t *ThreadPool) Add(f func()) {
	t.f <- f
}

func (t *ThreadPool) Wait() {
	close(t.f)
	_ = <-t.finished

}
