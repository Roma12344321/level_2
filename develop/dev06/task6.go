package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type options struct {
	fields    []int
	delimiter string
	separated bool
}

func parseOptions() options {
	fieldsFlag := flag.String("f", "", "Select fields (columns)")
	delimiterFlag := flag.String("d", "\t", "Use different delimiter")
	separatedFlag := flag.Bool("s", false, "Only lines with delimiter")
	flag.Parse()
	var fields []int
	if *fieldsFlag != "" {
		fieldStrs := strings.Split(*fieldsFlag, ",")
		for _, fieldStr := range fieldStrs {
			var field int
			_, _ = fmt.Sscanf(fieldStr, "%d", &field)
			fields = append(fields, field-1)
		}
	}
	return options{
		fields:    fields,
		delimiter: *delimiterFlag,
		separated: *separatedFlag,
	}
}

func cutLine(line string, opts options) string {
	parts := strings.Split(line, opts.delimiter)
	if opts.separated && len(parts) < 2 {
		return ""
	}
	var selectedParts []string
	for _, field := range opts.fields {
		if field >= 0 && field < len(parts) {
			selectedParts = append(selectedParts, parts[field])
		}
	}
	return strings.Join(selectedParts, opts.delimiter)
}

func main() {
	opts := parseOptions()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result := cutLine(line, opts)
		if result != "" || !opts.separated {
			fmt.Println(result)
		}
	}
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
