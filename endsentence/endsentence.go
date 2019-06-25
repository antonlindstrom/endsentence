package endsentence

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer returns an analyzer to use to detect issues in the code.
var Analyzer = &analysis.Analyzer{
	Name: "endsentence",
	Doc:  "reports that documentation ends with a period",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if doc := nodeCommentGroup(n); doc != nil {
				if !docHasEndingPeriod(doc) {
					pass.Reportf(doc.Pos(), nodeName(n)+" comment should end with period")
					return false
				}
			}
			return true
		})
	}

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

// docHasEndingPeriod returns if the CommentGroup has an ending period.
func docHasEndingPeriod(doc *ast.CommentGroup) bool {
	return strings.HasSuffix(strings.TrimSpace(doc.Text()), ".")
}
