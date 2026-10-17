# Go Reloaded

## Description

A text auto-correction tool written in go. it reads a .txt file applies a series of text transformations, and writes the corrected result to another .txt file

The project supports two processing strategies that produce output through different traversal approaches:

- **Left-to-Right:** the default. walks the token stream once, applying each rule as its token is encountered ("first come, first delivered").

- **Pipline:** runs each transformation as a seperate stage over the whole token slice.

## Features
_______________________________________________________
|       Rule                                          |
|-----------------------------------------------------|
| (hex) -> decimal [previous decimal]                 |
| (bin) -> decimal [previous decimal]                 |
| (up)  -> UPPERCASE previous word                    |
| (low) -> lowercase previous word                    |
| (cap) ->  Capitalize previous word                  |
| (cap|low|cap, n) -> apply to previous n words       |
| Punctuation hugs the previous word                  |
| Grouped punctuation stays together                  |
| Quotes hug their content                            |
-------------------------------------------------------

## Project Structure

```bash

module/
├── go.mod
├── main.go # CLI entry point, flag parsing, file I/O
│
├── tokenizer/
│ └── tokenizer.go # lexical analysis: text → []Token
│
├── format/
│ ├── apply_tags.go # (hex), (bin), (up), (low), (cap) transformations
│ ├── fix_article.go # a → an
│ ├── fix_punctuation.go # punctuation spacing
│ └── fix_quote.go # quote spacing
│
├── process/
│ ├── process.go # Strategy type
│ ├── pipeline.go # Strategy A: stage-by-stage
│ └── left_to_right.go # Strategy B: single-pass traversal
│
└── text/
└── back_to_text.go # tokens → final string
```

## Requirement

Requires **Go 1.24.4** or newer.

## Installation

```bash
git clone <go-reloaded repo>
cd go-reloaded
```

## Usage

### method 1:
```bash
# Build the program
go build -o go-reloaded

# Run it
./go-reloaded sample.txt result.txt

# Display the result
cat result.txt
```

### method 2:
```bash
go run . sample.txt result.txt
cat result.txt
```
## Testing

to run test
```bash
go test ./...
```

for more details
```bash
go test -v ./...
```
## Limitations

- The program only applies the text transformations explicitly supported by its implemented rules. It is not a general-purpose grammar or spell checker.

## Author
**Abdelkhaleq Elalaoui**
- **Github:** @abdeleaiso

## Licence
This project is licenced under the MIT.
See the [<a href="https://github.com/abdeleaiso/go-reloaded/tree/main?tab=MIT-1-ov-file">LICENCE</a>](LICENCE) file for details
