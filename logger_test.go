package main

import (
	"testing"
)

func TestLoggerFormat(t *testing.T) {
	msgs := formatMsg("@c", "hello ", "world")
	if msgs[0] != "@chello " {
		t.Errorf("Expected @chello , got %v", msgs[0])
	}
}
