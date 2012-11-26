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
	flag.Parse()

	var writers = []io.Writer{os.Stdout}

	for _, teeTo := range flag.Args() {
		var fileFlags int

		if *appendPtr {
			fileFlags = os.O_WRONLY|os.O_CREATE|os.O_APPEND
		} else {
			fileFlags = os.O_WRONLY|os.O_CREATE|os.O_TRUNC
		}

		w, err := os.OpenFile(teeTo, fileFlags, 0644)

		if err != nil {
			log.Fatal(err)
		}
		
		writer := io.Writer(w)
		bufwriter := bufio.NewWriter(writer)

		defer w.Close()
		defer bufwriter.Flush()

		writers = append(writers, bufwriter)
	}

	teeWriter := io.MultiWriter(writers...)
	bufwriter := bufio.NewWriter(teeWriter)

	defer bufwriter.Flush()

	io.Copy(teeWriter, os.Stdin)
}