# Speachy

![Speachy](speachy.png "Speachy")

Speachy is a command-line tool that allows you to get a list of words, from your text,
that match specific parts of speech.

## Examples

View a list of all types of adjectives from a text file:

```
speachy -i inputfile.txt JJ,JJR,JJS
```

Speachy will also take the text from STDIN, if you leave the input filename flag off of the command:

```
cat intputfile.txt | speachy JJS
```

You can also output the words to a text file:

*(If you leave the list of part of speech tags off of the command, Speachy
will present an interactive list from which to choose tags.)*

```
speachy -i inputfile.txt -o outputfile.txt
```

## Usage

```
speachy [OPTIONS] [item1,item2,...]

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
```

## Thanks

All the Natural Language Processing work is done by Go [prose package](https://github.com/jdkato/prose). Thanks, too, to Charm for the [Bubble Tea](https://github.com/charmbracelet/bubbletea) package for making the command-line fun. Also thanks to [Go](https://go.dev/), a fun language to learn and use.
