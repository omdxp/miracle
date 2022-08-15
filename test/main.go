package main

import (
	"fmt"

	"github.com/Omar-Belghaouti/miracle"
	"github.com/Omar-Belghaouti/miracle/ajzaa"
	"github.com/Omar-Belghaouti/miracle/suar"
)

func main() {
	quran := miracle.Book()
	err := quran.SetLanguage("en")
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err := quran.ReadSurah(suar.AL_FATIHA)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("count:", s.Count)
	fmt.Println("indx:", s.Index)
	fmt.Println("name:", s.Name)
	for k, v := range s.Verse {
		fmt.Println(k, v)
	}
	j, err := quran.GetJuzInfo(ajzaa.JUZ_30)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v", j)
}
