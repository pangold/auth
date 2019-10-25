package main

import (
	"./controller"
)

func main() {
	rounter := controller.New()
	rounter.Run(":8080")
}
