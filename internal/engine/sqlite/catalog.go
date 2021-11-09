package sqlite

import "github.com/MaksimSkorobogatov/sqlc/internal/sql/catalog"

func NewCatalog() *catalog.Catalog {
	c := catalog.New("main")
	return c
}
