package main

import (
	"log"

	"cmd/main.go/internal/bulb/controller"
)

func main() {
	light := controller.NewYeeLight()
	_, err := light.SetBrightness(
		"192.168.0.41:55443",
		10,
		1000,
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}
