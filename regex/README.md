## Using regex in Go

**replace_string.go** - Replacing a word in a sample string

**email_validator.go** - Email address validator using regexp package

Example usage:

```
go build -o evalidator
```

```
$ ./evalidator 12345           
Not a valid email address:  12345
```

```
$ ./evalidator testinggmail.com
Not a valid email address:  testinggmail.com
```

```
$ ./evalidator test@gmail.com  
Valid email address:  test@gmail.com
```


