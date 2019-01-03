package main

import "testing"

func TestSha(t *testing.T) {
	v := sha256Hash("carlos")
	if v != "7b85175b455060e3237e925f023053ca9515e8682a83c8b09911c724a1f8b75f" {
		t.Error("Expected another hash, got ", v)
	}
}
