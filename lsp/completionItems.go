package main

import (
	"regexp"

	"github.com/tliron/commonlog"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Fountain struct {
    kind    string
	matches []string
}
var defaultHeadings = []string {  "INT.", "EXT.", "INT/EXT" }

func UpdateCompletionList(content string) (symbols []protocol.CompletionItem) {
	matches := getEveryMatch(content)
    commonlog.NewInfoMessage(0, matches[0].matches...).Send()
	for _, match := range matches {
		for i, item := range match.matches {
			symbols = append(symbols, protocol.CompletionItem{
                Documentation: match.kind,
                Label:      item,
				InsertText: &match.matches[i],
			})
		}
	}

	return symbols
}

func getEveryMatch(content string) (matches []Fountain) {
	everyMatchPossibile := -1

	sceneHeadingsRE, err := regexp.Compile(`(INT\.|EXT\.|INT/EXT).*- (DAY|NIGHT|DUSK|DAWN)`)
    if err != nil {
        commonlog.NewInfoMessage(0, "Invalid regex").Send()
        commonlog.NewInfoMessage(0, err.Error()).Send()
    }

	sceneHeadings := sceneHeadingsRE.FindAllString(content, everyMatchPossibile)
	sceneHeadings = removeRepeatedValues(sceneHeadings)
    sceneHeadings = append(sceneHeadings, defaultHeadings...)
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
