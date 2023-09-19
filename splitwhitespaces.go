package piscine

func SplitWhiteSpaces(s string) []string {
	words := make([]string, 0)
	currentWord := ""
	for _, char := range s {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	return words
}
