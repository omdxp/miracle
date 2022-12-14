package miracle

type Surah struct {
	Index string            `json:"index"`
	Name  string            `json:"name"`
	Verse map[string]string `json:"verse"`
	Count uint              `json:"count"`
	Juz   []JuzInSurah      `json:"juz"`
}

type Verse struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type JuzInSurah struct {
	Index string `json:"index"`
	Verse `json:"verse"`
}

type SurahInfo struct {
	Place   string       `json:"place"`
	Type    string       `json:"type"`
	Count   uint         `json:"count"`
	Title   string       `json:"title"`
	TitleAr string       `json:"titleAr"`
	Index   string       `json:"index"`
	Pages   string       `json:"pages"`
	Juz     []JuzInSurah `json:"juz"`
}
