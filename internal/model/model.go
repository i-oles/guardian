package model

const (
	whiteType bulbType = "white"
	rgbType   bulbType = "rgb"
	On        State    = "on"
	Off       State    = "off"
	Offline   State    = "offline"
)

type bulbType string
type State string

type Bulb struct {
	ID         string
	Name       string
	BulbType   bulbType
	Brightness int
	Red        *int
	Green      *int
	Blue       *int
}

type BulbState struct {
	ID         string
	Name       string
	Location   string
	State      State
	Brightness int
}

type Preset struct {
	ID         int
	Name       string
	BulbID     string
	Brightness int
}
