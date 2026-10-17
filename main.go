package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go-reloaded/process"
	"go-reloaded/text"
	"go-reloaded/tokenizer"
)

func main() {
	f := flag.Bool("f", false, "use f strategy (default: left-to-right)")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Error: project accepts exactly 2 files")
		os.Exit(1)
	}

	file1 := args[0]
	file2 := args[1]

	if !strings.HasSuffix(file1, ".txt") || !strings.HasSuffix(file2, ".txt") {
		fmt.Println("Error: files must end with .txt")
		os.Exit(1)
	}

	var strategy process.Strategy = process.LeftToRight
	if *f {
		strategy = process.Pipeline
	}

	data, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	tokens := tokenizer.Tokenized(string(data))
	tokens = strategy(tokens)
	result := text.BacktoText(tokens)

	if err := os.WriteFile(file2, []byte(result), 0644); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
