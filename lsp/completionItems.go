package main

import (
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Fountain struct {
	kind    string
	matches []string
}

func UpdateCompletionList(content string) (symbols []protocol.CompletionItem) {
	matches := getEveryMatch(content)
	for _, match := range matches {
		for _, item := range match.matches {
			symbols = append(symbols, protocol.CompletionItem{
				Label:      match.kind,
				InsertText: &item,
			})
		}
	}

	return symbols
}

func getEveryMatch(content string) (matches []Fountain) {
	everyMatchPossibile := -1

	sceneHeadingsRE := regexp.MustCompile(`(?i)^[ \\t]*([.](?![.])|(?:[*]{0,3}_?)(?:int|ext|est|int[.]?\/ext|i[.]?\/e)[. ])(.+?)(#[-.0-9a-z]+#)?$`)
	sceneHeadings := sceneHeadingsRE.FindAllString(content, everyMatchPossibile)
	sceneHeadings = removeRepeatedValues(sceneHeadings)
	matches = append(matches, Fountain{"Heading", sceneHeadings})

	return matches
}

func removeRepeatedValues(matches []string) (uniqueMatches []string) {
	for _, match := range matches {
		isUnique := true
		for _, uniqueMatch := range uniqueMatches {
			if match == uniqueMatch {
				isUnique = false
			}
		}
		if isUnique {
			uniqueMatches = append(uniqueMatches, match)
		}
	}
	return uniqueMatches
}
