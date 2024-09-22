package main

import (
	"strings"
	"testing"
)

func TestBuildNoArgsGivenMsgContainsKeyWords(t *testing.T) {
	result := BuildNoArgsGivenMsg()

	if !strings.Contains(result, "usage") {
		t.Errorf("BuildNoArgsGivenMsg(): Does not contain 'usage'")
	}
}