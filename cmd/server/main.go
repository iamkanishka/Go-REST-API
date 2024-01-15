package main

import (
	"fmt"

	"github.com/kanishkanaik/go-rest-api-course/cmd/Internal/comment"
	transportHttp "github.com/kanishkanaik/go-rest-api-course/cmd/Internal/transport/http"

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

	// if err := db.Ping(context.Background()); err != nil {
	// 	return err
	// }
	fmt.Println("Succefully Connected and pinged to DB")

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to Migrate Databases")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
