package util

type Imagepacket struct {
	Name string       `json:"dump"`
	Img  [][]RawPixel `json:"img"`
}

type RawPixel struct {
	R uint32 `json:"R"`
	G uint32 `json:"G"`
	B uint32 `json:"B"`
	A uint32 `json:"A"`
}

func (t *RawPixel) Get() (uint32, uint32, uint32, uint32) {
	return t.R, t.G, t.B, t.A
}
