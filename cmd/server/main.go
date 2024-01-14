package main

import (
	"context"
	"fmt"

	"github.com/kanishkanaik/go-rest-api-course/cmd/Internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to Migrate Databases")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(context.Background(),
		comment.Comment{
			ID:     "0247c30c-9b9f-4192-8f74-5a7660308552",
			Author: "Kanishka",
			Body:   "First Comment",
			Slug:   "akjhsdvbcjhgasdvchjgasvdc",
		})

	fmt.Println(cmtService.GetComment(context.Background(), "1925b445-1c63-417f-b012-d7a27362791e"))

	fmt.Println("Succefully Connected and pinged to DB")

	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
