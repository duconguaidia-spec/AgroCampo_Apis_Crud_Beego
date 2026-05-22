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
	m.SQL("INSERT INTO veterinarias.serviciogeneral (id, nombre) VALUES (1, 'Consulta General');")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (id, nombre) VALUES (2, 'Vacunación');")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (id, nombre) VALUES (3, 'Desparasitación');")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (id, nombre) VALUES (4, 'Corte de Uñas');")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (id, nombre) VALUES (5, 'Limpieza Dental');")

}

// Reverse the migrations
func (m *InsertTableServiciogeneral_20260521_171502) Down() {
	m.SQL("DELETE FROM veterinarias.serviciogeneral WHERE id IN (1, 2, 3, 4, 5);")

}
