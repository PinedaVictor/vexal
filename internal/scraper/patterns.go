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

func FindTodos() []Findings {
	return scrape(currentDirectory, fileExtensions, ignoredDirectories, todoPattern, "todos.md")
}

func FindFixme() []Findings {
	return scrape(currentDirectory, fileExtensions, ignoredDirectories, fixMePattern, "fixme.md")
}

func ClearTodos() {
	deletePattern(currentDirectory, fileExtensions, ignoredDirectories, todoPattern)
}

func ClearFixme() {
	deletePattern(currentDirectory, fileExtensions, ignoredDirectories, fixMePattern)
}
