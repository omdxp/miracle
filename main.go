package miracle

import (
	"path/filepath"
	"runtime"
)

var f string

func init() {
	_, file, _, _ := runtime.Caller(0)
	f = file
}

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
	case "en", "ar":
		q.Language = l
		return nil
	}
	return &NotSupportedLanguageError{}
}

// ReadSurah read a surah by its number in the Quran
func (q *Quran) ReadSurah(n int) (Surah, error) {
	return Surah{}, nil
}

// GetSurahInfo gets surah information by its number in the Quran
func (q *Quran) GetSurahInfo(n uint) (*SurahInfo, error) {
	if n < 1 || n > 114 {
		return nil, &SurahNumberError{}
	}
	dirPath := filepath.Dir(f)
	bytes, err := openJsonFile(filepath.Join(dirPath, "data/surah.json"))
	if err != nil {
		return nil, err
	}
	suar, err := unmarshalSurahInfoJson(bytes)
	if err != nil {
		return nil, err
	}
	surah := (*suar)[n-1]
	return &surah, nil
}
