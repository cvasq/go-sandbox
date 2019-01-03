package main

import (
	"crypto/md5"
	"fmt"
)

func md5Sum(d string) string {
	data := []byte(d)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func main() {
	data := "hello"
	fmt.Println("Data:", data)
	fmt.Println("MD5 Sum:", md5Sum(data))
}
