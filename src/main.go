package main

import (
	"github.com/lvm/album2go/src/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
