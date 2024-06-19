package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type options struct {
	column  int
	numeric bool
	reverse bool
	unique  bool
}

func parseOptions() options {
	column := flag.Int("k", 0, "Column for sorting")
	numeric := flag.Bool("n", false, "Sort by numerical value")
	reverse := flag.Bool("r", false, "Sort in reverse order")
	unique := flag.Bool("u", false, "Output only unique lines")
	flag.Parse()
	return options{
		column:  *column - 1,
		numeric: *numeric,
		reverse: *reverse,
		unique:  *unique,
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func uniqueLines(lines []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range lines {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func compareLines(a, b string, opts options) bool {
	aFields := strings.Fields(a)
	bFields := strings.Fields(b)
	var aKey, bKey string
	if opts.column >= 0 && opts.column < len(aFields) {
		aKey = aFields[opts.column]
	} else {
		aKey = a
	}
	if opts.column >= 0 && opts.column < len(bFields) {
		bKey = bFields[opts.column]
	} else {
		bKey = b
	}
	if opts.numeric {
		aNum, aErr := strconv.ParseFloat(aKey, 64)
		bNum, bErr := strconv.ParseFloat(bKey, 64)
		if aErr == nil && bErr == nil {
			return aNum < bNum
		}
	}
	return aKey < bKey
}

func sortLines(lines []string, opts options) []string {
	if opts.unique {
		lines = uniqueLines(lines)
	}
	for i := 0; i < len(lines)-1; i++ {
		for j := 0; j < len(lines)-i-1; j++ {
			if (opts.reverse && compareLines(lines[j], lines[j+1], opts)) ||
				(!opts.reverse && !compareLines(lines[j], lines[j+1], opts)) {
				lines[j], lines[j+1] = lines[j+1], lines[j]
			}
		}
	}
	return lines
}

func reverseLines(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func main() {
	opts := parseOptions()
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: sort [options] <input_file> <output_file>")
		return
	}
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)
	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	sortedLines := sortLines(lines, opts)
	if opts.reverse {
		reverseLines(sortedLines)
	}
	err = writeLines(sortedLines, outputFile)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}
