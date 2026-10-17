package format

import "go-reloaded/tokenizer"

func IsVowelOrH(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'h',
		'A', 'E', 'I', 'O', 'U', 'H':
		return true
	}
	return false
}

func FixArticles(tokens []tokenizer.Token) []tokenizer.Token {
	for i := range tokens {
		t := &tokens[i]
		if t.Kind != "Words" || (t.Value != "a" && t.Value != "A") {
			continue
		}

		j := i + 1
		for j < len(tokens) &&
			(tokens[j].Kind == "Quotes" ||
				tokens[j].Kind == "Puncts" ||
				tokens[j].Kind == "Others") {
			j++
		}
		if j >= len(tokens) || tokens[j].Kind != "Words" || tokens[j].Value == "" {
			continue
		}
		first := rune(tokens[j].Value[0])
		if !IsVowelOrH(first) {
			continue
		}
		if t.Value == "a" {
			t.Value = "an"
		} else {
			t.Value = "An"
		}
	}
	return tokens
}
