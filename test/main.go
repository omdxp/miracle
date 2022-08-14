package main

import (
	"fmt"

	"github.com/Omar-Belghaouti/miracle"
	"github.com/Omar-Belghaouti/miracle/suar"
)

func main() {
	quran := miracle.Book()
	err := quran.SetLanguage("en")
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err := quran.ReadSurah(suar.AAL_IMRAN)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s.Verse["verse_22"])
}
