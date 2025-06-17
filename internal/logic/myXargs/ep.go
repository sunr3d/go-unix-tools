package myXargs

import "fmt"

func Run(args []string) error {
	opts, err := newOptions(args)
	if err != nil {
		return fmt.Errorf("newOptions: %w", err)
	}
	return opts.run()
}
