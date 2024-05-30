# A Fast and Flexible export data from remote mysql server

### how to use
`go run main.go` or you can `go build -o mydata main.go && ./mydata -h`

### cmd help

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

### How to export data
pls type help:
`./mydata exp -h`

```text
Export data from remote database server

Usage:
  mydata export [flags]

Aliases:
  export, exp

Flags:
  -a, --addr string                mysql database addr, format: ip:port
  -u, --username string            username for connect database
  -p, --password string            password for connect database
  -D, --dbname string              default database name
  -e, --query-sql string           select sql
  -o, --output string              output filename
      --fields-terminated string   fields terminated (default ",")
      --fields-enclosed string     fields enclosed
      --fields-escaped string      fields escaped (default "\\")
      --lines-terminated string    lines terminated (default "\n")
      --enclosed-optionally        fields enclosed optionally
      --params string              connection Params (default "timeout=3s")
      --buf-size int               buf size for write outfile (default 32768)
      --concurrency int            concurrency number (default 5)
      --not-merge                  merge chunks to one file
```
### example
export data by `select * from bigdata` from 127.0.0.1
```shell
./mydata exp -a 127.0.0.1:3306 -u username@tenant#clustername:clusterid -p xxx -D test -e "select * from bigdata" -o bigdata
```
