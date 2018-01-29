package linter

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestEndsWithPeriod(t *testing.T) {
	tests := []struct {
		name        string
		source      string
		feedbackLen int
	}{
		{
			name: "funcdoc with period in the end",
			source: `
				package main

				// TestWithPeriod has a period in the end.
				func TestWithPeriod() {}
			`,
			feedbackLen: 0,
		},
		{
			name: "funcdoc without period in the end",
			source: `
				package main

				// TestWithPeriod has a period in the end
				func TestWithPeriod() {}
			`,
			feedbackLen: 1,
		},
		{
			name: "funcdoc with period in the end and space in the middle",
			source: `
				package main

				// TestWithPeriod has a period in the end.
				//
				// This comment also has another line.
				func TestWithPeriod() {}
			`,
			feedbackLen: 0,
		},
		{
			name: "funcdoc without period in the end and space in the middle",
			source: `
				package main

				// TestWithPeriod has a period in the end.
				//
				// This comment also has another line
				func TestWithPeriod() {}
			`,
			feedbackLen: 1,
		},
		{
			name: "funcdoc with period in the end without first period and space in the middle",
			source: `
				package main

				// TestWithPeriod has a period in the end
				//
				// This comment also has another line.
				func TestWithPeriod() {}
			`,
			feedbackLen: 0,
		},
		{
			name: "multiline funcdoc with period in the end",
			source: `
				package main

				// TestWithPeriod that is multi line and has a
				// period in the end.
				func TestWithPeriod() {}
			`,
			feedbackLen: 0,
		},
		{
			name: "multiline funcdoc without period in the end",
			source: `
				package main

				// TestWithPeriod that is multi line and has
				// no period in the end
				func TestWithPeriod() {}
			`,
			feedbackLen: 1,
		},
		{
			name: "typedoc with period in the end",
			source: `
				package main

				// TestWithPeriod has a period in the end.
				type TestWithPeriod string
			`,
			feedbackLen: 0,
		},
		{
			name: "typedoc without period in the end",
			source: `
				package main

				// TestWithPeriod has a period in the end
				type TestWithPeriod string
			`,
			feedbackLen: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fs := token.NewFileSet()
			node, err := parser.ParseFile(fs, "testing.go", test.source, parser.ParseComments)
			if err != nil {
				t.Errorf("want err = nil, got %v", err)
				return
			}

			f := &file{
				filename: "testing.go",
				astFile:  node,
				fileSet:  fs,
			}

			feedback := f.checkDoc()
			if len(feedback) != test.feedbackLen {
				t.Errorf("want len(feedback) = %v, got %v", test.feedbackLen, len(feedback))
			}
		})
	}
}
