package tokenizer

import (
	"math/big"
	"regexp"
	"unicode"
)

type Token struct {
	Kind            string
	Value           string
	Count           *big.Int
	PrecededBySpace bool
}

//regex to match only (tags) & (tags, n) 
var CommandRe = regexp.MustCompile(`^\(\s*(hex|bin|up|low|cap)\s*(,\s*(-?\d+)\s*)?\)`)

func IsWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-'
}

func IsPunctRune(r rune) bool {
	switch r {
	case '.', ',', '!', '?', ':', ';':
		return true
	}
	return false
}

func IsSpace(r rune) bool {
	switch r {
	case ' ', '\t', '\r':
		return true
	}
	return false
}

func Tokenized(s string) []Token {
	var tokens []Token
	runes := []rune(s)
	space := false

	emit := func(t Token) {
		t.PrecededBySpace = space
		tokens = append(tokens, t)
		space = false
	}

	for i := 0; i < len(runes); {
		r := runes[i]
		switch {
		case r == '\n':
			emit(Token{Kind: "Newlines", Value: "\n"})
			i++

		case IsSpace(r):
			space = true
			i++

		case r == '(':
			rest := string(runes[i:])
			tag := CommandRe.FindStringSubmatch(rest)
			if tag == nil {
				start := i
				for i < len(runes) && runes[i] != ')' && runes[i] != '\n' {
					i++
				}
				if i < len(runes) && runes[i] == ')' {
					i++
				}
				emit(Token{Kind: "Others", Value: string(runes[start:i])})
				continue
			}
			name := tag[1]
			count := big.NewInt(1)
			if tag[3] != "" {
				n, _ := new(big.Int).SetString(tag[3], 10)
				count = n
			}
			i += len([]rune(tag[0]))
			emit(Token{Kind: "Tags", Value: name, Count: count})

		case r == '\'':
			emit(Token{Kind: "Quotes", Value: "'"})
			i++

		case IsPunctRune(r):
			start := i
			for i < len(runes) && IsPunctRune(runes[i]) {
				i++
			}
			emit(Token{Kind: "Puncts", Value: string(runes[start:i])})

		case IsWordRune(r):
			start := i
			for i < len(runes) {
				if IsWordRune(runes[i]) {
					i++
					continue
				}
				if runes[i] == '\'' && i+1 < len(runes) && IsWordRune(runes[i+1]) {
					i++
					continue
				}
				break
			}
			emit(Token{Kind: "Words", Value: string(runes[start:i])})

		default:
			start := i
			for i < len(runes) &&
				!IsWordRune(runes[i]) && !IsPunctRune(runes[i]) &&
				runes[i] != '(' && runes[i] != '\'' &&
				runes[i] != '\n' && !IsSpace(runes[i]) {
				i++
			}
			emit(Token{Kind: "Others", Value: string(runes[start:i])})
		}
	}
	return tokens
}
