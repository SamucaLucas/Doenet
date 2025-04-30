package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

func main() {
	// Configuração de conexão com o PostgreSQL
	connStr := "user=samucael password=2004 dbname=casa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	defer db.Close()

	// Consulta para listar tabelas no esquema public
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Erro ao executar a consulta:", err)
	}
	defer rows.Close()

	fmt.Println("Tabelas no banco de dados:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal("Erro ao ler os resultados:", err)
		}
		fmt.Println("- ", tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Erro ao iterar pelas linhas:", err)
	}
}