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
