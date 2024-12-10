package main

import (
	"flag"
	"fmt"

	"software-design/pattern"
)

func main() {
	flag.Parse()
	// Get the command line argument
	arg := flag.Arg(0)

	pattern, err := pattern.CreatePattern(arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok := pattern.ProcessPattern()
	if !ok {
		fmt.Println("Error processing pattern")
		return
	}

}
