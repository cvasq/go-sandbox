package main

import (
	"strings"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if !strings.Contains(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Hello World message is returned", func(t *testing.T) {
		got := helloWorld()
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

}
