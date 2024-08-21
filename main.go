package main

import (
	"log"
	"os"

	"github.com/lvm/album2go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
