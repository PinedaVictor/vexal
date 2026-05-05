/*
Copyright © 2026 Victor Pineda pinedavictor095@gmail.com
*/
package config

import (
	"fmt"
	"os"
	"strings"
)

const claudeMDTemplate = `## Dependency & Context Tools (vexal)
This repo has %d indexed dependency edges.

- ` + "`vx overview`" + ` — repo structure at a glance: top dependents, file count, edge count
- ` + "`vx deps <file>`" + ` — what a file imports and what depends on it
- ` + "`vx impact <file>`" + ` — blast radius: all files affected if this file changes

When to use:
- Orienting to an unfamiliar codebase → run ` + "`vx overview`" + ` first
- Planning a refactor or API change → run ` + "`vx impact <file>`" + ` before editing
- Investigating an unknown file → run ` + "`vx deps <file>`" + `

Skip for routine edits to files you already understand well.

If the snapshot is stale, run ` + "`vx init`" + ` to rebuild it.
`

const claudeMD = "CLAUDE.md"

// InitClaudeMD writes the vexal section to .vexal/CLAUDE.md with the current
// edge count embedded. Always overwrites — this file is fully generated.
func InitClaudeMD(edgeCount int) error {
	curDir, _ := os.Getwd()
	claudePath := fmt.Sprintf("%s/.vexal/%s", curDir, claudeMD)
	content := fmt.Sprintf(claudeMDTemplate, edgeCount)
	return os.WriteFile(claudePath, []byte(content), 0644)
}

const rootClaudeMDImport = "@.vexal/CLAUDE.md"

// InjectRootClaudeMD ensures the root CLAUDE.md imports the vexal-generated file.
// Creates root CLAUDE.md if absent. No-ops if the import is already present.
func InjectRootClaudeMD() error {
	curDir, _ := os.Getwd()
	rootPath := fmt.Sprintf("%s/%s", curDir, claudeMD)

	data, err := os.ReadFile(rootPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if strings.Contains(string(data), rootClaudeMDImport) {
		return nil
	}

	if os.IsNotExist(err) {
		return os.WriteFile(rootPath, []byte(rootClaudeMDImport+"\n"), 0644)
	}

	f, err := os.OpenFile(rootPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "\n%s\n", rootClaudeMDImport)
	return err
}
