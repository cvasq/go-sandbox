## Cobra CLI builder

A usage example of the Cobra CLI tool builder

```
go build .
```

```
$ ./cobra-cli -h

Usage:
  cobra-cli [flags]
  cobra-cli [command]

Available Commands:
  add-user    Add a new user
  help        Help about any command

Flags:
  -d, --dryrun    Dry run
  -h, --help      Help for cobra-cli
  -v, --verbose   Verbose output

Use "cobra-cli [command] --help" for more information about a command.
```

```
$ ./cobra-cli add-user cvasquez
[+] Added user: cvasquez
```

```
$ ./cobra-cli add-user cvasquez -d
[Dry-Run] Added user: cvasquez
```

```
$ ./cobra-cli add-user cvasquez -v    
[Verbose] Creating a new user: cvasquez
[+] Added user: cvasquez
```
```
$ ./cobra-cli add-user            
Error: requires at least 1 arg(s), only received 0
```