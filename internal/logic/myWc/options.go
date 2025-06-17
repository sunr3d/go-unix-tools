package myWc

import (
	"flag"
	"fmt"
)

type options struct {
	countLines bool
	countWords bool
	countChars bool
	files      []string
}

func newOptions(args []string) (*options, error) {
	fs := flag.NewFlagSet("myWc", flag.ContinueOnError)
	lines := fs.Bool("l", false, "Подсчет количества строк")
	words := fs.Bool("w", false, "Подсчет количества слов")
	chars := fs.Bool("m", false, "Подсчет количества символов")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf(errFlagParse, err)
	}

	if fs.NArg() < 1 {
		return nil, fmt.Errorf(errNoFiles)
	}

	flagCounter := 0
	fs.Visit(func(f *flag.Flag) {
		flagCounter++
	})

	if flagCounter > 1 {
		return nil, fmt.Errorf(errMultipleFlags)
	}

	opts := &options{
		countLines: *lines,
		countWords: *words,
		countChars: *chars,
		files:      fs.Args(),
	}

	if flagCounter == 0 {
		opts.countLines = false
		opts.countWords = true
		opts.countChars = false
	}
	return opts, nil
}
