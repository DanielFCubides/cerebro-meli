package main

import (
	"cerebro/adapters"
	"log"
)

func main() {
	// server setup:
	err := adapters.ServerSetup().
		Run()
	if err != nil {
		log.Println(err)
	}
}
