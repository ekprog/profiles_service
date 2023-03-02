package main

import (
	"log"
	"profile_service/bootstrap"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
