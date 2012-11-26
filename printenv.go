package main

import ("os"
		"fmt"
		"strings")

func main() {
	environment := os.Environ()

	var theOne = len(os.Args) > 1

	for i := 0; i < len(environment); i++ {
		if theOne {
			if 0 == strings.Index(environment[i], os.Args[1]) {
				temp := strings.Split(environment[i], "=")
				fmt.Println(temp[1])	
			}
		} else {
			fmt.Println(environment[i])
		}
	}
}