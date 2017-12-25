package main

import "fmt"
import "app/gateway"

func main() {
	fmt.Println("Running...")
	controller := gateway.NewTweeterController()
	controller.Index()
}
