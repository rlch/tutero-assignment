package main

import (
	"log"
	"os"
	"os/exec"
)

var logger *log.Logger

func main() {
	InitFlags()
	logger = log.New(os.Stdout, "[runner] ", 0)
	logger.Println("Running tests...")

	args := os.Args[1:]

	runner := exec.Command("go", append([]string{
		"test", ".", "-run=^TestRunner$",
	}, args...)...)
	runner.Stdout = os.Stdout
	runner.Stderr = os.Stderr
	runner.Run()
}
