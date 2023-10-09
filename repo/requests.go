package repo

import (
	"database/sql"
	"fmt"

	domain "github.com/alinonea/main/domain"
)

type RequestRepositoryInterface interface {
	SaveRequest(request domain.Request) error
}

type RequestRepository struct {
	db *sql.DB
}

func NewRequestRepository(db *sql.DB) RequestRepositoryInterface {
	return &RequestRepository{
		db: db,
	}
}

func (repo *RequestRepository) SaveRequest(request domain.Request) error {
	_, err := repo.db.Exec(`INSERT INTO requests (request, remote_addr) VALUES ($1,$2)`, request.RequestBody, request.RemoteAddr)
	if err != nil {
		return fmt.Errorf("Error while saving the request in the db: %v", err)
	}

	return nil
}
