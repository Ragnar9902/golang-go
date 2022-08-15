package main

import (
	"strings"
)

func tokenize(line string)[]string{
	
	tokens := strings.Fields(line)
	return tokens
	
}
