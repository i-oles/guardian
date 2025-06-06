package main

import (
	"log"

	"cmd/main.go/internal/repository/ylight"
)

func main() {
	light := ylight.NewYLight()
	_, err := light.PowerOff(
		"192.168.0.42:55443",
		2000,
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}
