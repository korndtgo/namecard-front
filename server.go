package main

import "card-service/internal/container"

func main() {
	c := container.NewContainer()

	if err := c.Run().Error; err != nil {
		panic(err)
	}
}
