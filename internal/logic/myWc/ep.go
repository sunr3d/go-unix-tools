package myWc

import "fmt"

func Run(args []string) error {
	p, err := newProcessor(args)
	if err != nil {
		return fmt.Errorf("newProcessor: %w", err)
	}
	return p.run()
}
