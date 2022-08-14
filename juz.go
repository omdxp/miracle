package miracle

type JuzPosition struct {
	Index uint   `json:"index"`
	Verse string `json:"verse"`
	Name  string `json:"name"`
}

type Juz struct {
	Index uint        `json:"index"`
	Start JuzPosition `json:"start"`
	End   JuzPosition `json:"end"`
}
