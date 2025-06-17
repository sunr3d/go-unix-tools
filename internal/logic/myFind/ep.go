package myFind

import "fmt"

func Run(args []string) error {
	f, err := newFinder(args)
	if err != nil {
		return fmt.Errorf("newFinder: %w", err)
	}
	return f.run()
}
