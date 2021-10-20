package main

import (
	"log"

	"github.com/xvbnm48/grpc-golang/internal/db"
	"github.com/xvbnm48/grpc-golang/internal/rocket"
)

func Run() error {
	//responsible for initializing and starting
	// out grpc
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	_ = rocket.New(rocketStore)
	
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
