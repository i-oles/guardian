package bulb

import "cmd/main.go/internal/bulb/controller"

type Controller interface {
	Toggle(loc string) (controller.Response, error)
	SetBrightness(loc string, brightness, duration int) (controller.Response, error)
	SetRGB(loc string, red, green, blue, duration int) (controller.Response, error)
	PowerOff(loc string, duration int) (controller.Response, error)
}
