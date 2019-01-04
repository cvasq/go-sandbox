## Structs in Golang

#### Example struct
```
type Address struct {
	Street  string
	Zipcode string
}
```

#### Files

**basic_struct.go** - Basic creation of a struct

```
$ go run basic_struct.go 

Name: Carlos V
Age: 26
Gender: Male
```

**nested_struct.go** - Using nested structs and promoted fields
```
$ go run nested_struct.go

Name: Robert Tappan
Age: 37
Street: 9 Main Street
Zipcode: 22303

Name: Alexandra Elbakyan
Age: 30
Street: 53 Industrial Park
Zipcode: 21520
```

**struct_constructor.go** - Using a constructer and returing a pointer
```
$ go run struct_constructor.go

Adding new student...

Name: Brian Peppers
Age: 16
ID: bgnfm0c7uh3cdv60ibf0
Points: 95
```
