package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTrVeterinariasespecialidad_20260522_130557 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTrVeterinariasespecialidad_20260522_130557{}
	m.Created = "20260522_130557"

	migration.Register("InsertTableTrVeterinariasespecialidad_20260522_130557", m)
}

// Run the migrations
func (m *InsertTableTrVeterinariasespecialidad_20260522_130557) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableTrVeterinariasespecialidad_20260522_130557) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
