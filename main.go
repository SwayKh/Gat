package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Not enough arguments, Usage: gat [FILE]s...")
		os.Exit(1)
	}

	for _, files := range os.Args[1:] {
		file, err := os.Open(files)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		// Read file line by line
		scanner.Split(bufio.ScanLines)

		// Returns a bool if newline character is found (I think)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				// End the iteration if we encounter an error
				fmt.Fprintln(os.Stderr, err)
				break
			}
			fmt.Fprintln(os.Stdout, scanner.Text())
		}
	}
}

// TODO
// Pull the license from git repo (DONE)
// Add codeblock to readme (DONE)
// Use os.READ to read, readfile() loads whole file in memory (Okay)
// Use bufio scanner to read file (Done)
// Use for loop and read in increment for large files (Used bool from
// scanner.scan())
// Check for io.EOF to stop reading (scanner.scan() stops on EOF)
// Make is work with stdin ( send file to it with < ) (Dont want to)
// Make is take multiple files (Done)
// Check cat implementation and bat implementation
// Check i billion row challenge to see what methods are used for performance
// Print errors to stderr (Done)
