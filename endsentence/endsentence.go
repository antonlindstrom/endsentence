package endsentence

import (
	"go/ast"
	"strings"

	"github.com/mingrammer/commonregex"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"gopkg.in/jdkato/prose.v2"
)

// Analyzer returns an analyzer to use to detect issues in the code.
var Analyzer = &analysis.Analyzer{
	Name:             "endsentence",
	Doc:              "reports that documentation ends with a period",
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if doc := nodeCommentGroup(n); doc != nil {
			hasEndingPeriod, err := docHasEndingPeriod(doc)
			if err != nil {
				pass.Reportf(doc.Pos(), nodeName(n)+" "+err.Error())
				return
			}
			if !hasEndingPeriod {
				pass.Reportf(doc.Pos(), nodeName(n)+" comment should end with period")
				return
			}
		}
	})

	return nil, nil
}

// nodeName returns the name of the node.
func nodeName(n ast.Node) string {
	switch v := n.(type) {
	case *ast.FuncDecl:
		return v.Name.Name
	case *ast.GenDecl:
		name := v.Tok.String()
		for _, g := range v.Specs {
			if g, ok := g.(*ast.TypeSpec); ok {
				name = g.Name.Name
			}
		}
		return name
	default:
		return ""
	}
}

// nodeCommentGroup returns the CommentGroup for the given node. This is only
// checked for FuncDecl and GenDecl. Returns nil if it does not exist.
func nodeCommentGroup(n ast.Node) *ast.CommentGroup {
	switch v := n.(type) {
	case *ast.FuncDecl:
		return v.Doc
	case *ast.GenDecl:
		return v.Doc
	default:
		return nil
	}
}

// docHasEndingPeriod returns if the CommentGroup has an ending period. This
// function uses the package prose to be able to make good use of splitting
// the comment into sentences (if there are multiple).
func docHasEndingPeriod(doc *ast.CommentGroup) (bool, error) {
	pdoc, err := prose.NewDocument(
		doc.Text(),
		prose.WithTokenization(false),
		prose.WithTagging(false),
		prose.WithExtraction(false),
	)
	if err != nil {
		return false, err
	}

	for _, s := range pdoc.Sentences() {
		if !strings.HasSuffix(strings.TrimSpace(s.Text), ".") {
			if !endsWithURL(s.Text) && !endsWithEmail(s.Text) && !isList(s) {
				return false, nil
			}
		}
	}
	return true, nil
}

// listPrefixes contains the valid prefixes that can be used for a list. For
// instance, the following are list prefixes:
// * A list prefix
// - Another list prefix
var listPrefixes = []string{
	"-",
	"*",
}

// isList is a way to check if the given sentence is a list by checking the
// prefix of the sentence.
func isList(s prose.Sentence) bool {
	for _, p := range listPrefixes {
		if strings.HasPrefix(s.Text, p) {
			return true
		}
	}
	return false
}

// endsWithURL checks if the string ends with a URL.
func endsWithURL(s string) bool {
	words := strings.Split(s, " ")
	return commonregex.LinkRegex.MatchString(words[len(words)-1])
}

// endsWithEmail checks if the string ends with an email.
func endsWithEmail(s string) bool {
	words := strings.Split(s, " ")
	return commonregex.EmailRegex.MatchString(words[len(words)-1])
}
