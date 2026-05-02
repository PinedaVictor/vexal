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

- ` + "`vx deps <file>`" + ` — show what a file imports and what depends on it
- ` + "`vx impact <file>`" + ` — show what files would be affected by changing this file

Always run ` + "`vx impact`" + ` before refactoring and ` + "`vx deps`" + ` before
editing a file you haven't read yet.

If the snapshot is stale, run ` + "`vx init`" + ` to rebuild it.
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
