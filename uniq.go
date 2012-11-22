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


func show(line string) {
	if *countPtr && line != "" {
		fmt.Printf("%4d %s", repeats + 1, line)
	} else if (*countPtr && repeats > 0) || (*uniqPtr && repeats != 0) {
		fmt.Printf("%s", line)
	}
}

func main() {

/*
	ignoreFieldsPtr := flag.Int("f", 0, "Ignore a number of fields in a line")
	ignoreCharsPtr := flag.Int("s", 0, "Ignore the first chars characters in each input line when doing comparisons.")
*/
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

	fmt.Println("countPtr: ", *countPtr, " dupPtr:", *dupPtr, " uniqPtr: ", *uniqPtr)
/*
	fmt.Println("uniq: ", *uniqPtr)
	fmt.Println("dup: ", *dupPtr)
	fmt.Println("ignoreCase: ", *ignoreCasePtr)
	fmt.Println("ignoreFields: ", *ignoreFieldsPtr)
	fmt.Println("ignoreChars: ", *ignoreCharsPtr)
	fmt.Println("tail: ", flag.Args())
*/
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
			var cmp bool

			if *ignoreCasePtr {
				cmp = strings.EqualFold(prevline, thisline)
			} else {
				cmp = prevline == thisline
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

		//fmt.Println("at: ", index, " we have: ", element)
	}
}