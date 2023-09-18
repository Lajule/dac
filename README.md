# dac

Typing terminal app

## Build

```sh
go build
```

## Release

```sh
goreleaser release
```

## Usage

```sh
./dac -h
Dac is typing training sessions program, it's help you to improve your typing skills.

Usage:
  dac [flags]
  dac [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Typing training sessions listing
  stats       Typing training sessions statistics
  version     Program version

Flags:
  -c, --closable            close on session timeout
  -d, --duration duration   duration of the training session
  -h, --help                help for dac
  -s, --statistic string    statistic to display (default "speed")

Use "dac [command] --help" for more information about a command.
```

# Input files

```sh
cat text | tr '\t\r\n' ' ' >input
```
