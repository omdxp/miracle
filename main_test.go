package miracle

import (
	"testing"

	"github.com/Omar-Belghaouti/miracle/ajzaa"
	"github.com/Omar-Belghaouti/miracle/suar"
)

func TestSetLanguage(t *testing.T) {
	quran := Book()
	tests := []struct {
		l        string
		expected error
	}{
		{
			l:        "ar",
			expected: nil,
		},
		{
			l:        "en",
			expected: nil,
		},
		{
			l:        "zh",
			expected: &NotSupportedLanguageError{},
		},
	}
	for _, tt := range tests {
		err := quran.SetLanguage(tt.l)
		if err != tt.expected {
			t.Fatalf("expected %v, got %v", tt.expected, err)
		}
	}
}

func TestGetSurahInfo(t *testing.T) {
	quran := Book()
	tests := []struct {
		n        uint
		expected struct {
			surah *SurahInfo
			err   error
		}
	}{
		{n: suar.AL_FATIHA, expected: struct {
			surah *SurahInfo
			err   error
		}{
			surah: &SurahInfo{
				Place:   "Mecca",
				Type:    "Makkiyah",
				Count:   7,
				Title:   "Al-Fatiha",
				TitleAr: "الفاتحة",
				Index:   "001",
				Pages:   "1",
				Juz: []JuzInSurah{
					{
						Index: "01",
						Verse: Verse{
							Start: "verse_1",
							End:   "verse_7",
						},
					},
				},
			},
		}},
		{n: 255, expected: struct {
			surah *SurahInfo
			err   error
		}{
			surah: nil,
			err:   &SurahNumberError{},
		}},
	}
	for _, tt := range tests {
		surah, err := quran.GetSurahInfo(tt.n)
		if surah != nil && surah.Title != tt.expected.surah.Title {
			t.Fatalf("expected %v, got %v", tt.expected.surah, surah)
		}
		if err != tt.expected.err {
			t.Fatalf("expected %v, got %v", tt.expected.err, err)
		}
	}
}

func TestReadSurah(t *testing.T) {
	quran := Book()
	tests := []struct {
		n        uint
		expected struct {
			surah *Surah
			err   error
		}
	}{
		{
			n: suar.AL_FAJR,
			expected: struct {
				surah *Surah
				err   error
			}{
				surah: &Surah{
					Index: "089",
				},
			},
		},
		{
			n: 255,
			expected: struct {
				surah *Surah
				err   error
			}{
				surah: nil,
				err:   &SurahNumberError{},
			},
		},
	}

	for _, tt := range tests {
		surah, err := quran.ReadSurah(tt.n)
		if surah != nil && surah.Index != tt.expected.surah.Index {
			t.Fatalf("expected %v, got %v", tt.expected.surah, surah)
		}
		if err != tt.expected.err {
			t.Fatalf("expected %v, got %v", tt.expected.err, err)
		}
	}
}

func TestGetJuzInfo(t *testing.T) {
	quran := Book()
	tests := []struct {
		n        uint
		expected struct {
			juz *Juz
			err error
		}
	}{
		{
			n: ajzaa.JUZ_1,
			expected: struct {
				juz *Juz
				err error
			}{
				juz: &Juz{
					Index: "01",
				},
				err: nil,
			},
		},
		{
			n: 255,
			expected: struct {
				juz *Juz
				err error
			}{
				juz: nil,
				err: &JuzNumberError{},
			},
		},
	}
	for _, tt := range tests {
		juz, err := quran.GetJuzInfo(tt.n)
		if juz != nil && juz.Index != tt.expected.juz.Index {
			t.Fatalf("expected %v, got %v", tt.expected.juz.Index, juz.Index)
		}
		if err != tt.expected.err {
			t.Fatalf("expected %v, got %v", tt.expected.err, err)
		}
	}
}
