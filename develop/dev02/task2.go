package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func unpack(input string) (string, error) {
	var res strings.Builder
	arr := []rune(input)
	for i := 0; i < len(arr); i++ {
		if unicode.IsDigit(arr[i]) {
			return "", errors.New("invalid input: digit encountered without preceding character")
		}
		if i < len(arr)-1 && unicode.IsDigit(arr[i+1]) {
			count, err := strconv.Atoi(string(arr[i+1]))
			if err != nil {
				return "", err
			}
			for j := 0; j < count; j++ {
				res.WriteRune(arr[i])
			}
			i++
		} else {
			res.WriteRune(arr[i])
		}
	}
	return res.String(), nil
}
