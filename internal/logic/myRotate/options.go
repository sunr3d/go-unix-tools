package myRotate

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type options struct {
	files       []string
	archiveDir  string
	workerCount int
}

func newOptions(args []string) (*options, error) {
	opts := &options{
		workerCount: runtime.NumCPU(),
	}

	fs := flag.NewFlagSet("myRotate", flag.ContinueOnError)
	fs.StringVar(&opts.archiveDir, "a", "", "директория для архивации")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf(errParseFlags, err)
	}

	opts.files = fs.Args()
	if len(opts.files) == 0 {
		return nil, fmt.Errorf(errNoFiles)
	}

	for _, file := range opts.files {
		if filepath.Ext(file) != ".log" {
			return nil, fmt.Errorf(errNotALogFile, file)
		}
	}

	if opts.archiveDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf(errCwd, err)
		}
		opts.archiveDir = cwd
	} else {
		if err := opts.validateArchiveDir(); err != nil {
			return nil, fmt.Errorf(errArchiveDir, opts.archiveDir, err)
		}
	}

	return opts, nil
}

func (opts *options) validateArchiveDir() error {
	info, err := os.Stat(opts.archiveDir)
	if err != nil {
		return fmt.Errorf("не удалось получить инфо о директории")
	}
	if !info.IsDir() {
		return fmt.Errorf("не является директорией")
	}

	return nil
}
