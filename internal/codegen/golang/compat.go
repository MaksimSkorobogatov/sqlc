package golang

import (
	"github.com/MaksimSkorobogatov/sqlc/internal/core"
	"github.com/MaksimSkorobogatov/sqlc/internal/sql/ast"
)

func sameTableName(n *ast.TableName, f core.FQN, defaultSchema string) bool {
	if n == nil {
		return false
	}
	schema := n.Schema
	if n.Schema == "" {
		schema = defaultSchema
	}
	return n.Catalog == f.Catalog && schema == f.Schema && n.Name == f.Rel
}
