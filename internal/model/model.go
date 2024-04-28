package model

const (
	whiteType bulbType = "white"
	//TODO: check this type
	rgbType bulbType = "rgb"
)

type bulbType string

type Bulb struct {
	ID        string
	Name      string
	BulbType  bulbType
	Luminance int
	Red       *int
	Green     *int
	Blue      *int
}

type State string

const (
	On      State = "on"
	Off     State = "off"
	Offline State = "offline"
)

type BulbState struct {
	ID       string
	Name     string
	Location string
	State    State
}

type Preset struct {
	ID        int
	Name      string
	BulbID    string
	Luminance int
}
