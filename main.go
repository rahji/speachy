package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jdkato/prose/v2"
	"github.com/rahji/speachy/internal/textinput"
	"github.com/spf13/pflag"
)

func main() {

	var (
		inputFile  string
		outputFile string
		parts      string
		help       bool
	)

	pflag.StringVarP(&inputFile, "infile", "i", "", "input file (if not specified, reads from STDIN)")
	pflag.StringVarP(&outputFile, "outfile", "o", "", "optional output file")
	pflag.StringVarP(&parts, "parts", "p", "", "comma-separated list of parts of speech to return")
	pflag.BoolVarP(&help, "help", "h", false, "show help message")
	pflag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	// get the text, from STDIN if inputFile is empty
	text, err := textinput.GetText(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// get the remaining args, if any
	args := pflag.Args()

	// make a map of parts of speech abbrevations
	// the keys come from either the comma-separated list in args[0]
	// or from a bubble tea interactive list
	partsMap := make(map[string]bool)

	// if no arguments then show the interactive list
	if len(args) != 1 {
		p := tea.NewProgram(initialModel())
		m, err := p.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: %v", err)
		}

		finalModel := m.(model)
		for i := range finalModel.selected {
			partsMap[finalModel.choices[i].abbr] = true
		}
	} else {
		// split the comma-separated list from the argument
		for _, part := range strings.Split(args[0], ",") {
			part = strings.TrimSpace(part)
			part = strings.ToUpper(part)
			if part != "" {
				partsMap[part] = true
			}
		}
	}

	// double check that something is in the map
	if len(partsMap) == 0 {
		fmt.Fprintln(os.Stderr, "Error: comma-separated list cannot be empty")
		os.Exit(1)
	}

	// create a new prose document from the text string
	doc, err := prose.NewDocument(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// write to output file if that flag was used
	var out *os.File
	if outputFile != "" {
		out, err = os.Create(outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	err = outputTags(doc, partsMap, out)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// outputTags iterates over a doc and outputs the chosen tags to an *os.File
func outputTags(doc *prose.Document, parts map[string]bool, f *os.File) error {
	writer := bufio.NewWriter(f)
	for _, tok := range doc.Tokens() {
		if parts[tok.Tag] {
			_, err := writer.WriteString(tok.Text + "\n")
			if err != nil {
				return err
			}
		}
	}
	writer.Flush()
	return nil
}

func usage() {
	fmt.Print(`
Usage: speachy [OPTIONS] [item1,item2,...]

  -i, --infile  string    input file (if not specified, reads from STDIN)
  -o, --outfile string    optional output file
  -h, --help              show help message

The optional argument is a comma-separated list of abbreviations for 
parts of speech to find in the text. If none are specified, 
an interactive list is presented.

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
