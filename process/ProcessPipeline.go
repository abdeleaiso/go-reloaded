package process

import (
	"go-reloaded/format"
	"go-reloaded/tokenizer"
)

func Pipeline(tokens []tokenizer.Token) []tokenizer.Token {
	tokens = format.ApplyTags(tokens)
	tokens = format.FixArticles(tokens)
	tokens = format.FixQuotes(tokens)
	tokens = format.FixPunctuation(tokens)
	return tokens
}
