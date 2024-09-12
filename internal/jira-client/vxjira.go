package jiraclient

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"vx/config"
	"vx/internal/scraper"
)

// ALG:
// 1. Get Issue Type IDs
// 2. Scrape codebase
// 3. Use Go concurrency to update Jira per issueType
// ---3.1 How will Vexal make sure NOT to create dublicates????
// ---3.2 Delete TODO comments
// 4. Clear comments
func PushTodos() {
	currentDir, _ := os.Getwd()
	issueTypeID, _, _ := parseIssueTypes("todo")
	cfg, _ := config.LoadRepoConfig()
	todoPattern := regexp.MustCompile(`^\s*//\s*TODO:\s*`)
	var wg sync.WaitGroup
	pattern := scraper.FindPatterns()
	for _, value := range pattern {
		wg.Add(1)
		relativePath, _ := filepath.Rel(currentDir, value.FilePath)
		updateComment := todoPattern.ReplaceAllString(value.Comment, "")
		description := formatDescription(fmt.Sprintf("%s/%s", cfg.Repo, relativePath), value.LineNumber, updateComment)
		go func() {
			defer wg.Done()
			CreateIssue(issueTypeID, cfg.Jira_Project_key, updateComment, description)
		}()
	}
	wg.Wait()
	// TODO: Delete TODO comments here

}

func formatDescription(file string, lineNumber int, comment string) string {
	return fmt.Sprintf("File: %s \nLine Number: %d \nComment: %s \n ", file, lineNumber, comment)
}
