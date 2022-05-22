package sql

import (
	"database/sql"
	_ "database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// Test sql api
func TestSqlBasicOperator(t *testing.T) {
	// Open
	db, err := sql.Open("mysql", "root:ug3213748560@/dbname")

	if err != nil {
		t.Error("Can't open sql")
	}
	defer db.Close()

}
