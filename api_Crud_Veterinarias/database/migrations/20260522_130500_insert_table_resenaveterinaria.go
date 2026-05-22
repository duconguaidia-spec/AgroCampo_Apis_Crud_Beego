package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableResenaveterinaria_20260522_130500 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableResenaveterinaria_20260522_130500{}
	m.Created = "20260522_130500"

	migration.Register("InsertTableResenaveterinaria_20260522_130500", m)
}

// Run the migrations
func (m *InsertTableResenaveterinaria_20260522_130500) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableResenaveterinaria_20260522_130500) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
