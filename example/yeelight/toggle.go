package main

import (
	"fmt"
	"github.com/akominch/yeelight"
)

func main() {
	err := toggleBulb("192.168.0.15")
	if err != nil {
		fmt.Println(err)
	}
}

func toggleBulb(ip string) error {
	config := yeelight.BulbConfig{
		Ip:     ip,
		Effect: yeelight.Smooth,
	}

	bulb := yeelight.New(config)
	_, err := bulb.TurnOn()
	if err != nil {
		return fmt.Errorf("could not turn off bulb: %s, err: %w", ip, err)
	}
	return nil
}
