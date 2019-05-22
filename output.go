package main

import (
	"fmt"
)

func writeToOutput(m string) {
	fmt.Println(m)
}

func convertToString(r []byte) string {
	return fmt.Sprintf("%s", r)
}
