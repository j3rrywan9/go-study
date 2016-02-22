package main

import (
	"flag"
	"fmt"
)

var runFile string

func init() {
	flag.StringVar(&runFile, "r", "", "The input file is a run file instead of a JSON file.")
	flag.StringVar(&runFile, "run-file", "", "The input file is a run file instead of a JSON file.")
}

func main() {
	flag.Parse()
	fmt.Println(runFile)
	fmt.Println(flag.Args())
}
