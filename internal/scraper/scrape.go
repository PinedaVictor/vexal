package scraper

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	currentDirectory, _ = os.Getwd()
	fileExtensions      = []string{
		".py",         // Python
		".ts",         // TypeScript
		".tsx",        // TypeScript with JSX
		".js",         // JavaScript
		".jsx",        // JavaScript with JSX
		".go",         // Go
		".java",       // Java
		".cs",         // C#
		".cpp",        // C++
		".c",          // C
		".rb",         // Ruby
		".php",        // PHP
		".html",       // HTML
		".css",        // CSS
		".swift",      // Swift
		".kt",         // Kotlin
		".rs",         // Rust
		".scala",      // Scala
		".sql",        // SQL
		".md",         // Markdown
		".sh",         // Shell script
		".xml",        // XML
		".json",       // JSON
		".yaml",       // YAML
		".yml",        // YAML
		".dockerfile", // Dockerfile
		".gradle",     // Gradle
		".makefile",   // Makefile
	}
	ignoredDirectories = []string{
		"node_modules",
		"dist",
		"env",
		"instance",
		"__pycache__",
		".next",
		".ruff_cache",
		"scripts",
		"todos.md",
	}

	todoPattern  = regexp.MustCompile(`(?i)\s*//\s*TODO\b.*$`)
	fixMePattern = regexp.MustCompile(`(?i)\s*//\s*FIXME\b.*$`)
)

type Findings struct {
	Comment    string
	FilePath   string
	LineNumber int
}

func scrape(directory string, extensions []string, ignoredDirs []string, pattern *regexp.Regexp, outputFileName string) []Findings {
	var findings []Findings

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() && contains(ignoredDirs, info.Name()) {
			fmt.Println("Skipping:", info.Name())
			return filepath.SkipDir
		}

		// Skip the output file (e.g., todos.md)
		if !info.IsDir() && info.Name() == outputFileName {
			fmt.Println("Skipping file:", info.Name())
			return nil
		}

		if !info.IsDir() && hasValidExtension(info.Name(), extensions) {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			lineNumber := 0
			for scanner.Scan() {
				lineNumber++
				line := scanner.Text()
				if pattern.MatchString(line) {
					findings = append(findings, Findings{
						Comment:    strings.TrimSpace(pattern.FindString(line)),
						FilePath:   path,
						LineNumber: lineNumber,
					})
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}

	return findings
}

func deletePattern(directory string, extensions []string, ignoredDirs []string, pattern *regexp.Regexp) error {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && contains(ignoredDirs, info.Name()) {
			return filepath.SkipDir
		}

		if !info.IsDir() && hasValidExtension(info.Name(), extensions) {
			input, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			lines := strings.Split(string(input), "\n")
			for i, line := range lines {
				if pattern.MatchString(line) {
					lines[i] = pattern.ReplaceAllString(line, "")
				}
			}

			output := strings.Join(lines, "\n")
			err = os.WriteFile(path, []byte(output), info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func hasValidExtension(filename string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func generateReadme(findings []Findings, outputFile string) {
	// Create or truncate the file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating README file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fileTitle := strings.Split(outputFile, ".")[0]
	ft := strings.ToUpper(fileTitle)
	writer.WriteString(fmt.Sprintf("# %s \n\n", ft))
	for _, comment := range findings {
		relativePath, _ := filepath.Rel(currentDirectory, comment.FilePath)
		writer.WriteString(fmt.Sprintf("**%s:%d** - %s\n\n", relativePath, comment.LineNumber, comment.Comment))
	}
	writer.Flush()
}
