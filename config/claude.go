/*
Copyright © 2026 Victor Pineda pinedavictor095@gmail.com
*/
package config

import (
	"fmt"
	"os"
	"strings"
	"vx/pkg/paths"
)

const claudeMDMarker = "## Dependency & Context Tools (vexal)"

const claudeMDSection = `
## Dependency & Context Tools (vexal)
This repo uses vexal for dependency tracking. Before editing files,
use these commands to get accurate context:

- ` + "`vx deps <file>`" + ` — get imports and callers for a file
- ` + "`vx impact <file>`" + ` — check what a change will affect
- ` + "`vx map <feature>`" + ` — find all files related to a feature
- ` + "`vx contracts <file>`" + ` — see env vars, config, API contracts
- ` + "`vx snapshot`" + ` — restore session context
- ` + "`vx query \"<sql>\"`" + ` — run SQL queries over dependency data

Always run ` + "`vx impact`" + ` before refactoring and ` + "`vx deps`" + ` before
editing a file you haven't read yet.
`

const claudeMD = "CLAUDE.md"

// InitClaudeMD injects the vexal dependency tracking section into CLAUDE.md.
// Creates CLAUDE.md if it does not exist. No-ops if already injected.
func InitClaudeMD() error {
	curDir, _ := os.Getwd()
	// TODO: Revisit best user workflows at the end
	claudePath := fmt.Sprintf("%s/.vexal/%s", curDir, claudeMD)

	if paths.PathExists(claudePath) {
		content := paths.GetContent(claudePath)
		if strings.Contains(content, claudeMDMarker) {
			return nil
		}
		return paths.AppendToFile(claudePath, claudeMDSection)
	}

	return os.WriteFile(claudePath, []byte(strings.TrimLeft(claudeMDSection, "\n")), 0644)
}
