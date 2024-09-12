package jiraclient

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"vx/config"
	"vx/internal"
	"vx/internal/scraper"
)

func PushAllComments() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		Push("todo")
	}()
	go func() {
		defer wg.Done()
		Push("fixme")
	}()
	wg.Wait()
}

func Push(commentType string) {
	currentDir, _ := os.Getwd()
	issueTypeID, _, _ := parseIssueTypes(commentType)
	cfg, _ := config.LoadRepoConfig()
	comments, replacePattern := getCommentsAndRegexp(commentType)
	var wg sync.WaitGroup
	msg := fmt.Sprintf("Number of %s comments found: %d", commentType, len(comments))
	internal.UserFeedback(msg)
	for _, value := range comments {
		wg.Add(1)
		relativePath, _ := filepath.Rel(currentDir, value.FilePath)
		updateComment := replacePattern.ReplaceAllString(value.Comment, "")
		description := formatDescription(fmt.Sprintf("%s/%s", cfg.Repo, relativePath), value.LineNumber, updateComment)
		go func() {
			defer wg.Done()
			CreateIssue(issueTypeID, cfg.Jira_Project_key, updateComment, description)
		}()
	}
	wg.Wait()

	// FIXME: This is switch statement is used twice. Try to abstract out into 1 funciton
	switch commentType {
	case "todo":
		scraper.ClearTodos()
	case "fixme":
		scraper.ClearFixme()
	}
}

func getCommentsAndRegexp(commentType string) ([]scraper.Findings, *regexp.Regexp) {
	switch commentType {
	case "todo":
		return scraper.FindTodos(), regexp.MustCompile(`^\s*//\s*TODO:\s*`)
	case "fixme":
		return scraper.FindFixme(), regexp.MustCompile(`^\s*//\s*FIXME:\s*`)
	}
	return nil, nil
}

func formatDescription(file string, lineNumber int, comment string) string {
	return fmt.Sprintf("File: %s \nLine Number: %d \nComment: %s \n ", file, lineNumber, comment)
}
