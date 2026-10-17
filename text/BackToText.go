package text

import (
	"strings"

	"go-reloaded/tokenizer"
)

func BacktoText(tokens []tokenizer.Token) string {
	var b strings.Builder

	for i, t := range tokens {
		if i > 0 {
			prev := tokens[i-1]
			needSpace := false

			if prev.Kind != "Newlines" && t.Kind != "Newlines" {
				switch {
				case t.Kind == "Puncts":
					needSpace = false

				case t.Kind == "Quotes":
					needSpace = t.PrecededBySpace

				case prev.Kind == "Puncts":
					needSpace = true

				default:
					needSpace = t.PrecededBySpace
				}
			}

			if needSpace {
				b.WriteByte(' ')
			}
		}
		b.WriteString(t.Value)
	}
	return b.String()
}
