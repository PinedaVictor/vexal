package scraper

import (
	"fmt"
	"regexp"
	"strings"
)

func ScrapeTodos() {
	findPattern(todoPattern, "todos.md")
}

func ScrapeFixMe() {
	findPattern(fixMePattern, "fixme.md")
}

func findPattern(pattern *regexp.Regexp, outputFile string) {
	todos := scrape(currentDirectory, fileExtensions, ignoredDirectories, pattern, outputFile)
	if len(todos) > 0 {
		generateReadme(todos, outputFile)
	} else {
		fmt.Printf("No comments found in %s files.\n", strings.Join(fileExtensions, ", "))
	}
}
