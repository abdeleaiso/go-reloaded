package main

import (
	"testing"

	"go-reloaded/process"
	"go-reloaded/text"
	"go-reloaded/tokenizer"
)

func run(s string, strat process.Strategy) string {
	return text.BacktoText(strat(tokenizer.Tokenized(s)))
}

func TestProcessText(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "project_test1",
			input:  "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			output: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:   "project_test2",
			input:  "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			output: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:   "project_test3",
			input:  "There is no greater agony than bearing a untold story inside you.",
			output: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:   "project_test4",
			input:  "Punctuation tests are ... kinda boring ,what do you think ?",
			output: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			name:   "audit_test1",
			input:  "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			output: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:   "audit_test2",
			input:  "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			output: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:   "audit_test3",
			input:  "Don not be sad ,because sad backwards is das . And das not good",
			output: "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:   "audit_test4",
			input:  "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			output: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			name:   "official capitalization example",
			input:  "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			output: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:   "hex and binary conversion",
			input:  "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			output: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:   "article correction",
			input:  "There is no greater agony than bearing a untold story inside you.",
			output: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:   "punctuation",
			input:  "Punctuation tests are ... kinda boring ,what do you think ?",
			output: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			name:   "lowercase count and uppercase count",
			input:  "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			output: "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:   "hexadecimal letters",
			input:  "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			output: "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:   "quotes, capitalization and articles",
			input:  "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			output: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
		{
			name:   "uppercase command with count",
			input:  "the fans shouted: WOW (up, 2)",
			output: "the fans SHOUTED: WOW",
		},
		{
			name:   "single uppercase command",
			input:  "Ready, set, go (up) !",
			output: "Ready, set, GO!",
		},
		{
			name:   "single lowercase command",
			input:  "I should stop SHOUTING (low)",
			output: "I should stop shouting",
		},
		{
			name:   "single capitalization command",
			input:  "Welcome to the Brooklyn bridge (cap)",
			output: "Welcome to the Brooklyn Bridge",
		},
		{
			name:   "quotation mark spacing",
			input:  "I am exactly how they describe me: ' awesome '",
			output: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:   "empty input",
			input:  "",
			output: "",
		},
		{
			name:   "negative hexadecimal",
			input:  "-1E (hex) files",
			output: "-30 files",
		},
		{
			name:   "negative binary",
			input:  "-101 (bin) files",
			output: "-5 files",
		},
		{
			name:   "hex followed by punctuation",
			input:  "1E (hex), files",
			output: "30, files",
		},
		{
			name:   "bin followed by punctuation",
			input:  "101 (bin)!",
			output: "5!",
		},
		{
			name:   "up followed by punctuation",
			input:  "hello (up)!",
			output: "HELLO!",
		},
		{
			name:   "low followed by punctuation",
			input:  "HELLO (low)!",
			output: "hello!",
		},
		{
			name:   "cap followed by punctuation",
			input:  "hello (cap)!",
			output: "Hello!",
		},
		{
			name:   "apostrophe",
			input:  "He said: ' hello world '",
			output: "He said: 'hello world'",
		},
		{
			name:   "a before vowel",
			input:  "This is a apple.",
			output: "This is an apple.",
		},
		{
			name:   "a before h",
			input:  "This is a house.",
			output: "This is an house.",
		},
		{
			name:   "Hex conversion",
			input:  "1E (hex) files were added",
			output: "30 files were added",
		},
		{
			name:   "Hex conversion lowercase/uppercase",
			input:  "Simply add 42 (hex) and ff (hex)",
			output: "Simply add 66 and 255",
		},
		{
			name:   "bin conversion",
			input:  "It has been 10 (bin) years",
			output: "It has been 2 years",
		},

		{
			name:   "Up multiple words",
			input:  "This is so exciting (up, 2)",
			output: "This is SO EXCITING",
		},
		{
			name:   "Low multiple words",
			input:  "IT WAS THE (low, 3) winter",
			output: "it was the winter",
		},
		{
			name:   "Cap multiple words",
			input:  "it was the age of foolishness (cap, 6)",
			output: "It Was The Age Of Foolishness",
		},

		{
			name:   "Basic punctuation spacing",
			input:  "I was sitting over there ,and then BAMM !!",
			output: "I was sitting over there, and then BAMM!!",
		},
		{
			name:   "Multiple grouped punctuations",
			input:  "I was thinking ... You were right",
			output: "I was thinking... You were right",
		},
		{
			name:   "Multiple words in single quotes",
			input:  "As Elton John said: ' I am the most well-known homosexual in the world '",
			output: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},

		{
			name:   "Article 'a' before vowel",
			input:  "There it was. A amazing rock!",
			output: "There it was. An amazing rock!",
		},
		{
			name:   "Article 'a' before 'h' sound",
			input:  "It was a hour ago",
			output: "It was an hour ago",
		},

		{
			name:   "Computed transformations",
			input:  "1a (hex) (up) items",
			output: "26 items",
		},

		{
			name:   "Computed conversions: bin then Hex then Up",
			input:  "10 (bin) (hex) (up) units",
			output: "2 units",
		},
		{
			name:   "Computed text modifiers: Low then Cap",
			input:  "HELLO (low) (cap) WORLD",
			output: "Hello WORLD",
		},
		{
			name:   "Multiple modifiers back-to-back with numbers",
			input:  "this is a crazy test (up, 2) (low, 3)",
			output: "this is a crazy test",
		},

		{
			name:   "Flag before comma",
			input:  "hello world (up) , welcome",
			output: "hello WORLD, welcome",
		},
		{
			name:   "Flag attached to punctuation",
			input:  "are you sure (up)?",
			output: "are you SURE?",
		},

		{
			name:   "Quotes with flag inside quotes",
			input:  "he said ' hello (cap) world ' !",
			output: "he said 'Hello world'!",
		},
		{
			name:   "Quotes with multi-word up and punctuation directly outside",
			input:  "she screamed : ' stop it now (up, 3) ' !",
			output: "she screamed: 'STOP IT NOW'!",
		},
		{
			name:   "Multiple pairs of single quotes in one sentence",
			input:  "the ' first ' and ' second ' options",
			output: "the 'first' and 'second' options",
		},
		{
			name:   "Quotes wrapping punctuation and words",
			input:  "I love ' apples , oranges , and bananas ! '",
			output: "I love 'apples, oranges, and bananas!'",
		},

		{
			name:   "Article 'a' modified by flag into a vowel word",
			input:  "a 10 (hex) fold increase",
			output: "a 16 fold increase",
		},
		{
			name:   "Article 'a' before case transformation",
			input:  "a owl (cap) flew by",
			output: "an Owl flew by",
		},
		{
			name:   "Capital 'A' to 'An' conversion before vowel",
			input:  "A unoutput error occurred",
			output: "An unoutput error occurred",
		},
		{
			name:   "Multiple articles in sentence with mixed case",
			input:  "A apple , a banana , and a orange",
			output: "An apple, a banana, and an orange",
		},

		{
			name:   "Mixed punctuation cluster like !? and ...",
			input:  "What !? Is that true ... ?",
			output: "What!? Is that true...?",
		},
		{
			name:   "Punctuation at the very start and end of string",
			input:  " ... start and end ! ",
			output: "... start and end!",
		},
		{
			name:   "Consecutive commas and colons",
			input:  "item1 , item2 : item3 ; item4 .",
			output: "item1, item2: item3; item4.",
		},

		{
			name:   "Modifier number larger than available words",
			input:  "one two (up, 10)",
			output: "ONE TWO",
		},
		{
			name:   "Modifier count skipping punctuation tokens",
			input:  "one , two . three (up, 3)",
			output: "ONE, TWO. THREE",
		},

		{
			name:   "Hex with letters A-F in mixed case",
			input:  "ff (hex) and 1A (hex) and 0 (hex)",
			output: "255 and 26 and 0",
		},
		{
			name:   "binary zero",
			input:  "0 (bin) or 1 (bin)",
			output: "0 or 1",
		},

		{name: "Hex zero padded", input: "000F (hex) items", output: "15 items"},
		{name: "Hex mixed case letters", input: "aBcDeF (hex) value", output: "11259375 value"},
		{name: "Hex maximum common byte value", input: "FF (hex) value", output: "255 value"},
		{name: "binary long number", input: "11111111 (bin) value", output: "255 value"},
		{name: "binary leading zeros", input: "00001010 (bin) value", output: "10 value"},
		{name: "binary zero with punctuation", input: "0 (bin), 1 (bin), 10 (bin)!", output: "0, 1, 2!"},

		{name: "Uppercase already uppercase", input: "HELLO (up) WORLD", output: "HELLO WORLD"},
		{name: "Lowercase already lowercase", input: "hello (low) WORLD", output: "hello WORLD"},
		{name: "Cap lowercase", input: "hello (cap) world", output: "Hello world"},
		{name: "Cap uppercase", input: "HELLO (cap) WORLD", output: "Hello WORLD"},
		{name: "Cap mixed case", input: "hELLo (cap) WORLD", output: "Hello WORLD"},
		{name: "Uppercase punctuation attached", input: "hello! (up)", output: "HELLO!"},
		{name: "Lowercase punctuation attached", input: "HELLO! (low)", output: "hello!"},
		{name: "Cap punctuation attached", input: "hello! (cap)", output: "Hello!"},

		{name: "Up exactly three words", input: "one two three (up, 3)", output: "ONE TWO THREE"},
		{name: "Low exactly four words", input: "ONE TWO THREE FOUR (low, 4)", output: "one two three four"},
		{name: "Cap exactly five words", input: "one two three four five (cap, 5)", output: "One Two Three Four Five"},
		{name: "Up count one", input: "hello world (up, 1)", output: "hello WORLD"},
		{name: "Low count one", input: "HELLO WORLD (low, 1)", output: "HELLO world"},
		{name: "Cap count one", input: "hello WORLD (cap, 1)", output: "hello World"},
		{name: "Up count greater than sentence", input: "hello world (up, 100)", output: "HELLO WORLD"},
		{name: "Low count greater than sentence", input: "HELLO WORLD (low, 100)", output: "hello world"},
		{name: "Cap count greater than sentence", input: "hello WORLD (cap, 100)", output: "Hello World"},
		{name: "Multiple sequential up operations", input: "one two three (up, 2) four (up)", output: "one TWO THREE FOUR"},
		{name: "Multiple sequential low operations", input: "ONE TWO THREE (low, 2) FOUR (low)", output: "ONE two three four"},
		{name: "Multiple sequential cap operations", input: "one two three (cap, 2) four (cap)", output: "one Two Three Four"},

		{name: "Hex then uppercase", input: "ff (hex) (up)", output: "255"},
		{name: "binary then uppercase", input: "1010 (bin) (up)", output: "10"},
		{name: "Hex then lowercase", input: "FF (hex) (low)", output: "255"},
		{name: "Hex followed by cap", input: "2A (hex) (cap)", output: "42"},
		{name: "Multiple conversions", input: "10 (bin) 1F (hex) 101 (bin)", output: "2 31 5"},
		{name: "Conversions mixed with case", input: "AA (hex) (up) BB (hex) (low)", output: "170 187"},
		{name: "Conversion inside punctuation", input: "Values: FF (hex), 10 (bin), 1A (hex)!", output: "Values: 255, 2, 26!"},

		{name: "Many commas", input: "one , two , three , four", output: "one, two, three, four"},
		{name: "Many periods", input: "one . two . three . four", output: "one. two. three. four"},
		{name: "Many exclamations", input: "hello ! ! ! world", output: "hello!!! world"},
		{name: "Many question marks", input: "what ? ? ? really", output: "what??? really"},
		{name: "Colon and semicolon", input: "first : second ; third", output: "first: second; third"},
		{name: "All supported punctuation", input: "a , b . c ! d ? e : f ; g", output: "a, b. c! d? e: f; g"},
		{name: "Punctuation groups", input: "Really !? What ... Why ?!", output: "Really!? What... Why?!"},
		{name: "Long punctuation group", input: "What !!!!!????? happened", output: "What!!!!!????? happened"},
		{name: "Mixed punctuation group", input: "Really !?!?!!...?? wow", output: "Really!?!?!!...?? wow"},
		{name: "Punctuation immediately after punctuation", input: "Hello , , , world", output: "Hello,,, world"},
		{name: "Punctuation with spaces everywhere", input: "hello , world ! how ? are you", output: "hello, world! how? are you"},
		{name: "Punctuation at beginning", input: "! hello , world", output: "! hello, world"},
		{name: "Multiple punctuation at beginning", input: "... hello !!! world", output: "... hello!!! world"},
		{name: "Multiple punctuation at end", input: "hello world . . .", output: "hello world..."},
		{name: "Punctuation around conversion", input: "FF (hex) , 10 (bin) !", output: "255, 2!"},

		{name: "Up before comma", input: "hello world (up) , next", output: "hello WORLD, next"},
		{name: "Up before punctuation cluster", input: "hello world (up) !?", output: "hello WORLD!?"},
		{name: "Low before punctuation", input: "HELLO WORLD (low) ... next", output: "HELLO world... next"},
		{name: "Cap before punctuation", input: "hello world (cap) !", output: "hello World!"},
		{name: "Up multi-word before punctuation", input: "one two three (up, 3) ... next", output: "ONE TWO THREE... next"},
		{name: "Low multi-word before punctuation", input: "ONE TWO THREE (low, 3) !!!", output: "one two three!!!"},
		{name: "Cap multi-word before punctuation", input: "one two three (cap, 3) !?", output: "One Two Three!?"},

		{name: "Empty spaces inside quotes", input: "' hello '", output: "'hello'"},
		{name: "Quotes with many words", input: "' one two three four five '", output: "'one two three four five'"},
		{name: "Two quote pairs", input: "' hello ' and ' goodbye '", output: "'hello' and 'goodbye'"},
		{name: "Three quote pairs", input: "' one ' ' two ' ' three '", output: "'one' 'two' 'three'"},
		{name: "Quotes with punctuation", input: "' hello ! '", output: "'hello!'"},
		{name: "Quotes with question", input: "' are you sure ? '", output: "'are you sure?'"},
		{name: "Quotes with punctuation cluster", input: "' really !? '", output: "'really!?'"},
		{name: "Quotes containing comma", input: "' hello , world '", output: "'hello, world'"},
		{name: "Quotes containing colon", input: "' first : second '", output: "'first: second'"},
		{name: "Quote + flag", input: "' hello (up) '", output: "'HELLO'"},
		{name: "Quote + low flag", input: "' HELLO (low) '", output: "'hello'"},
		{name: "Quote + cap flag", input: "' hello (cap) '", output: "'Hello'"},
		{name: "Quote + multi-word flag", input: "' one two three (up, 3) '", output: "'ONE TWO THREE'"},
		{name: "Quote with punctuation and flag", input: "' hello world (up, 2) ! '", output: "'HELLO WORLD!'"},
		{name: "Multiple quoted phrases with punctuation", input: "he said : ' hello ' , then she said : ' goodbye ' !", output: "he said: 'hello', then she said: 'goodbye'!"},

		{name: "A before e", input: "a elephant", output: "an elephant"},
		{name: "A before i", input: "a igloo", output: "an igloo"},
		{name: "A before o", input: "a orange", output: "an orange"},
		{name: "A before u", input: "a umbrella", output: "an umbrella"},
		{name: "A before uppercase vowel", input: "a Apple", output: "an Apple"},
		{name: "A before uppercase H", input: "a House", output: "an House"},
		{name: "Lowercase a before consonant", input: "a banana", output: "a banana"},
		{name: "Uppercase A before consonant", input: "A banana", output: "A banana"},
		{name: "Multiple articles", input: "a apple and a orange and a banana", output: "an apple and an orange and a banana"},
		{name: "Articles with punctuation", input: "a apple , a banana , a orange !", output: "an apple, a banana, an orange!"},
		{name: "Article after period", input: "Hello. a apple is here.", output: "Hello. an apple is here."},
		{name: "Article before quoted word", input: "a ' apple '", output: "an 'apple'"},

		{name: "A then cap vowel word", input: "a apple (cap)", output: "an Apple"},
		{name: "A then cap consonant word", input: "a banana (cap)", output: "a Banana"},
		{name: "A before converted hex number", input: "a FF (hex) value", output: "a 255 value"},
		{name: "A before converted binary number", input: "a 101 (bin) value", output: "a 5 value"},

		{name: "Unicode lowercase", input: "école (up)", output: "ÉCOLE"},
		{name: "Unicode uppercase", input: "ÉCOLE (low)", output: "école"},
		{name: "Unicode capitalization", input: "école (cap)", output: "École"},
		{name: "French Unicode sentence", input: "bonjour , ça va ?", output: "bonjour, ça va?"},
		{name: "Unicode words with punctuation", input: "café , thé ! gâteau ?", output: "café, thé! gâteau?"},
		{name: "Arabic text punctuation", input: "مرحبا , كيف حالك ?", output: "مرحبا, كيف حالك?"},
		{name: "Arabic text with quotes", input: "' مرحبا بالعالم ' !", output: "'مرحبا بالعالم'!"},
		{name: "Arabic mixed with English", input: "hello , مرحبا ! world ?", output: "hello, مرحبا! world?"},
		{name: "Unicode multi-word uppercase", input: "bonjour ça va (up, 3)", output: "BONJOUR ÇA VA"},
		{name: "Unicode multi-word lowercase", input: "BONJOUR ÇA VA (low, 3)", output: "bonjour ça va"},
		{name: "Unicode multi-word capitalization", input: "ça va bien (cap, 3)", output: "Ça Va Bien"},

		{name: "Everything combined", input: "hello (cap) world , FF (hex) (up) and 1010 (bin) !", output: "Hello world, 255 and 10!"},
		{name: "Quotes conversions punctuation articles", input: "a apple , ' hello world (up, 2) ' !", output: "an apple, 'HELLO WORLD'!"},
		{name: "Multiple flags conversions and punctuation", input: "this is a test (up, 4) , FF (hex) , 1010 (bin) !", output: "THIS IS A TEST, 255, 10!"},
		{name: "Long mixed sentence", input: "it was a amazing day , the weather was SO NICE (low) ! ' HELLO WORLD (up, 2) ' , and FF (hex) birds flew !", output: "it was an amazing day, the weather was SO nice! 'HELLO WORLD', and 255 birds flew!"},
		{name: "Multiple transformations around punctuation", input: "hello world (up, 2) , this is a apple (cap) !", output: "HELLO WORLD, this is an Apple!"},
		{name: "Complex quote sentence", input: "he said : ' this is a amazing TEST (low, 4) ! ' , then she said : ' FF (hex) is 1010 (bin) ' !", output: "he said: 'this is an amazing test!', then she said: '255 is 10'!"},
		{name: "Everything repeated", input: "a apple , FF (hex) , 1010 (bin) , hello (up) ! ' goodbye (cap) ' ?", output: "an apple, 255, 10, HELLO! 'Goodbye'?"},
		{name: "Heavy punctuation and flags", input: "hello !!! world (up) , how are you ??? ' fine (cap) ' !", output: "hello!!! WORLD, how are you??? 'Fine'!"},
		{name: "Heavy punctuation with conversions", input: "FF (hex) !!! 1010 (bin) ??? 10 (hex) ...", output: "255!!! 10??? 16..."},
		{name: "Long modifier chain", input: "hello world (up, 2) (low, 2) (cap, 2)", output: "Hello World"},
		{name: "Repeated modifier chain", input: "one two three (up, 3) (low, 2) (cap)", output: "ONE two Three"},

		{name: "Many spaces before punctuation", input: "hello , world ! how ?", output: "hello, world! how?"},
		{name: "Many spaces inside quotes", input: "' hello world '", output: "'hello world'"},
		{name: "Many spaces around flags", input: "hello (up) world", output: "HELLO world"},
		{name: "Many spaces around punctuation clusters", input: "hello !? world", output: "hello!? world"},

		{name: "Very long mixed sentence", input: "this is a amazing test of the text editor , FF (hex) was converted correctly ! then 1010 (bin) was converted too ... ' hello world (up, 2) ' and finally the quick brown fox (cap, 4) jumped over a orange dog !", output: "this is an amazing test of the text editor, 255 was converted correctly! then 10 was converted too... 'HELLO WORLD' and finally The Quick Brown Fox jumped over an orange dog!"},

		{
			name:   "up attached to comma",
			input:  "hello (up),",
			output: "HELLO,",
		},
		{
			name:   "hex attached to punctuation",
			input:  "1E (hex)!",
			output: "30!",
		},
		{
			name:   "bin attached to punctuation",
			input:  "10 (bin)!",
			output: "2!",
		},
		{
			name:   "hex attached to comma",
			input:  "1F (hex), files",
			output: "31, files",
		},
		{
			name:   "bin attached to comma",
			input:  "101 (bin), files",
			output: "5, files",
		},

		{
			name:   "punctuation before up",
			input:  "hello , (up)",
			output: "HELLO,",
		},
		{
			name:   "punctuation before counted up",
			input:  "hello world , (up, 2)",
			output: "HELLO WORLD,",
		},
		{
			name:   "punctuation before counted cap",
			input:  "hello world , (cap, 2)",
			output: "Hello World,",
		},
		{
			name:   "punctuation before counted low",
			input:  "HELLO WORLD , (low, 2)",
			output: "hello world,",
		},

		{
			name:   "apostrophe one word",
			input:  "She said ' hello ' world",
			output: "She said 'hello' world",
		},
		{
			name:   "apostrophe after colon",
			input:  "hello: ' world '",
			output: "hello: 'world'",
		},
		{
			name:   "apostrophe after comma",
			input:  "hello, ' world '",
			output: "hello, 'world'",
		},
		{
			name:   "contraction stays unchanged",
			input:  "that's a good idea",
			output: "that's a good idea",
		},
		{
			name:   "contraction plus quoted word",
			input:  "that's ' something ' good",
			output: "that's 'something' good",
		},

		{
			name:   "a before vowel",
			input:  "a apple",
			output: "an apple",
		},
		{
			name:   "A before vowel",
			input:  "A amazing rock",
			output: "An amazing rock",
		},
		{
			name:   "a before h",
			input:  "a happy person",
			output: "an happy person",
		},
		{
			name:   "A before H",
			input:  "A Huge problem",
			output: "An Huge problem",
		},
		{
			name:   "final a must not disappear",
			input:  "this is a",
			output: "this is a",
		},

		{
			name:   "up then low",
			input:  "hello (up) (low)",
			output: "hello",
		},
		{
			name:   "low then up",
			input:  "HELLO (low) (up)",
			output: "HELLO",
		},
		{
			name:   "up then cap",
			input:  "hello (up) (cap)",
			output: "Hello",
		},
		{
			name:   "cap then up",
			input:  "hello (cap) (up)",
			output: "HELLO",
		},
		{
			name:   "up low cap",
			input:  "hello (up) (low) (cap)",
			output: "Hello",
		},

		{
			name:   "counted low",
			input:  "ONE TWO THREE (low, 3)",
			output: "one two three",
		},
		{
			name:   "counted cap",
			input:  "one two three (cap, 3)",
			output: "One Two Three",
		},

		{
			name:   "all punctuation",
			input:  "one . two , three ! four ? five : six ; seven",
			output: "one. two, three! four? five: six; seven",
		},
		{
			name:   "punctuation groups",
			input:  "hello . . . world ! ? really",
			output: "hello... world!? really",
		},
		{
			name:   "punctuation attached both sides",
			input:  "hello...world",
			output: "hello... world",
		},
		{
			name:   "mixed punctuation group",
			input:  "hello!?world",
			output: "hello!? world",
		},
		{
			name:   "many exclamation marks",
			input:  "hello ! ! !",
			output: "hello!!!",
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			got := run(ts.input, process.LeftToRight)
			if got != ts.output {
				t.Fatalf("\ninput:    %q\nGot:      %q\noutput: %q", ts.input, got, ts.output)
			}
		})
	}
}
