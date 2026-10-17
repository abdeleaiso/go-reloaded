package format

import "go-reloaded/tokenizer"

func FixQuotes(tokens []tokenizer.Token) []tokenizer.Token {
	open := true
	for i := range tokens {
		if tokens[i].Kind != "Quotes" {
			continue
		}

		if open {
			if i > 0 && tokens[i-1].Kind == "Puncts" {
				tokens[i].PrecededBySpace = true
			} else {
				tokens[i].PrecededBySpace = false
			}

			if i+1 < len(tokens) {
				tokens[i+1].PrecededBySpace = false
			}
		} else {
			tokens[i].PrecededBySpace = false

			if i+1 < len(tokens) &&
				tokens[i+1].Kind != "Puncts" &&
				tokens[i+1].Kind != "Newlines" &&
				tokens[i+1].Kind != "Quotes" {
				tokens[i+1].PrecededBySpace = true
			}
		}
		open = !open
	}
	return tokens
}
