package main

import "fmt"

var NC = "\033[0m"

func green(text string) string {
	GREEN := "\033[0;32m"
	return fmt.Sprintf("%v%v%v", GREEN, text, NC)
}

func red(text string) string {
	GREEN := "\033[0;31m"
	return fmt.Sprintf("%v%v%v", GREEN, text, NC)
}
