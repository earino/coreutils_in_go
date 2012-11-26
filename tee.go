package main

import (
	"io"
	"bufio"
	"os"
	"flag"
	"log"
)

var appendPtr = flag.Bool("a", false, "Append the output to the files rather than overwriting them.")
var ignoreCasePtr = flag.Bool("i", false, "Ignore the SIGINT signal.")

func main() {
	var err error

	flag.Parse()

	bufferedStdin := bufio.NewReader(os.Stdin)
	bufferedStdout := bufio.NewWriter(os.Stdout)

	var writers = []bufio.Writer{*bufferedStdout}

	for _, teeTo := range flag.Args() {
		var w io.Writer
		if *appendPtr {
			w, err = os.OpenFile(teeTo, os.O_CREATE|os.O_APPEND, 0644)
		} else {
			w, err = os.OpenFile(teeTo, os.O_CREATE, 0644)
		}

		if err != nil {
			log.Fatal(err)
		}

		bufwriter := bufio.NewWriter(w)
		writers = append(writers, *bufwriter)
	}

	teeWriter := io.MultiWriter(writers)

	io.Copy(teeWriter, bufferedStdin)
}