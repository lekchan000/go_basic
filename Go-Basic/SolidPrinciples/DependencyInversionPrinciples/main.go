package main

import "fmt"

type (
	Database interface {
		Insert()
	}

	PostgresDb struct{}
	MockDb     struct{}
)

func NewPostgresDatabase() Database {
	return &PostgresDb{}
}

func (m *MockDb) Insert() {
	fmt.Println("MockDb Insert")
}

func NewMockDatabase() Database {
	return &MockDb{}
}

func (p *PostgresDb) Insert() {
	fmt.Println("PostgresDb Insert")
}

func InsertPlayerItem(db Database) {
	db.Insert()
}

func main() {
	postgresDb := NewPostgresDatabase()
	mockDb := NewMockDatabase()

	InsertPlayerItem(postgresDb)
	InsertPlayerItem(mockDb)
}
