package main

import (
	"log"
	"os"

	"go-unix-tools/internal/logic/myWc"
)

func main() {
	if err := myWc.Run(os.Args[1:]); err != nil {
		log.Fatalf("myWc.Run: %v", err)
	}
}
