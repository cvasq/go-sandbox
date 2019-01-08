## Using the Go flag package

**bool_flag.go** - Boolean flag usage
```
$ go run bool_flag.go
Dry Run Flag Enabled: false

$ go run bool_flag.go --dry-run
Dry Run Flag Enabled: true
```

**string_flag.go** - String flag usage
```
$ go run string_flag.go
Target host is empty

$ go run string_flag.go --target-host my_host.local
The target host is: my_host.local
```
