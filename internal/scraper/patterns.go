package scraper

import (
	"fmt"
	"regexp"
)

func ScrapeTodos() {
	findPattern(todoPattern, "todos.md")
}

func ScrapeFixMe() {
	findPattern(fixMePattern, "fixme.md")
}

func findPattern(pattern *regexp.Regexp, outputFile string) {
	results := scrape(currentDirectory, fileExtensions, ignoredDirectories, pattern, outputFile)
	if len(results) > 0 {
		generateReadme(results, outputFile)
	} else {
		fmt.Printf("No comments found. \n")
	}
}

func FindPatterns() []Findings {
	return scrape(currentDirectory, fileExtensions, ignoredDirectories, todoPattern, "nil")
}

func ClearTodos() {
	deletePattern(currentDirectory, fileExtensions, ignoredDirectories, todoPattern)
}
