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
	quran.GetSurahInfo(suar.AL_AALA)
}
