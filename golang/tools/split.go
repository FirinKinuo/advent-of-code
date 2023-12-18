package tools

import "strings"

func SplitNewLines(in string) []string {
	return strings.Split(in, "\r\n")
}

func SplitSeparatedLines(in string) []string {
	return strings.Split(in, "\r\n\r\n")
}
