package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Not enough arguments, Usage: gat [FILE]s...")
		os.Exit(1)
	}

	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(os.Stdout, string(f))

	// for file := range os.Args[1:] {
	// f, err := os.Open(os.Args[file])
	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	fmt.Fprintln(os.Stdout, scanner.Text())
	// }
	// }
}

// TODO
// Pull the license from git repo (DONE)
// Add codeblock to readme (DONE)
// Use os.READ to read, readfile() loads whole file in memory
// Use bufio scanner to read file
// Use for loop and read in increment for large files
// Check for io.EOF to stop reading
// Make is work with stdin ( send file to it with < )
// Make is take multiple files
// Check cat implementation and bat implementation
// Check i billion row challenge to see what methods are used for performance
// Print errors to stderr
