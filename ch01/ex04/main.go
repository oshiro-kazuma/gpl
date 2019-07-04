package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Dup(os.Args[1:])
}

func Dup(files []string) {
	counts := make(map[string]int)
	filenames := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !contains(filenames[line], f.Name()) {
			filenames[line] = append(filenames[line], f.Name())
		}
	}
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
