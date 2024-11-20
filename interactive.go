package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type part struct {
	abbr string
	name string
}

type model struct {
	choices  []part
	selected map[int]struct{}
	cursor   int
}

func initialModel() model {
	return model{
		choices: []part{
			{abbr: "CC", name: "conjunction, coordinating"},
			{abbr: "CD", name: "cardinal number"},
			{abbr: "DT", name: "determiner"},
			{abbr: "EX", name: "existential there"},
			{abbr: "FW", name: "foreign word"},
			{abbr: "IN", name: "conjunction, subordinating or preposition"},
			{abbr: "JJ", name: "adjective"},
			{abbr: "JJR", name: "adjective, comparative"},
			{abbr: "JJS", name: "adjective, superlative"},
			{abbr: "LS", name: "list item marker"},
			{abbr: "MD", name: "verb, modal auxiliary"},
			{abbr: "NN", name: "noun, singular or mass"},
			{abbr: "NNP", name: "noun, proper singular"},
			{abbr: "NNPS", name: "noun, proper plural"},
			{abbr: "NNS", name: "noun, plural"},
			{abbr: "PDT", name: "predeterminer"},
			{abbr: "POS", name: "possessive ending"},
			{abbr: "PRP", name: "pronoun, personal"},
			{abbr: "PRP$", name: "pronoun, possessive"},
			{abbr: "RB", name: "adverb"},
			{abbr: "RBR", name: "adverb, comparative"},
			{abbr: "RBS", name: "adverb, superlative"},
			{abbr: "RP", name: "adverb, particle"},
			{abbr: "SYM", name: "symbol"},
			{abbr: "TO", name: "infinitival to"},
			{abbr: "UH", name: "interjection"},
			{abbr: "VB", name: "verb, base form"},
			{abbr: "VBD", name: "verb, past tense"},
			{abbr: "VBG", name: "verb, gerund or present participle"},
			{abbr: "VBN", name: "verb, past participle"},
			{abbr: "VBP", name: "verb, non-3rd person singular present"},
			{abbr: "VBZ", name: "verb, 3rd person singular present"},
			{abbr: "WDT", name: "wh-determiner"},
			{abbr: "WP", name: "wh-pronoun, personal"},
			{abbr: "WP$", name: "wh-pronoun, possessive"},
			{abbr: "WRB", name: "wh-adverb"},
		},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			_, exists := m.selected[m.cursor]
			if exists {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select parts of speech using space bar, confirm with enter:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, exists := m.selected[i]; exists {
			checked = "✓"
		}

		s += fmt.Sprintf("%s [%s] (%s) %s\n", cursor, checked, choice.abbr, choice.name)
	}

	s += "\nPress q to quit.\n"
	return s
}
