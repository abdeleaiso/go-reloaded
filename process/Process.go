package process

import "go-reloaded/tokenizer"

type Strategy func([]tokenizer.Token) []tokenizer.Token
