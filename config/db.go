package config

import (
	"database/sql"
	"fmt"
	"log"
)

func NewInstanceDatabase() *sql.DB {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("erro ao carregar configurações: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("erro ao abrir conexão com banco de dados: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("erro ao conectar com banco de dados: %v", err)
	}

	log.Println("Conexão com banco de dados estabelecida com sucesso!")
	return db
}
