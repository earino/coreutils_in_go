package main

import (
	"io"
	"os"
	"os/signal"
	"flag"
	"log"
)

var appendPtr = flag.Bool("a", false, "Append the output to the files rather than overwriting them.")
var ignoreIntPtr = flag.Bool("i", false, "Ignore the SIGINT signal.")

func main() {
	flag.Parse()

	var writers = []io.Writer{os.Stdout}
	var fileFlags int = os.O_WRONLY|os.O_CREATE

	if *appendPtr {
		fileFlags |= os.O_APPEND
	} else {
		fileFlags |= os.O_TRUNC
	}

	if *ignoreIntPtr {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
	}

	for _, teeTo := range flag.Args() {
		w, err := os.OpenFile(teeTo, fileFlags, 0644)

		if err != nil {
			log.Fatal(err)
		}
		
		writers = append(writers, w)
	}

	teeWriter := io.MultiWriter(writers...)

	io.Copy(teeWriter, os.Stdin)
	os.Exit(0)
}