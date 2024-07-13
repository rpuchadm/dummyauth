package ini

import (
	"database/sql"
	"fmt"
	"log"
)

func DropTableAplicaciones(db *sql.DB) (string, error) {
	dropTableSQL := `DROP TABLE IF EXISTS APLICACIONES;`

	_, err := db.Exec(dropTableSQL)
	if err != nil {
		msg := fmt.Sprintf("Error dropping table: %s", err)
		log.Println(msg)
		return msg, err
	}
	log.Println("Table 'APLICACIONES' dropped successfully")
	return "Table 'APLICACIONES' dropped successfully", nil
}

func CreateTableAplicaciones(db *sql.DB) (string, error) {
	// Consulta para verificar si la tabla ya existe
	query := `SELECT EXISTS (
		SELECT 1
		FROM information_schema.tables
		WHERE table_name = 'APLICACIONES'
	);`

	var exists bool
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		msg := fmt.Sprintf("Error checking if table exists: %s", err)
		log.Println(msg)
		return msg, err
	}

	// Si la tabla no existe, entonces la crea
	if !exists {
		createTableSQL := `CREATE TABLE APLICACIONES (
			id SERIAL PRIMARY KEY,
			nombre VARCHAR(50) NOT NULL,
			token VARCHAR(32) NOT NULL,
			token_old VARCHAR(32),
			token_last_update TIMESTAMP
		);`

		_, err = db.Exec(createTableSQL)
		if err != nil {
			msg := fmt.Sprintf("Error creating table: %s", err)
			log.Println(msg)
			return msg, err
		}
		log.Println("Table 'APLICACIONES' created successfully")
		return "Table 'APLICACIONES' created successfully", nil
	}

	log.Println("Table 'APLICACIONES' already exists, skipping creation")
	return "Table 'APLICACIONES' already exists, skipping creation", nil
}
