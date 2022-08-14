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
	surah, err := quran.GetSurahInfo(suar.AL_FATIHA)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", surah)
}
