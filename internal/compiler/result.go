package compiler

import (
	"github.com/MaksimSkorobogatov/sqlc/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
