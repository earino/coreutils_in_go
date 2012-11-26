package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var newlinePtr = flag.Bool("n", false, "Do not print the trailing newline character.")

	flag.Parse();
	var args = os.Args[1:]

	if *newlinePtr {
		args = os.Args[2:]
	}

	for _, value := range args {
		fmt.Printf("%s ", value)
	}
	if ! *newlinePtr {
		fmt.Println()
	}
}