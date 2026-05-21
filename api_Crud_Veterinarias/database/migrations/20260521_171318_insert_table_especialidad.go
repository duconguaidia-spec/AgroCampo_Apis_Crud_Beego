package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableEspecialidad_20260521_171318 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableEspecialidad_20260521_171318{}
	m.Created = "20260521_171318"

	migration.Register("InsertTableEspecialidad_20260521_171318", m)
}

// Run the migrations
func (m *InsertTableEspecialidad_20260521_171318) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableEspecialidad_20260521_171318) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
