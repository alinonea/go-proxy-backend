package repo

import (
	"encoding/json"
	"testing"

	domain "github.com/alinonea/main/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveRequest(t *testing.T) {
	envVars := domain.EnvironmentVariables{
		DB_HOST:      "localhost",
		DB_PORT:      5432,
		DB_USER:      "postgres",
		DB_PASSWORD:  "postgres",
		TEST_DB_NAME: "go-proxy-test",
	}

	db, err := NewDbTest(envVars)
	defer db.db.Close()
	repo := NewRequestRepository(db.db)

	t.Run("should return error while saving the request", func(t *testing.T) {
		db.prepareDataBase()
		responseBody := []byte(`
	{
		"key": "value
	}
	`)
		err = repo.SaveRequest(domain.Request{
			RequestBody: (*json.RawMessage)(&responseBody),
			RemoteAddr:  "foo_test",
		})
		assert.Error(t, err, "Error while saving the request in the db: pq: invalid input syntax for type json")
	})

	t.Run("success", func(t *testing.T) {
		db.prepareDataBase()
		responseBody := []byte(`
	{
		"key": "value"
	}
	`)
		err = repo.SaveRequest(domain.Request{
			RequestBody: (*json.RawMessage)(&responseBody),
			RemoteAddr:  "foo_test",
		})
		assert.Nil(t, err)
	})

}
