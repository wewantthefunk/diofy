package main

import (
	"strconv"
	"strings"
)

type wordInfo struct {
	word   string
	count  int
	before []wordInfo
	after  []wordInfo
}

var listOfWordInfos []wordInfo

func main() {
	test := "this is a test. this is only a test. we are here to test the system."

	clean := strings.Replace(test, ".", "", -1)
	clean = strings.Replace(clean, "  ", " ", -1)
	clean = strings.Replace(clean, ",", "", -1)
	words := strings.Split(clean, " ")

	count := 0

	for _, word := range words {
		beforeWord := ""
		afterWord := ""
		if count > 0 {
			beforeWord = words[count-1]
		}
		if count < len(words)-1 {
			afterWord = words[count+1]
		}
		listOfWordInfos = findWord(word, listOfWordInfos, beforeWord, afterWord)
		count++
	}

	for _, word := range listOfWordInfos {
		writeToOutput(word.word + " - " + strconv.Itoa(word.count))
	}
}

func findWord(word string, wordInfoList []wordInfo, beforeWord string, afterWord string) []wordInfo {
	pos := 0
	found := false
	for _, w := range wordInfoList {
		if w.word == word {
			wordInfoList[pos].count++
			found = true
			break
		}
		pos++
	}

	if !found {
		lowi := wordInfo{}
		lowi.count = 1
		lowi.word = word
		wordInfoList = append(wordInfoList, lowi)
	}

	if beforeWord != "" {
		wordInfoList[pos].before = findWord(beforeWord, wordInfoList[pos].before, "", "")
	}

	if afterWord != "" {
		wordInfoList[pos].after = findWord(afterWord, wordInfoList[pos].after, "", "")
	}

	return wordInfoList
}
