package main

import (
	"log"
	"os"

	"Go_Day02/internal/logic/myWc"
)

func main() {
	if err := myWc.Run(os.Args[1:]); err != nil {
		log.Fatalf("myWc.Run: %v", err)
	}
}
