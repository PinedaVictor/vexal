# vexal

Developer workflow CLI with dependency graph, AI tooling, and repo automation.

## Installation

```bash
npm install -g vexal
```

## Usage

```bash
vx [command]
```

### Commands

- `vx overview` — repo structure at a glance: top dependents, file count, edge count
- `vx deps <file>` — what a file imports and what depends on it
- `vx impact <file>` — blast radius: all files affected if this file changes
- `vx init` — initialize vexal and build the dependency graph
- `vx pr` — AI-assisted PR generation based on your commit history
- `vx todos` — find all TODO comments in your codebase
- `vx fixme` — find all FIXME comments in your codebase

## Why vexal

Tested on Excalidraw, a real well-maintained open-source monorepo:

```
$ grep -r "@excalidraw/common" .
# 6 callers

$ vx impact packages/excalidraw/common/index.ts
# 290 files affected
```

284 files completely missed — they import via `@excalidraw/common`, a path alias grep can't follow. A dev relying on grep would have shipped a breaking change across 284 files they didn't know existed.

## Links

- [Website](https://www.vexal.io/)
- [Quickstart](https://www.vexal.io/quickstart)
- [GitHub](https://github.com/PinedaVictor/vexal)
