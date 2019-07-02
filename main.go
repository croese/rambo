package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE: rambo <path-to-mach-o-binary>\n")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Printf("Analyzing '%s'...\n", filePath)
}
