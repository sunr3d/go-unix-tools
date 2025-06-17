package workerpool

import "sync"

type Job[T any] interface {
	Process() T
}

type Pool[T any] struct {
	maxWorkers int
	jobs       chan Job[T]
	results    chan T
	wg         sync.WaitGroup
}

func New[T any](workers, buffer int) *Pool[T] {
	return &Pool[T]{
		maxWorkers: workers,
		jobs:       make(chan Job[T], buffer),
		results:    make(chan T, buffer),
	}
}

func (p *Pool[T]) worker() {
	defer p.wg.Done()
	for job := range p.jobs {
		p.results <- job.Process()
	}
}

func (p *Pool[T]) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *Pool[T]) Submit(job Job[T]) {
	p.jobs <- job
}

func (p *Pool[T]) Results() <-chan T {
	return p.results
}

func (p *Pool[T]) Stop() {
	close(p.jobs)
	p.wg.Wait()
	close(p.results)
}
