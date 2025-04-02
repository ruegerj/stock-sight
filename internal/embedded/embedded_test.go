package embedded

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDDL_ContainsDbSchema(t *testing.T) {
	schemaPath := filepath.Join("db", "schema.sql")
	currentSqlSchema, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, string(currentSqlSchema))
	assert.Equal(t, string(currentSqlSchema), DDL)
}
