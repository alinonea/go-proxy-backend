package repo

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/alinonea/main/domain"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type DbTest struct {
	db       *sql.DB
	fixtures *testfixtures.Loader
}

func NewDbTest(envVars domain.EnvironmentVariables) (*DbTest, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		envVars.DB_HOST, envVars.DB_PORT, envVars.DB_USER, envVars.DB_PASSWORD, envVars.TEST_DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS requests(
		id serial PRIMARY KEY,
		request jsonb NOT NULL,
		remote_addr VARCHAR(256) NOT NULL

	)`)
	if err != nil {
		return nil, err
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),     // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("./fixtures"), // The directory containing the YAML files
	)
	if err != nil {
		return nil, err
	}

	return &DbTest{
		db:       db,
		fixtures: fixtures,
	}, nil
}

func (dbTest *DbTest) prepareDataBase() {
	if err := dbTest.fixtures.Load(); err != nil {
		log.Fatalf("Error when loading the fixtures: %v", err)
	}
}

func TestCreateConnection(t *testing.T) {
	t.Run("should return error because of a wrong connection string", func(t *testing.T) {
		envVars := domain.EnvironmentVariables{}
		_, err := CreateConnection(envVars)
		assert.Error(t, err)

	})

	t.Run("success", func(t *testing.T) {
		envVars := domain.EnvironmentVariables{
			DB_HOST:     "localhost",
			DB_PORT:     5432,
			DB_USER:     "postgres",
			DB_PASSWORD: "postgres",
			DB_NAME:     "go-proxy-test",
		}

		_, err := CreateConnection(envVars)
		assert.Nil(t, err)
	})

}

func TestCreateDB(t *testing.T) {
	envVars := domain.EnvironmentVariables{
		DB_HOST:     "localhost",
		DB_PORT:     5432,
		DB_USER:     "postgres",
		DB_PASSWORD: "postgres",
		DB_NAME:     "go-proxy-test",
	}

	db, _ := CreateConnection(envVars)

	t.Run("success", func(t *testing.T) {
		err := db.CreateDB()
		assert.Nil(t, err)
	})

	t.Run("should return error when executing the query", func(t *testing.T) {
		db.Db.Close()
		err := db.CreateDB()
		assert.Error(t, err)
	})
}
