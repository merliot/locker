package main

import (
	"log"
	"os"

	"github.com/merliot/locker"
)

//go:generate go run main.go
func main() {
	locker := locker.New("proto", "locker", "proto").(*locker.Locker)
	if err := locker.GenerateUf2s("../.."); err != nil {
		log.Println("Error generating UF2s:", err)
		os.Exit(1)
	}
}
