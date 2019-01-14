## Using the Go flag package

**bool_flag.go** - Boolean flag usage
```
$ ./bool_flag
Dry Run Flag Enabled: false

$ ./bool_flag --dry-run
Dry Run Flag Enabled: true
```

**string_flag.go** - String flag usage
```
$ ./string_flag
Target host is empty

$ ./string_flag --target-host my_host.local
The target host is: my_host.local
```

**int_flag.go** - Int flag usage
```
$ ./int_flag --quantity 3
The quantity is: 3
```
