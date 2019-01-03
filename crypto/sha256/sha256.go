package main

import "crypto/sha256"
import "fmt"

func sha256Hash(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func main() {

	data := "carlos"

	hash := sha256Hash(data)

	fmt.Println("Data:", data)
	fmt.Printf("SHA256 Hash: %v\n", hash)
}
