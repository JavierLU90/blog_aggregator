package main

import (
	"blog_aggregator/internal/config"
	"fmt"
	"log"
)

func main() {
	data, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", data)

	err = data.SetUser("javier")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	data, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", data)
}
