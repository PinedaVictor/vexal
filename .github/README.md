<div align="center">
<img   alt="banner maker" src="https://github.com/PinedaVictor/vexal/blob/main/docs/images/vexal_banner.jpg">
</div>

<div align="center">
  <a href="https://www.vexal.io/" target="_blank" rel="noopener noreferrer">website</a> | 
  <a href="https://www.vexal.io/init" target="_blank" rel="noopener noreferrer">Init</a> | 
  <a href="https://www.vexal.io/quickstart" target="_blank" rel="noopener noreferrer">Quickstart</a> |
  <a href="https://www.vexal.io/deps" target="_blank" rel="noopener noreferrer">deps</a> |
  <a href="https://www.vexal.io/impact" target="_blank" rel="noopener noreferrer">impact</a>
</div>

## Overview

Vexal is a developer workflow CLI with a built-in dependency graph, AI tooling, and repo automation. `vx` gives you and your AI agents accurate codebase context so you can move faster with fewer surprises.

Before touching a shared file, `vx impact` tells you exactly what breaks. Before exploring an unfamiliar module, `vx deps` shows what it imports and what depends on it. Before starting a refactor, `vx overview` ranks the highest-risk files in the repo.

**grep vs. vexal — the gap is real:**

Tested on Excalidraw, a real well-maintained open-source monorepo:

```
$ grep -r "@excalidraw/common" .
# 6 callers

$ vx impact packages/excalidraw/common/index.ts
# 290 files affected
```

284 files completely missed — they import via `@excalidraw/common`, a path alias grep can't follow. A dev relying on grep would have shipped a breaking change across 284 files they didn't know existed.

**Core dependency commands:**
- `vx overview` — repo structure at a glance: top dependents, file count, edge count
- `vx deps <file>` — what a file imports and what depends on it
- `vx impact <file>` — blast radius: all files affected if this file changes

**Workflow automation:**
- Leverage AI combined with your commit history to create PRs ✅
- Keep track of all "TODO" comments in your codebase ✅
- Document and track "FIXME" comments ✅
- Configure repositories independently ✅

Follow the [Quick Start Guide](https://www.vexal.io/quickstart) to get started!

## Installation

**macOS / Linux**

```bash
brew tap PinedaVictor/vx
brew install pinedavictor/vx/vexal
```

**Windows**

```powershell
scoop bucket add vexal https://github.com/PinedaVictor/scoop-vexal
scoop install vexal
```

**npm**

```bash
npm install -g vexal
```

## Once Installed

Run

```
vx
```

You will see the following prompt

```
vexal.io - Dependency graph, AI tooling, and repo automation for developers and AI agents

Usage:
  vx [command]

Commands:
  overview    Repo structure at a glance: top dependents, file count, edge count
  deps        Show what a file imports and what depends on it
  impact      Show what files would be affected by changing the given file
  init        Initialize vexal and build the dependency graph
  pr          AI-assisted generated PRs based on your commit history
  todos       Find all "TODO:" comments in your codebase.
  fixme       Find all "FIXME:" comments in your codebase.
  config      Edit vx configuration
  context     Manage Vexal contexts for external service configuration
  enable      Enable supported API integrations by vx
  jira        Jira utils
  login       Use login to authenticate into the vexal platform
  completion  Generate the autocompletion script for the specified shell

Flags:
  -h, --help      help for vx
  -v, --version   version for vx

Use "vx [command] --help" for more information about a command.
```
