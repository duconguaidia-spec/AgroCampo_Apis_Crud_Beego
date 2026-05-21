package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTrVeterinariaservicio_20260521_171611 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTrVeterinariaservicio_20260521_171611{}
	m.Created = "20260521_171611"

	migration.Register("InsertTableTrVeterinariaservicio_20260521_171611", m)
}

// Run the migrations
func (m *InsertTableTrVeterinariaservicio_20260521_171611) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableTrVeterinariaservicio_20260521_171611) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
