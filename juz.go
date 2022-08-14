package miracle

type JuzPosition struct {
	Index string `json:"index"`
	Verse string `json:"verse"`
	Name  string `json:"name"`
}

type Juz struct {
	Index string      `json:"index"`
	Start JuzPosition `json:"start"`
	End   JuzPosition `json:"end"`
}
