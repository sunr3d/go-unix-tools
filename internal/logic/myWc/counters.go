package myWc

import (
	"bufio"
	"fmt"
	"io"
)

type Counter interface {
	count(r io.Reader) (int, error)
}

type lineCounter struct{}
type wordCounter struct{}
type charCounter struct{}

func (l *lineCounter) count(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func (w *wordCounter) count(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func (c *charCounter) count(r io.Reader) (int, error) {
	reader := bufio.NewReader(r)
	count := 0
	for {
		_, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("ошибка чтения: %w", err)
		}
		count += size
	}
	return count, nil
}
