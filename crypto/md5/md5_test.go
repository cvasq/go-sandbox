package main

import "testing"

func TestMD5(t *testing.T) {
	v := md5Sum("hello")
	if v != "5d41402abc4b2a76b9719d911017c592" {
		t.Error("Expected another hash, got ", v)
	}
}
