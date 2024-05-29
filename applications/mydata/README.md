# A Fast and Flexible export data from remote mysql server

# how to use
`go run main.go` or you can `go build -o mydata main.go && ./mydata -h`

# cmd help 

```text
Usage:
    mydata [flags]
    mydata [command]

Available Commands:
    completion  Generate the autocompletion script for the specified shell
    export      Export data from remote database server
    help        Help about any command
    hex         Convert hex string to byte
    import      Import data to remote database server
    version     Print the version number of tool

Flags:
    --debug   print stack log
    -h, --help    help for mydata

Use "mydata [command] --help" for more information about a command.
```