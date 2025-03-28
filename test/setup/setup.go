package setup

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func SetupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("erro ao abrir banco de dados: %v", err)
	}

	err = RunMigrations(db)
	if err != nil {
		t.Fatalf("erro ao executar migrações: %v", err)
	}

	return db
}
