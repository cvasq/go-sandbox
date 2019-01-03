## Cobra CLI builder

A usage example of the Cobra CLI tool builder

```
go build -o ./testctl
```

```
$ ./testctl -h

Usage:
  testctl [flags]
  testctl [command]

Available Commands:
  add-user    Add a new user
  help        Help about any command

Flags:
  -d, --dryrun    Dry run
  -h, --help      Help for testctl
  -v, --verbose   Verbose output

Use "testctl [command] --help" for more information about a command.
```

```
$ ./testctl add-user cvasquez
[+] Added user: cvasquez
```

```
$ ./testctl add-user cvasquez -d
[Dry-Run] Added user: cvasquez
```

```
$ ./testctl add-user cvasquez -v    
[Verbose] Creating a new user: cvasquez
[+] Added user: cvasquez
```
```
$ ./testctl add-user            
Error: requires at least 1 arg(s), only received 0
```