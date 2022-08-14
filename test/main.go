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
	s, err := quran.ReadSurah(suar.AL_FAJR)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("count:", s.Count)
	fmt.Println("indx:", s.Index)
	fmt.Println("name:", s.Name)
}
