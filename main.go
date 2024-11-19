package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jdkato/prose/v2"
	"github.com/rahji/speachy/internal/textinput"
	"github.com/spf13/pflag"
)

func main() {

	var (
		inputFile string
		parts     string
		help      bool
	)

	pflag.StringVarP(&inputFile, "file", "f", "", "input file (if not specified, reads from STDIN)")
	pflag.StringVarP(&parts, "parts", "p", "", "comma-separated list of parts of speech to return")
	pflag.BoolVarP(&help, "help", "h", false, "show help message")
	pflag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	// remaining args
	args := pflag.Args()

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Error: a comma-separated list of parts of speech is required")
		os.Exit(1)
	}

	// split the comma-separated list into a map
	partsMap := make(map[string]bool)
	for _, part := range strings.Split(args[0], ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			partsMap[part] = true
		}
	}

	if len(partsMap) == 0 {
		fmt.Fprintln(os.Stderr, "Error: comma-separated list cannot be empty")
		os.Exit(1)
	}

	// get the text, from STDIN if inputFile is empty
	text, err := textinput.GetText(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create a new prose document with the default configuration:
	doc, err := prose.NewDocument(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		if partsMap[tok.Tag] {
			fmt.Println(tok.Text)
		}
	}

}

func usage() {
	fmt.Print(`

usage: speachy [--file] [--help] <arg1[,arg2]>

  -f, --file string    input file (if not specified, reads from STDIN)
  -h, --help           show help message

the required argument is a comma-separated list of parts of speech to find:

  CC    conjunction, coordinating
  CD    cardinal number
  DT    determiner
  EX    existential there
  FW    foreign word
  IN    conjunction, subordinating or preposition
  JJ    adjective
  JJR   adjective, comparative
  JJS   adjective, superlative
  LS    list item marker
  MD    verb, modal auxiliary
  NN    noun, singular or mass
  NNP   noun, proper singular
  NNPS  noun, proper plural
  NNS   noun, plural
  PDT   predeterminer
  POS   possessive ending
  PRP   pronoun, personal
  PRP$  pronoun, possessive
  RB    adverb
  RBR   adverb, comparative
  RBS   adverb, superlative
  RP    adverb, particle
  SYM   symbol
  TO    infinitival to
  UH    interjection
  VB    verb, base form
  VBD   verb, past tense
  VBG   verb, gerund or present participle
  VBN   verb, past participle
  VBP   verb, non-3rd person singular present
  VBZ   verb, 3rd person singular present
  WDT   wh-determiner
  WP    wh-pronoun, personal
  WP$   wh-pronoun, possessive
  WRB   wh-adverb
`)
}
