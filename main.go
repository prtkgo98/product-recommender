package main

import (
	"log"

	"./server"
)

func main() {
	s := server.NewServer()
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
