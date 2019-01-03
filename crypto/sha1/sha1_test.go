package main

import "testing"

func TestSha(t *testing.T) {
	v := sha1Hash("carlos")
	if v != "ab5e2bca84933118bbc9d48ffaccce3bac4eeb64" {
		t.Error("Expected another hash, got ", v)
	}
}
