<div align="center">
<img  width="698" alt="banner maker" src="https://github.com/PinedaVictor/vexal/blob/main/docs/images/vexal_banner.jpg">
</div>

<div align="center">
[website](https://www.vexal.io/) | [Documentation](https://www.vexal.io/init) 
</div>

## Overview

Vexal is a CLI tool written in [Go](https://go.dev/) helping developers automate tasks and improve the quality of their codebase.

- Leverage AI combined with your commit history to create PRs ✅
- Keep track of all "TODO" comments in your codebase ✅
- Document and track "FIXME" comments ✅
- Configure repositories independently ✅

Follow the [Quick Start Guide](https://www.vexal.io/quickstart) to get started configuring your repository!

## Installation

Currently, Vexal is only supported on macOS via Homebrew tap. Support for other package managers and operating systems is in the backlog.

1. Run Homebrew tap

```
brew tap PinedaVictor/vx
```

2. Install vx

```
brew install pinedavictor/vx/vexal
```

## Once Installed

Run

```
vx
```

You will see the following prompt

```
vexal.io - Developer Tools

Usage:
  vx [command]

Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Edit vx configuartion
  enable      Enable supported API integrations by vx
  fixme       Find all "FIXME:" comments in your codebase.
  help        Help about any command
  init        Initialize repository utilities. (Only needed if you plan on using github and OpenAI)
  login       Use login to authenticate into the vexal platform
  pr          AI generated PRs based on your commit history
  todos       Find all "TODO:" comments in your codebase.

Flags:
  -h, --help      help for vx
  -v, --version   version for vx

Use "vx [command] --help" for more information about a command.
```
