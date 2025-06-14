package main

import (
	"log"

	"cmd/main.go/internal/bulb/controller"
)

func main() {
	bulbController := controller.NewYeeLight()
	_, err := bulbController.SetRGB(
		"192.168.0.41:55443",
		210,
		1,
		1,
		5000,
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}
