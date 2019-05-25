package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func write(s string, fn string) {
	writeFull(s, fn, false)
}

func writeFull(s string, fn string, silent bool) {
	f, err := os.Create(fn)
	if err != nil {
		writeToOutput(err.Error())
		return
	}
	l, err := f.WriteString(s)
	if err != nil {
		writeToOutput(err.Error())
		f.Close()
		return
	}
	if !silent {
		writeToOutput(strconv.Itoa(l) + " bytes written successfully to " + fn)
	}
	err = f.Close()
	if err != nil {
		writeToOutput(err.Error())
		return
	}
}

func writeBytes(b []byte, fn string) {
	writeBytesFull(b, fn, false)
}

func writeBytesFull(b []byte, fn string, silent bool) {
	f, err := os.Create(fn)
	if err != nil {
		writeToOutput(err.Error())
		return
	}
	l, err := f.Write(b)
	if err != nil {
		writeToOutput(err.Error())
		f.Close()
		return
	}

	if !silent {
		writeToOutput(strconv.Itoa(l) + " bytes written successfully to " + fn)
	}

	err = f.Close()
	if err != nil {
		writeToOutput(err.Error())
		return
	}
}

func read(fn string) []byte {
	b, _ := ioutil.ReadFile(fn)
	return b
}

func readAsText(fn string) string {
	r := read(fn)
	s := convertToString(r)
	return s
}

func makeDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func fileExists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func fileSize(f string) int64 {
	file, err := os.Open(f)
	if err != nil {
		writeToOutput(err.Error())
	}
	fi, err := file.Stat()
	if err != nil {
		writeToOutput(err.Error())
	}

	return fi.Size()
}
