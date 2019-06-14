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

var _listofWordInfos []wordInfo

var _sentence []string

func main() {
	argsWithoutProg := os.Args[1:]

	action := argsWithoutProg[0]

	rand.Seed(time.Now().UnixNano())

	switch action {
	case "parse":
		parseText(readAsText(argsWithoutProg[1]))
		out, _ := json.Marshal(_listofWordInfos)
		writeBytes(out, argsWithoutProg[2])
		break
	case "import":
		json.Unmarshal(read(argsWithoutProg[1]), &_listofWordInfos)
		writeToOutput(_listofWordInfos[0].Word + " - " + strconv.Itoa(_listofWordInfos[0].Count))
		break
	case "speech":
		json.Unmarshal(read(argsWithoutProg[1]), &_listofWordInfos)
		totalWords := len(_listofWordInfos)
		startWord := wordInfo{}
		pos := -1
		if len(argsWithoutProg) > 2 {
			startWord = getWordPosition(argsWithoutProg[2])
		} else {

			pos = rand.Intn(totalWords)
			startWord = _listofWordInfos[pos]
		}
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
		getWord(startWord, beforeWord, afterWord, 0, rand.Intn(20-5)+5)
		writeToOutput(strings.Join(_sentence[:], " "))
		break
	case "find":
		json.Unmarshal(read(argsWithoutProg[1]), &_listofWordInfos)
		count := 0
		foundWord := wordInfo{}
		for _, word := range _listofWordInfos {
			if word.Word == argsWithoutProg[2] {
				foundWord = _listofWordInfos[count]
				break
			}
			count++
		}
		writeToOutput(foundWord.Word + " - " + strconv.Itoa(foundWord.Count))
		break
	default:
		writeToOutput("no action specified")
	}
}

func getWord(startWord wordInfo, beforeWord wordInfo, afterWord wordInfo, count int, limit int) {
	count++
	if count > limit {
		return
	}
	newStartWord := wordInfo{}
	newBeforeWord := wordInfo{}
	newAfterWord := wordInfo{}
	pos := 0
	_sentence = append(_sentence, startWord.Word)
	if afterWord.Word != "" {
		_sentence = append(_sentence, afterWord.Word)
	}
	if len(afterWord.After) > 0 {
		pos = rand.Intn(len(afterWord.After))
		newStartWord = afterWord.After[pos]
	} else {
		newStartWord = getWordPosition(afterWord.Word)
		if len(newStartWord.After) > 0 {
			pos = rand.Intn(len(newStartWord.After))
			newStartWord = getWordPosition(newStartWord.After[pos].Word)
		}
	}
	if newStartWord.Word == "" {
		pos = rand.Intn(len(_listofWordInfos))
		newStartWord = _listofWordInfos[pos]
	}
	if len(newStartWord.After) > 0 {
		pos = rand.Intn(len(newStartWord.After))
		newAfterWord = newStartWord.After[pos]
	}
	if len(newStartWord.Before) > 0 {
		pos = rand.Intn(len(newStartWord.Before))
		newBeforeWord = newStartWord.Before[pos]
	}
	getWord(newStartWord, newBeforeWord, newAfterWord, count, limit)
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
		_listofWordInfos = findWord(word, _listofWordInfos, beforeWord, afterWord)
		count++
	}

	for _, word := range _listofWordInfos {
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

func getWordPosition(word string) wordInfo {
	for _, w := range _listofWordInfos {
		if w.Word == word {
			return w
		}
	}

	return wordInfo{}
}
