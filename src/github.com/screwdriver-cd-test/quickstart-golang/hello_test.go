package hello

import (
	"strings"
	"testing"
)

func TestGreetings(t *testing.T) {
	expected := "Hello, Go"
	result := greetings()

	if !strings.EqualFold(result, expected) {
		t.Errorf("result == %q, want %q", result, expected)
	}
}

func TestMain(t *testing.T) {
	expected := "Hi, this is test"
	result := echo(expected)

	if !strings.EqualFold(result, expected) {
		t.Errorf("result == %q, want %q", result, expected)
	}
}
