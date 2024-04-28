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

type BulbState struct {
	ID       string
	Name     string
	Location string
	IsOn     bool
}

type Preset struct {
	ID        int
	Name      string
	BulbID    string
	Luminance int
}
