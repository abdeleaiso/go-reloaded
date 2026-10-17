package process

import (
	"go-reloaded/format"
	"go-reloaded/tokenizer"
)

func LeftToRight(tokens []tokenizer.Token) []tokenizer.Token {
	out := make([]tokenizer.Token, 0, len(tokens))

	quoteOpen := true

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		switch t.Kind {

		case "Words":
			if (t.Value == "a" || t.Value == "A") && nextWordStartsWithVowelOrH(tokens, i) {
				if t.Value == "a" {
					t.Value = "an"
				} else {
					t.Value = "An"
				}
			}
			out = append(out, t)

		case "Tags":
			out = format.ApplyTagToPrev(out, t)

		case "Puncts":
			t.PrecededBySpace = false
			if len(out) > 0 && out[len(out)-1].Kind != "Newlines" {
				out[len(out)-1].Value += t.Value
				if i+1 < len(tokens) &&
					tokens[i+1].Kind != "Puncts" &&
					tokens[i+1].Kind != "Newlines" {
					tokens[i+1].PrecededBySpace = true
				}
			} else {
				out = append(out, t)
			}

		case "Quotes":
			if quoteOpen {
				out = append(out, t)
				if i+1 < len(tokens) {
					tokens[i+1].PrecededBySpace = false
				}
			} else {
				t.PrecededBySpace = false
				out = append(out, t)
				if i+1 < len(tokens) &&
					tokens[i+1].Kind != "Puncts" &&
					tokens[i+1].Kind != "Newlines" &&
					tokens[i+1].Kind != "Quotes" {
					tokens[i+1].PrecededBySpace = true
				}
			}
			quoteOpen = !quoteOpen

		default:
			out = append(out, t)
		}
	}
	return out
}

func nextWordStartsWithVowelOrH(tokens []tokenizer.Token, i int) bool {
	j := i + 1
	for j < len(tokens) &&
		(tokens[j].Kind == "Quotes" ||
			tokens[j].Kind == "Puncts" ||
			tokens[j].Kind == "Others") {
		j++
	}
	if j >= len(tokens) || tokens[j].Kind != "Words" || tokens[j].Value == "" {
		return false
	}
	return format.IsVowelOrH(rune(tokens[j].Value[0]))
}