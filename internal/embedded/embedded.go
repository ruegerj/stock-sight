package embedded

import (
	_ "embed"
)

//go:embed db/schema.sql
var DDL string
