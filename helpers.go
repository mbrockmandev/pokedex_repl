package main

import "strings"

func normalizeInput(text string) []string {
	result := strings.ToLower(text)
	words := strings.Fields(result)
	return words
}
