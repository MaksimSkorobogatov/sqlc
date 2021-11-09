package compiler

import (
	"strings"
	"testing"

	"github.com/MaksimSkorobogatov/sqlc/internal/engine/dolphin"
)

func Test_findParameters_withOrderByCase(t *testing.T) {
	const sqlSrc = `
SELECT *
FROM items
ORDER BY CASE ?
             WHEN 'type' THEN type
             WHEN 'name' THEN name
             WHEN 'created_at' THEN created_at
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
}
