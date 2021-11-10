package compiler

import (
	"strings"
	"testing"

	"github.com/MaksimSkorobogatov/sqlc/internal/engine/dolphin"
	"github.com/MaksimSkorobogatov/sqlc/internal/sql/ast"
)

func Test_findParameters_OrderByCase(t *testing.T) {
	const sqlSrc = `
SELECT *
FROM items
ORDER BY CASE ?
             WHEN 'type' THEN type
             WHEN TRUE THEN name
             WHEN FALSE THEN created_at
             WHEN 'updated_at' THEN updated_at
             END, id
LIMIT ? OFFSET ?;
`
	parser := dolphin.NewParser()
	stmts, err := parser.Parse(strings.NewReader(sqlSrc))
	if err != nil {
		t.Fatal("parser.Parse error: ", err)
	}
	if len(stmts) < 1 {
		t.Fatal("no statements")
	}

	parameters, err := findParameters(stmts[0].Raw)
	if err != nil {
		t.Fatal("findParameters error: ", err)
	}

	const want = 3
	if got := len(parameters); got != want {
		t.Fatalf("len(parameters) = %d, want %d", got, want)
	}

	t.Logf("%#v", parameters[0].parent.(*ast.CaseExpr).Args.Items[0].(*ast.CaseWhen).Expr)

	// a := []paramRef{paramRef{
	// 	parent: (*ast.ParamRef)(0xc00003c480),
	// 	rv: (*ast.RangeVar)(nil),
	// 	ref: (*ast.ParamRef)(0xc00003c480),
	// 	name: "",
	// }, paramRef{
	// 	parent: (*limitOffset)(0x17eb8a0),
	// 	rv: (*ast.RangeVar)(nil),
	// 	ref: (*ast.ParamRef)(0xc00003c4b0),
	// 	name: "",
	// }, paramRef{
	// 	parent: (*limitCount)(0x17eb8a0),
	// 	rv: (*ast.RangeVar)(nil),
	// 	ref: (*ast.ParamRef)(0xc00003c498),
	// 	name: "",
	// }}
}
