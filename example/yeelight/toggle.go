package main

import (
	"log"

	"cmd/main.go/internal/bulb/controller"
)

func main() {
	bulbController := controller.NewYeeLight()
	_, err := bulbController.Toggle("192.168.0.42:55443")
	if err != nil {
		log.Fatal(err)
	}

	return
}
