package jiraclient

import (
	"fmt"
	"regexp"
	"vx/config"
	"vx/internal/scraper"
)

// ALG:
// 1. Get Issue Type IDs
// 2. Scrape codebase
// 3. Use Go concurrency to update Jira per issueType
// ---3.1 How will Vexal make sure NOT to create dublicates????
// 4. Clear comments
func PushTodos() {
	id, _, _ := parseIssueTypes("todo")
	cfg, _ := config.LoadRepoConfig()
	todoPattern := regexp.MustCompile(`^\s*//\s*TODO:\s*`)
	pattern := scraper.FindPatterns()
	for _, value := range pattern {
		fmt.Println("issue type id", id)
		fmt.Println("Line Number:", value.LineNumber)
		fmt.Println("File:", value.FilePath)
		updateComment := todoPattern.ReplaceAllString(value.Comment, "")
		fmt.Println(updateComment)
		fmt.Println("Jira project key:", cfg.Jira_Project_key)
		fmt.Println("--------------------")
	}
	GetJiraPrjtMeta("SCRUM")

}
