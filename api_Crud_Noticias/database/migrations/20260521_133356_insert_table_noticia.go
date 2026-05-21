package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableNoticia_20260521_133356 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableNoticia_20260521_133356{}
	m.Created = "20260521_133356"

	migration.Register("InsertTableNoticia_20260521_133356", m)
}

// Run the migrations
func (m *InsertTableNoticia_20260521_133356) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableNoticia_20260521_133356) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
