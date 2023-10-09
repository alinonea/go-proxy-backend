package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	domain "github.com/alinonea/main/domain"
)

type DB struct {
	Db *sql.DB
}

func CreateConnection(env domain.EnvironmentVariables) (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Succesfully connected to database")
	return &DB{
		Db: db,
	}, nil
}

func (db *DB) CreateDB() error {
	_, err := db.Db.Exec(`CREATE TABLE IF NOT EXISTS requests(
		id serial PRIMARY KEY,
		request jsonb NOT NULL,
		remote_addr VARCHAR(256) NOT NULL
)`)
	if err != nil {
		return err
	}

	return nil
}
