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
		if err := printFile(files); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func printFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error reading file: %s, %w", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("Error reading file: %s, %w", path, err)
		}
		fmt.Fprintln(os.Stdout, scanner.Text())
	}

	// Since EOF also breaks the loop, add second err check to see if any other
	// errors occurred
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading file: %s, %w", path, err)
	}
	return nil
}
