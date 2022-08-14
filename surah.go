package miracle

type Surah struct {
	Index uint       `json:"index"`
	Name  string     `json:"name"`
	Verse any        `json:"verse"`
	Count uint       `json:"count"`
	Juz   JuzInSurah `json:"juz"`
}

type Verse struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type JuzInSurah struct {
	Index uint `json:"index"`
	Verse `json:"verse"`
}

type SurahInfo struct {
	Place   string       `json:"place"`
	Type    string       `json:"type"`
	Count   uint         `json:"count"`
	Title   string       `json:"title"`
	TitleAr string       `json:"titleAr"`
	Index   uint         `json:"index"`
	Pages   uint         `json:"pages"`
	Juz     []JuzInSurah `json:"juz"`
}

const (
	AL_FATIHA = iota + 1
	AL_BAQARA
	AAL_IMRAN
	AL_NISAA
	AL_MAIDA
	AL_ANAAM
	AL_AARAF
	AL_ANFAL
	AL_TAWBA
	YUNUS
	HUD
	YUSUF
	AL_RAAD
	IBRAHIM
	AL_HIJR
	AL_NAHL
	AL_ISRAA
	AL_KAHF
	MARYAM
	TAHA
	AL_ANBIYA
	AL_HAJJ
	AL_MUMINUN
	AL_NUR
	AL_FURQAN
	AL_SHUARA
	AL_NAML
	AL_QASAS
	AL_ANKABUT
	AL_RUM
	LUQMAN
	AL_SAJDAH
	AL_AHZAB
	SABA
	FATIR
	YASIN
	AL_SAFFAT
	SAD
	AL_ZUMAR
	GHAFIR
	FUSSILAT
	AL_SHURA
	AL_ZUKHRUF
	AL_DUKHAN
	AL_JATHIA
	AL_AHQAF
	MUHAMMAD
	AL_FATH
	AL_HUJURAT
	QAF
	AL_DHARIYAT
	AL_TUR
	AL_NAJM
	AL_QAMAR
	AL_RAHMAN
	AL_WAQIA
	AL_HADID
	AL_MUJADILAH
	AL_HASHR
	AL_MUMTAHINAH
	AL_SAFF
	AL_JUMUAH
	AL_MUNAFIQUN
	AL_TAGHABUN
	AL_TALAQ
	AL_TAHRIM
	AL_MULK
	AL_QALAM
	AL_HAQQAH
	AL_MAARIJ
	NUH
	AL_JINN
	AL_MUZZAMMIL
	AL_MUDDATHIR
	AL_QIYAMAH
	AL_INSAN
	AL_MUSALAT
	AL_NABAA
	AL_NAZIAT
	ABASA
	AL_TAKWIR
	AL_INFITAR
	AL_MUTAFIFIN
	AL_INSHIQAQ
	AL_BURUJ
	AL_TARIQ
	AL_AALA
	AL_GHASHIYAH
	AL_FAJR
	AL_BALAD
	AL_SHAMS
	AL_LAIL
	AL_DUHA
	AL_SHARH
	AL_TIN
	AL_ALAQ
	AL_QADR
	AL_BAYINAH
	AL_ZALZALAH
	AL_ADIYAT
	AL_QARIAH
	AL_TAKATHUR
	AL_ASR
	AL_HUMAZAH
	AL_FIL
	QURAISH
	AL_MAUN
	AL_KAUTHAR
	AL_KAFIRUN
	AL_NASR
	AL_MASAD
	AL_IKHLAS
	AL_FALAQ
	AL_NAS
)
