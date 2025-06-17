package main

import (
	"log"
	"os"

	"Go_Day02/internal/logic/myFind"
)

func main() {
	if err := myFind.Run(os.Args[1:]); err != nil {
		log.Fatalf("myFind.Run: %v", err)
	}
}
