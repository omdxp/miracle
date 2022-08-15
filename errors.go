package miracle

type SurahNumberError struct{}

func (s *SurahNumberError) Error() string {
	return "surah number is not between 1 and 114"
}

type NotSupportedLanguageError struct{}

func (n *NotSupportedLanguageError) Error() string {
	return "not supported language"
}

type JuzNumberError struct{}

func (j *JuzNumberError) Error() string {
	return "juz number is not between 1 and 30"
}
