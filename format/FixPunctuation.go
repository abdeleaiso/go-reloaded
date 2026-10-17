package format

import "go-reloaded/tokenizer"

func FixPunctuation(tokens []tokenizer.Token) []tokenizer.Token {
	for i := range tokens {
		if tokens[i].Kind != "Puncts" {
			continue
		}
		if i > 0 {
			tokens[i].PrecededBySpace = false
		}
		if i+1 < len(tokens) &&
			tokens[i+1].Kind != "Puncts" &&
			tokens[i+1].Kind != "Newlines" &&
			tokens[i+1].Kind != "Quotes" {
			tokens[i+1].PrecededBySpace = true
		}
	}
	return tokens
}
