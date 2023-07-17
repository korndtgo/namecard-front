package main

import (
	"fmt"
	"log"
	"namecard-gateway/internal/container"
	"os"
)

func main() {
	c := container.NewContainer()

	cmd := os.Args[1]
	if err := c.Container.Invoke(func() {
		log.Printf("Running %v...", cmd)
		// ctx := metadata.NewIncomingContext(context.Background(),
		// 	metadata.New(map[string]string{
		// 		"txid":        uuid.New().String(),
		// 		"userid":      "system",
		// 		"deviceid":    "system",
		// 		"deviceos":    "system",
		// 		"devicemodel": "system",
		// 	}))
		switch cmd {
		case "cmd-1":

		case "cmd-2":

		case "cmd-3":

		case "cmd-4":

		case "cmd-5":

		}
	}); err != nil {
		log.Println(fmt.Errorf("campaignctl Failed: %v", err.Error()))
	}
}
