package myWc

import (
	"fmt"
	"os"
)

type countJob struct {
	filename string
	counter  Counter
}

func newCountJob(filename string, counter Counter) *countJob {
	return &countJob{
		filename: filename,
		counter:  counter,
	}
}

func (cj *countJob) Process() result {
	file, err := os.Open(cj.filename)
	if err != nil {
		return result{
			filename: cj.filename,
			err:      fmt.Errorf(errFileOpen, cj.filename, err),
		}
	}
	defer file.Close()

	count, err := cj.counter.count(file)
	if err != nil {
		return result{
			filename: cj.filename,
			err:      fmt.Errorf(errFileProcess, cj.filename, err),
		}
	}

	return result{
		filename: cj.filename,
		count:    count,
		err:      nil,
	}
}
