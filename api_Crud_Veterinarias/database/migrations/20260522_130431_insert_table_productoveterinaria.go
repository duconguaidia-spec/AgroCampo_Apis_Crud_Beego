package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableProductoveterinaria_20260522_130431 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableProductoveterinaria_20260522_130431{}
	m.Created = "20260522_130431"

	migration.Register("InsertTableProductoveterinaria_20260522_130431", m)
}

// Run the migrations
func (m *InsertTableProductoveterinaria_20260522_130431) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *InsertTableProductoveterinaria_20260522_130431) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
