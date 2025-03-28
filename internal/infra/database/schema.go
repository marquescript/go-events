package database

import (
	"database/sql"
	"embed"
	"fmt"
	"strings"
)

//go:embed schema.sql
var schema embed.FS

func RunMigrations(db *sql.DB) error {
	content, err := schema.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo schema.sql: %w", err)
	}

	// Divide o conteúdo em comandos SQL individuais
	commands := strings.Split(string(content), ";")

	// Executa cada comando separadamente
	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		_, err = db.Exec(cmd)
		if err != nil {
			return fmt.Errorf("erro ao executar comando SQL: %s\nErro: %w", cmd, err)
		}
	}

	return nil
}
