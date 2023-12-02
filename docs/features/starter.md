# Starter Cli

Start From a Template Project

## How

Run commands as following help:

```shell
❯ ./fluent starter --help
fluent start java,python

Usage:
  fluent starter [flags]

Examples:
fluent start java

Flags:
  -h, --help   help for starter

```

## Configure to Add New Starter Project script

Add New Action in ***starters.json*** file:
   1. name: starter flag name, if name is go , then ./fluent start go
   2. desc: describe this command 
   3. commands: the commands to run in starter action
   4. set the ***FLUENT_HOME*** as the starters.json location
   5. Add any command in following format, cli app will take it 

```json
{
  "actions": [
    {
      "name": "java",
      "desc": "Create Java Template Project",
      "commands": [
        "git clone https://github.com/fluent-qa/fluent-java-tpl"
      ]
    },
    {
      "name": "python",
      "desc": "Create Python Template Project",
      "commands": [
        "cookiecutter https://github.com/fluent-qa/fluentqa-pytpl.git --no-input"
      ]
    }
  ]
}
```