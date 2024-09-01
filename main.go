package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// If there are arguments, cat each file
	// Else use stdin, enables use of redirection with <
	if len(os.Args) > 1 {
		for _, files := range os.Args[1:] {
			if err := catWithPath(files); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	} else {
		if err := gat(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func catWithPath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error reading file: %s, %w", path, err)
	}
	defer file.Close()

	if err := gat(os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return nil
}

func gat(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("Error reading file: %w", err)
		}
		fmt.Fprintln(os.Stdout, scanner.Text())
	}

	// Since EOF also breaks the loop, add second err check to see if any other
	// errors occurred
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading file: %w", err)
	}
	return nil
}
