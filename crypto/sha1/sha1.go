package main

import "crypto/sha1"
import "fmt"

func sha1Hash(d string) string {
	h := sha1.New()
	h.Write([]byte(d))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func main() {
	data := "carlos"
	fmt.Println("DATA:", data)
	fmt.Println("SHA1 Hash:", sha1Hash(data))
}
