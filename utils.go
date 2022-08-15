package miracle

import (
	"encoding/json"
	"os"
)

// openJsonFile opens a local json file with the given name
func openJsonFile(n string) ([]byte, error) {
	bytes, err := os.ReadFile(n)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// unmarshalSurahInfoJson unmarshal the ./data/surah.json file
func unmarshalSurahInfoJson(b []byte) (*[]SurahInfo, error) {
	var suar *[]SurahInfo
	err := json.Unmarshal(b, &suar)
	if err != nil {
		return nil, err
	}
	return suar, nil
}

// unmarshalSurahJson unmarshal the ./data/surah/surah_X.json file
// (or ./data/translation/en/en_translation_1.json in case of en)
func unmarshalSurahJson(b []byte) (*Surah, error) {
	var surah *Surah
	err := json.Unmarshal(b, &surah)
	if err != nil {
		return nil, err
	}
	return surah, nil
}

// unmarshalJuzInfoJson unmarshal the ./data/juz.json file
func unmarshalJuzInfoJson(b []byte) (*[]Juz, error) {
	var ajzaa *[]Juz
	err := json.Unmarshal(b, &ajzaa)
	if err != nil {
		return nil, err
	}
	return ajzaa, nil
}
