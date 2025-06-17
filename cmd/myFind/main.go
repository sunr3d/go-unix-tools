package main

import (
	"log"
	"os"

	"go-unix-tools/internal/logic/myFind"
)

func main() {
	if err := myFind.Run(os.Args[1:]); err != nil {
		log.Fatalf("myFind.Run: %v", err)
	}
}
