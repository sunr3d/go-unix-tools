package main

import (
	"log"
	"os"

	"go-unix-tools/internal/logic/myXargs"
)

func main() {
	if err := myXargs.Run(os.Args); err != nil {
		log.Fatalf("myXargs.Run: %v", err)
	}
}
