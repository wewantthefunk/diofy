package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type wordInfo struct {
	Word   string     `json:"word"`
	Count  int        `json:"count"`
	Before []wordInfo `json:"before"`
	After  []wordInfo `json:"after"`
}

var listOfWordInfos []wordInfo

func main() {
	argsWithoutProg := os.Args[1:]

	action := argsWithoutProg[0]

	switch action {
	case "parse":
		parseText(readAsText(argsWithoutProg[1]))
		out, _ := json.Marshal(listOfWordInfos)
		writeBytes(out, argsWithoutProg[2])
		break
	case "import":
		json.Unmarshal(read(argsWithoutProg[1]), &listOfWordInfos)
		writeToOutput(listOfWordInfos[0].Word + " - " + strconv.Itoa(listOfWordInfos[0].Count))
		break
	case "speech":
		json.Unmarshal(read(argsWithoutProg[1]), &listOfWordInfos)
		totalWords := len(listOfWordInfos)
		rand.Seed(time.Now().UnixNano())
		pos := rand.Intn(totalWords)
		startWord := listOfWordInfos[pos]
		beforeWord := wordInfo{}
		afterWord := wordInfo{}
		if len(startWord.Before) > 0 {
			pos = rand.Intn(len(startWord.Before))
			beforeWord = startWord.Before[pos]
		}
		if len(startWord.After) > 0 {
			pos = rand.Intn(len(startWord.After))
			afterWord = startWord.After[pos]
		}
		getWord(startWord, beforeWord, afterWord)
		break
	}
}

func getWord(startWord wordInfo, beforeWord wordInfo, afterWord wordInfo) {
	writeToOutput(startWord.Word)
	writeToOutput(beforeWord.Word)
	writeToOutput(afterWord.Word)
}

func parseText(test string) {
	clean := strings.Replace(test, ".", "", -1)
	clean = strings.Replace(clean, "  ", " ", -1)
	clean = strings.Replace(clean, ",", "", -1)
	clean = strings.ToLower(clean)
	words := strings.Split(clean, " ")

	count := 0

	for _, word := range words {
		if word == "" {
			continue
		}
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
		writeToOutput(word.Word + " - " + strconv.Itoa(word.Count))
	}
}

func findWord(word string, wordInfoList []wordInfo, beforeWord string, afterWord string) []wordInfo {
	pos := 0
	found := false
	for _, w := range wordInfoList {
		if w.Word == word {
			wordInfoList[pos].Count++
			found = true
			break
		}
		pos++
	}

	if !found {
		lowi := wordInfo{}
		lowi.Count = 1
		lowi.Word = word
		wordInfoList = append(wordInfoList, lowi)
	}

	if beforeWord != "" {
		wordInfoList[pos].Before = findWord(beforeWord, wordInfoList[pos].Before, "", "")
	}

	if afterWord != "" {
		wordInfoList[pos].After = findWord(afterWord, wordInfoList[pos].After, "", "")
	}

	return wordInfoList
}
