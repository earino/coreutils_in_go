package main

import (
	"log"
	"os"
	"io"
	"bufio"
	"strings"
	"fmt"
	"flag"
)

var repeats int = 0
var countPtr = flag.Bool("c", false, "Generate an output report in default style except that each line is preceded by a count of the number of times it occurred. If this option is specified, the -u and -d options are ignored if either or both are also present.")
var	ignoreCasePtr = flag.Bool("i", false, "Ignore case differences when comparing lines")
var uniqPtr = flag.Bool("u", false, "Print only those lines which are not repeated (unique) in the input.")
var dupPtr = flag.Bool("d", false, "Print only those lines which are repeated in the input.")
var ignoreFieldsPtr = flag.Int("f", 0, "Ignore a number of fields in a line")
var ignoreCharsPtr = flag.Int("s", 0, "Ignore the first chars characters in each input line when doing comparisons.")

func show(line string) {
	if *countPtr && line != "" {
		fmt.Printf("%4d %s", repeats + 1, line)
	} else if (*countPtr && repeats > 0) || (*uniqPtr) {
		fmt.Printf("%s", line)
	}
}

func skip(line string) string {
	var infield bool = false
	var nchars, nfields, i int = 0, 0, 0
	
	nfields = *ignoreFieldsPtr

	for ; nfields > 0 && i < len(line); i++ {
		if line[i] == ' ' {
			if infield {
				infield = false
				nfields--
			}
		} else if ! infield {
			infield = true
		}
	}

	line = line[i:len(line)]

	nchars = *ignoreCharsPtr
	i = 0

	for ; nchars > 0 && i < len(line); i++ {
		nchars--
	}
	
	return line[i:len(line)]
}

func main() {
	flag.Parse()

	if *countPtr  {
		if *dupPtr || *uniqPtr {
			fmt.Println("usage:")
		} 
	} else {
		if ! *dupPtr && ! *uniqPtr {
			*dupPtr = true
			*uniqPtr = true
		}
	}

	for _, element := range flag.Args() {
		var r io.Reader

		if element ==  "-" {
			r = os.Stdin
		} else {
			var err error
			r, err = os.Open(element)
			if err != nil {
				log.Fatal(err)
			}
		}

		bufreader := bufio.NewReader(r)

		var thisline, prevline string
		var err error

		for {
			thisline, err = bufreader.ReadString('\n')

			var t1, t2 string
			if *ignoreFieldsPtr != 0 || *ignoreCharsPtr != 0 {
				t1 = skip(thisline)
				t2 = skip(prevline)
			} else {
				t1 = thisline
				t2 = prevline
			}

			var cmp bool

			if *ignoreCasePtr {
				cmp = strings.EqualFold(t1, t2)
			} else {
				cmp = t1 == t2
			}

			if ! cmp {
				show(prevline)
    			thisline, prevline = prevline, thisline
    			repeats = 0
    		} else {
    			repeats++
    		}

			if err == io.EOF {
 				break
    		}    	
		}

		show(prevline)
	}
}