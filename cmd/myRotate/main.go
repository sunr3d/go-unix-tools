package main

import (
	"log"
	"os"

	"go-unix-tools/internal/logic/myRotate"
)

func main() {
	if err := myRotate.Run(os.Args[1:]); err != nil {
		log.Fatalf("myRotate.Run: %v", err)
	}
}
