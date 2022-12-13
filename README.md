## Installation

```bash
brew tap monlor/taps
brew install monlor/taps/spug-cli
spug-cli version
```

## How to use

```
Spug command line tool, support application release, log query, audit and so on

Usage:
  spug-cli [command]

Examples:
  spug-cli login
  spug-cli publish
  spug-cli publish -e dev -a job
  spug-cli publish -e dev -a base -v dev-latest -w
  spug-cli logs 6634
  spug-cli status 6634

Available Commands:
  approve     Approve spug apply
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List spug apply
  login       Login Spug
  logs        Show apply logs
  publish     Publish your application
  status      Show apply status
  version     Print the version number

Flags:
  -h, --help   help for spug-cli

Use "spug-cli [command] --help" for more information about a command.
```