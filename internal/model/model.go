package model

const (
	whiteType bulbType = "white"
	//TODO: check this type
	rgbType bulbType = "rgb"
)

type bulbType string

type Bulb struct {
	ID        string
	IP        string
	BulbType  bulbType
	Luminance int
	Red       *int
	Green     *int
	Blue      *int
}

type Preset struct {
	ID        int
	Name      string
	BulbID    string
	Luminance int
}
