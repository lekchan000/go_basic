package main

import "fmt"

type (
	Database interface {
		Insert() error
		Update() error
	}

	PostgresDb struct{} // Real Database

	MockDb struct{}
)

func (p *PostgresDb) Insert() error {
	fmt.Println("Inserting into Postgres")
	return nil
}

func (p *PostgresDb) Update() error {
	fmt.Println("Updating in Postgres")
	return nil
}

func (m *MockDb) Insert() error {
	fmt.Println("Inserting into Mock Database")
	return nil
}

func (m *MockDb) Update() error {
	fmt.Println("Updating in Mock Database")
	return nil
}

func InsertPlayerItem(db Database) error {
	return db.Insert()
}

func UpdatePlayerItem(db Database) error {
	return db.Update()
}

func main() {
	postgres := &PostgresDb{}
	mock := &MockDb{}

	InsertPlayerItem(postgres)
	InsertPlayerItem(mock)

	UpdatePlayerItem(postgres)
	UpdatePlayerItem(mock)
}
