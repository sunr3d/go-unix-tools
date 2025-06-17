package myWc

import (
	"fmt"
	"os"

	"go-unix-tools/pkg/workerpool"
)

type processor struct {
	opts    *options
	counter Counter
}

type result struct {
	filename string
	count    int
	err      error
}

func newProcessor(args []string) (*processor, error) {
	opts, err := newOptions(args)
	if err != nil {
		return nil, fmt.Errorf("newOptions: %w", err)
	}

	var ctr Counter
	switch {
	case opts.countLines:
		ctr = &lineCounter{}
	case opts.countWords:
		ctr = &wordCounter{}
	case opts.countChars:
		ctr = &charCounter{}
	}

	return &processor{
		opts:    opts,
		counter: ctr,
	}, nil
}

func (p *processor) processFile() []result {
	filesLen := len(p.opts.files)

	pool := workerpool.New[result](filesLen, filesLen)
	pool.Start()

	for _, file := range p.opts.files {
		pool.Submit(newCountJob(file, p.counter))
	}

	results := make([]result, filesLen)
	for i := 0; i < filesLen; i++ {
		results[i] = <-pool.Results()
	}

	pool.Stop()
	return results
}

func (p *processor) printResults(results []result) {
	for _, res := range results {
		if res.err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", res.filename, res.err)
			continue
		}
		fmt.Printf("%d\t%s\n", res.count, res.filename)
	}
}

func (p *processor) run() error {
	results := p.processFile()
	if len(results) == 0 {
		return fmt.Errorf(errNoFiles)
	}

	p.printResults(results)
	return nil
}
