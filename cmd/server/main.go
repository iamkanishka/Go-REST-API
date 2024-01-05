package main

import "fmt"

//Run is responsible for instantiating and
//Startup of our Application
func Run() error {
	fmt.Println("Starting up our AAplication")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
