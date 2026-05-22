package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTrVeterinariaservicio_20260522_130609 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTrVeterinariaservicio_20260522_130609{}
	m.Created = "20260522_130609"

	migration.Register("InsertTableTrVeterinariaservicio_20260522_130609", m)
}

// Run the migrations
func (m *InsertTableTrVeterinariaservicio_20260522_130609) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableTrVeterinariaservicio_20260522_130609) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
