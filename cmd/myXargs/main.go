package main

import (
	"log"
	"os"

	"Go_Day02/internal/logic/myXargs"
)

func main() {
	if err := myXargs.Run(os.Args); err != nil {
		log.Fatalf("myXargs.Run: %v", err)
	}
}
