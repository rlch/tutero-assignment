package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var logger *log.Logger

func main() {
	Init()
	logger = log.New(os.Stdout, "[runner]", 0)
	logger.SetPrefix("[runner] ")
	logger.Println("Running tests...")

	args := os.Args[1:]

	runner := exec.Command("go", append([]string{
		"test", ".", "-run=^TestRunner$",
	}, args...)...)
	runner.Stdout = os.Stdout
	runner.Stderr = os.Stderr
	runner.Run()

	fmt.Print("\n\n")
	logger.Println("Running benchmarks...")
	benchmarker := exec.Command("go", append([]string{
		"test", ".", "-run=^$", "-bench=BenchmarkRunner",
	}, args...)...)
	benchmarker.Stdout = os.Stdout
	benchmarker.Stderr = os.Stderr
	benchmarker.Run()
}
