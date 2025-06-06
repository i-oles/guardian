package main

import (
	"log"

	"cmd/main.go/internal/repository/ylight"
)

func main() {
	light := ylight.NewYLight()
	_, err := light.Toggle("192.168.0.42:55443")
	if err != nil {
		log.Fatal(err)
	}

	return
}
