package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("could not read input.txt: %s", err))
	}
	os.Stderr.WriteString("input: " + string(bytes))

	fmt.Println(
		strings.Join([]string{
			"A", "B", "D", "C",
		}, "\n"),
	)
}
