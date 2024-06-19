package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test_input.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer func() {
		_ = os.Remove(tmpfile.Name())
	}()
	expectedLines := []string{"line 1", "line 2", "line 3"}
	err = writeLines(expectedLines, tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	lines, err := readLines(tmpfile.Name())
	if err != nil {
		t.Errorf("readLines failed: %v", err)
	}
	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("Expected %v, but got %v", expectedLines, lines)
	}
}

func TestWriteLines(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test_output.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tmpfile.Name())
	lines := []string{"line 1", "line 2", "line 3"}
	err = writeLines(lines, tmpfile.Name())
	if err != nil {
		t.Errorf("writeLines failed: %v", err)
	}

	writtenLines, err := readLines(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read temporary file: %v", err)
	}
	if !reflect.DeepEqual(writtenLines, lines) {
		t.Errorf("Expected %v, but got %v", lines, writtenLines)
	}
}

func TestUniqueLines(t *testing.T) {
	lines := []string{"line 1", "line 2", "line 1", "line 3", "line 2"}
	expected := []string{"line 1", "line 2", "line 3"}
	result := uniqueLines(lines)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSortLines(t *testing.T) {
	lines := []string{"b 2", "a 1", "c 3"}
	expected := []string{"a 1", "b 2", "c 3"}
	opts := options{column: 1, numeric: false, reverse: false, unique: false}
	result := sortLines(lines, opts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSortLinesNumeric(t *testing.T) {
	lines := []string{"b 2", "a 10", "c 3"}
	expected := []string{"b 2", "c 3", "a 10"}
	opts := options{column: 1, numeric: true, reverse: false, unique: false}
	result := sortLines(lines, opts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSortLinesReverse(t *testing.T) {
	lines := []string{"b 2", "a 1", "c 3"}
	expected := []string{"c 3", "b 2", "a 1"}
	opts := options{column: 1, numeric: false, reverse: true, unique: false}
	result := sortLines(lines, opts)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCompareLines(t *testing.T) {
	a := "a 1"
	b := "b 2"
	opts := options{column: 1, numeric: false, reverse: false, unique: false}
	result := compareLines(a, b, opts)
	if !result {
		t.Errorf("Expected true, but got false")
	}
	opts.numeric = true
	result = compareLines(a, b, opts)
	if !result {
		t.Errorf("Expected true, but got false")
	}
	a = "a 10"
	b = "b 2"
	result = compareLines(a, b, opts)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestReverseLines(t *testing.T) {
	lines := []string{"a", "b", "c"}
	expected := []string{"c", "b", "a"}
	reverseLines(lines)
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %v, but got %v", expected, lines)
	}
}
