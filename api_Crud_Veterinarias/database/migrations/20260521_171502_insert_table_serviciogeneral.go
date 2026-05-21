package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableServiciogeneral_20260521_171502 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableServiciogeneral_20260521_171502{}
	m.Created = "20260521_171502"

	migration.Register("InsertTableServiciogeneral_20260521_171502", m)
}

// Run the migrations
func (m *InsertTableServiciogeneral_20260521_171502) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableServiciogeneral_20260521_171502) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
