<h1 align="center">miracle</h1>
<p align="center">Read Quran and Get Info about Suar and Ajzaa</p>

[![Build Status](https://github.com/Omar-Belghaouti/miracle/workflows/Go%20test/badge.svg)](https://github.com/Omar-Belghaouti/miracle/actions?query=branch%3Amain)

## Installation

- Install the gomodule package with:

```bash
go get -u github.com/Omar-Belghaouti/miracle
```

- Or import the package in your project:

```go
import "github.com/Omar-Belghaouti/miracle"
```

And run:

```bash
go mod tidy
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/Omar-Belghaouti/miracle"
	"github.com/Omar-Belghaouti/miracle/ajzaa"
	"github.com/Omar-Belghaouti/miracle/suar"
)

func main() {
	quran := miracle.Book()
	quran.SetLanguage("en")

	surah, err := quran.ReadSurah(suar.AAL_IMRAN)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(surah.Name)
	for k, v := range surah.Verse {
		fmt.Printf("%s: %s\n", k, v)
	}

	juz, err := quran.GetJuzInfo(ajzaa.JUZ_19)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(juz.Start)

	surahInfo, err := quran.GetSurahInfo(suar.AL_ADIYAT)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(surahInfo.Title)
}
```

## Documentation

Check the documentation [website](https://omar-belghaouti.github.io/miracle-docs/) for more information.
