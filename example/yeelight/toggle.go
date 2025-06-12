package main

import (
	"log"

	"cmd/main.go/internal/bulb/controller"
)

func main() {
	light := controller.NewYeeLight()
	_, err := light.Toggle("192.168.0.42:55443")
	if err != nil {
		log.Fatal(err)
	}

	return
}
