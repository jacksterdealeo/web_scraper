package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) >= 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %v\n", os.Args[1])

}
