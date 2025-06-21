package main

import (
	"fmt"
	"strings"
)

func rleLike(input string) string {
	if strings.TrimSpace(input) == "" {
		return ""
	}

	var b strings.Builder

	var current byte = input[0]
	count := 1

	//	b.WriteByte(current)
	//	b.WriteString(fmt.Sprint(count))

	for _, e := range input[1:] {
		if byte(e) == byte(current) {
			count = count + 1
			continue
		}

		b.WriteByte(current)
		b.WriteString(fmt.Sprint(count))

		current = byte(e)
		count = 1

		// early stopping
		if b.Len() > len(input) {
			return input
		}
	}

	// Dump the last character
	b.WriteByte(current)
	b.WriteString(fmt.Sprint(count))

	return b.String()
}
