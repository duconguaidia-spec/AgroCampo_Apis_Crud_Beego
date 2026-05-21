package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTrVeterinaria_20260521_171617 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTrVeterinaria_20260521_171617{}
	m.Created = "20260521_171617"

	migration.Register("InsertTableTrVeterinaria_20260521_171617", m)
}

// Run the migrations
func (m *InsertTableTrVeterinaria_20260521_171617) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableTrVeterinaria_20260521_171617) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
