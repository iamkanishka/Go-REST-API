package main

import (
	"context"
	"fmt"

	"github.com/kanishkanaik/go-rest-api-course/cmd/Internal/db"
)

// Run is responsible for instantiating and
// Startup of our Application
func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to Connect Database", err)

		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("Succefully Connected and pinged to DB")

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
