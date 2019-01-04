## Using regex in Go

**email_validator.go** - Email address validator using regex package

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

