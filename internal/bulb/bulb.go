package bulb

import "cmd/main.go/internal/bulb/controller"

type Toggler interface {
	Toggle(loc string) (controller.Response, error)
}

type BrightnessSetter interface {
	SetBrightness(loc string, brightness, duration int) (controller.Response, error)
}
