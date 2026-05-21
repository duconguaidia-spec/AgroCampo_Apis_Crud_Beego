package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertSubasta_20260521_143441 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertSubasta_20260521_143441{}
	m.Created = "20260521_143441"

	migration.Register("InsertSubasta_20260521_143441", m)
}

// Run the migrations
func (m *InsertSubasta_20260521_143441) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO subasta (nombre, descripcion) VALUES ('Suba Casanare', 'Subasta de Yopal');")

}

// Reverse the migrations
func (m *InsertSubasta_20260521_143441) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM subasta WHERE nombre = 'Suba Casanare';")

}
