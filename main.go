package main

type wordInfo struct {
	word string
	count int
	before []wordInfo
	after []wordInfo
}

func main() {
	writeToOutput("hello, world")
}
