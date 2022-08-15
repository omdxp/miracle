package miracle

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	f       string
	dirPath string
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	f = file
	dirPath = filepath.Dir(f)
}

type Quran struct {
	language string
}

// Book returns a pointer to Quran struct
func Book() *Quran {
	return &Quran{
		language: "ar",
	}
}

// SetLanguage sets the language for Quran struct
// supported languages are: "ar" and "en"
func (q *Quran) SetLanguage(l string) error {
	switch l {
	case "en", "ar":
		q.language = l
		return nil
	}
	return &NotSupportedLanguageError{}
}

// ReadSurah read a surah by its number in the Quran
func (q *Quran) ReadSurah(n uint) (*Surah, error) {
	if n < 1 || n > 114 {
		return nil, &SurahNumberError{}
	}
	var file string
	var parentDir string
	switch q.language {
	case "ar":
		file = fmt.Sprintf("surah_%d.json", n)
		parentDir = "data/surah/"
	case "en":
		file = fmt.Sprintf("en_translation_%d.json", n)
		parentDir = "data/translation/en/"
	}

	bytes, err := openJsonFile(filepath.Join(dirPath, parentDir, file))
	if err != nil {
		return nil, err
	}
	surah, err := unmarshalSurahJson(bytes)
	if err != nil {
		return nil, err
	}
	return surah, nil
}

// GetSurahInfo gets surah information by its number in the Quran
func (q *Quran) GetSurahInfo(n uint) (*SurahInfo, error) {
	if n < 1 || n > 114 {
		return nil, &SurahNumberError{}
	}
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

// GetJuzInfo gets juz information by its number in the Quran
func (q *Quran) GetJuzInfo(n uint) (*Juz, error) {
	if n < 1 || n > 30 {
		return nil, &JuzNumberError{}
	}
	bytes, err := openJsonFile(filepath.Join(dirPath, "data/juz.json"))
	if err != nil {
		return nil, err
	}
	ajzaa, err := unmarshalJuzInfoJson(bytes)
	if err != nil {
		return nil, err
	}
	juz := (*ajzaa)[n-1]
	return &juz, nil
}
