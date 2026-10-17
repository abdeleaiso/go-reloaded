package format

import (
	"math/big"
	"strings"
	"unicode"

	"go-reloaded/tokenizer"
)

func HexToDec(s string) string {
	n, ok := new(big.Int).SetString(s, 16)
	if !ok {
		return s
	}
	return n.String()
}

func BinToDec(s string) string {
	n, ok := new(big.Int).SetString(s, 2)
	if !ok {
		return s
	}
	return n.String()
}

func CapStr(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func TransformWord(tag, word string) string {
	switch tag {
	case "hex":
		return HexToDec(word)
	case "bin":
		return BinToDec(word)
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	case "cap":
		return CapStr(word)
	}
	return word
}

func ApplyTagToPrev(out []tokenizer.Token, tag tokenizer.Token) []tokenizer.Token {
	if tag.Count != nil && tag.Count.Sign() < 0 {
		return out
	}

	count := big.NewInt(1)
	if tag.Count != nil {
		count = tag.Count
	}

	applied := big.NewInt(0)
	for i := len(out) - 1; i >= 0 && applied.Cmp(count) < 0; i-- {
		if out[i].Kind == "Newlines" {
			break
		}
		if out[i].Kind != "Words" {
			continue
		}
		out[i].Value = TransformWord(tag.Value, out[i].Value)
		applied.Add(applied, big.NewInt(1))
	}
	return out
}

func ApplyTags(tokens []tokenizer.Token) []tokenizer.Token {
	out := make([]tokenizer.Token, 0, len(tokens))
	for _, t := range tokens {
		if t.Kind == "Tags" {
			out = ApplyTagToPrev(out, t)
			continue
		}
		out = append(out, t)
	}
	return out
}
