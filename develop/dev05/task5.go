package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type options struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func parseOptions() options {
	after := flag.Int("A", 0, "Print N lines of trailing context after matching lines")
	before := flag.Int("B", 0, "Print N lines of leading context before matching lines")
	context := flag.Int("C", 0, "Print N lines of output context")
	count := flag.Bool("c", false, "Only print a count of matching lines per FILE")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Select non-matching lines")
	fixed := flag.Bool("F", false, "Interpret PATTERN as a fixed string")
	lineNum := flag.Bool("n", false, "Prefix each line of output with the line number")
	flag.Parse()
	return options{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
	}
}

func compilePattern(pattern string, opts options) (*regexp.Regexp, error) {
	if opts.fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if opts.ignoreCase {
		pattern = "(?i)" + pattern
	}
	return regexp.Compile(pattern)
}

func matchLine(line, pattern string, opts options) bool {
	re, err := compilePattern(pattern, opts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error compiling pattern: %v\n", err)
		return false
	}

	return re.MatchString(line)
}

func grepLines(lines []string, pattern string, opts options) []string {
	var result []string
	var matchCount int
	for i := 0; i < len(lines); i++ {
		match := matchLine(lines[i], pattern, opts)
		if opts.invert {
			match = !match
		}
		if match {
			matchCount++
			if opts.context > 0 {
				start := max(0, i-opts.context)
				end := min(len(lines), i+opts.context+1)
				for j := start; j < end; j++ {
					result = append(result, formatLine(lines[j], j+1, opts.lineNum))
				}
			} else {
				if opts.before > 0 {
					start := max(0, i-opts.before)
					for j := start; j < i; j++ {
						result = append(result, formatLine(lines[j], j+1, opts.lineNum))
					}
				}
				result = append(result, formatLine(lines[i], i+1, opts.lineNum))
				if opts.after > 0 {
					end := min(len(lines), i+opts.after+1)
					for j := i + 1; j < end; j++ {
						result = append(result, formatLine(lines[j], j+1, opts.lineNum))
					}
				}
			}
		}
	}
	if opts.count {
		return []string{fmt.Sprintf("%d", matchCount)}
	}
	return unique(result)
}

func formatLine(line string, lineNum int, showLineNum bool) string {
	if showLineNum {
		return fmt.Sprintf("%d:%s", lineNum, line)
	}
	return line
}

func unique(lines []string) []string {
	set := make(map[string]struct{})
	var result []string
	for _, line := range lines {
		if _, exists := set[line]; !exists {
			set[line] = struct{}{}
			result = append(result, line)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func main() {
	opts := parseOptions()
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: grep [options] <pattern> <file>")
		return
	}
	pattern := flag.Arg(0)
	inputFile := flag.Arg(1)
	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	result := grepLines(lines, pattern, opts)
	for _, line := range result {
		fmt.Println(line)
	}
}
