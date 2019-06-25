package endsentence

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestEndsWithPeriod(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
