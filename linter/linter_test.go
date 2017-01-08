package linter

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestEndsWithPeriod(t *testing.T) {
	tests := []struct {
		source      string
		feedbackLen int
	}{
		{
			source: `
				package main

				// TestWithPeriod has a period in the end.
				func TestWithPeriod() {}
			`,
			feedbackLen: 0,
		},
		{
			source: `
				package main

				// TestWithPeriod has a period in the end
				func TestWithPeriod() {}
			`,
			feedbackLen: 1,
		},
	}

	for _, test := range tests {
		fs := token.NewFileSet()
		node, err := parser.ParseFile(fs, "testing.go", test.source, parser.ParseComments)
		if err != nil {
			t.Error("want err = nil, got %v", err)
		}

		f := &file{
			filename: "testing.go",
			astFile:  node,
			fileSet:  fs,
		}

		feedback := f.checkDoc()
		if len(feedback) != test.feedbackLen {
			t.Error("want len(feedback) = 0, got %v", len(feedback))
		}
	}
}
