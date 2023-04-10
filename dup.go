package main 

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		countLinesFromFiles(files, counts)
	}

	printLines(counts)
}

func countLines(f *os.File, lines map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines[input.Text()]++
	} 
	//Note: ignoring potential errors from input.Err()

	fmt.Println("input read successfully")
}

func countLinesFromFiles(files []string, lines map[string]int) {
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, lines)
		f.Close()
	}
}


func printLines(lines map[string]int) {
	for line, n := range lines{
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

