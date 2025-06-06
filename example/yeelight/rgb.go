package main

import (
	"log"

	"cmd/main.go/internal/repository/ylight"
)

func main() {
	light := ylight.NewYLight()
	_, err := light.SetRGB(
		"192.168.0.41:55443",
		1,
		2,
		3,
		3000,
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}
