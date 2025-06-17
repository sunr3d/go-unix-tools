package myRotate

import "fmt"

func Run(args []string) error {
	ar, err := newArchiver(args)
	if err != nil {
		return fmt.Errorf("newArchiver: %w", err)
	}
	return ar.run()
}
