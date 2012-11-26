package main

import (
	"fmt"
	"os/user"
)

func main() {
	me, _ := user.Current()
	fmt.Println(me.Username)
}