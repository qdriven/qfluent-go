# Setup Golang Cli Project

- [cobra](https://cobra.dev/)

## Install and init Project

```shell
go get -u github.com/spf13/cobra/cobra
```


## Main Concepts

- Commands: Actions
- Args: Things
- Flags: modifiers of actions

Command Patter:
```shell
 APPNAME VERB NOUN --ADJECTIVE
 # or 
 APPNAME COMMAND ARG --FLAG
```

## Commands

```sh
go get -u github.com/spf13/cobra/cobra
go install github.com/spf13/cobra-cli@latest
```

- hugo server --port=1313
- cobra init appname
- cobra add cmdname
- 

## Cli Project Generator

```shell

```