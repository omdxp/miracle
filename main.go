package miracle

import "errors"

type Quran struct {
	Language string
}

// Book returns a pointer to Quran struct
func Book() *Quran {
	return &Quran{
		Language: "ar",
	}
}

// SetLanguage sets the language for Quran struct
// supported languages are: "ar" and "en"
func (q *Quran) SetLanguage(l string) error {
	switch l {
	case "ar":
	case "en":
		q.Language = l
		return nil
	}
	return errors.New("not supported language")
}

// ReadSurah read a surah by its number in the Quran
func (q *Quran) ReadSurah(n int) (Surah, error) {
	return Surah{}, nil
}

// GetSurahInfo gets surah information by its number in the Quran
func (q *Quran) GetSurahInfo(n int) (SurahInfo, error) {
	return SurahInfo{}, nil
}
