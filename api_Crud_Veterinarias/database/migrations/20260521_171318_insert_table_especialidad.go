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
	m.SQL("INSERT INTO veterinarias.especialidad (id, nombre) VALUES (1, 'Medicina General');")
	m.SQL("INSERT INTO veterinarias.especialidad (id, nombre) VALUES (2, 'Cirugía');")
	m.SQL("INSERT INTO veterinarias.especialidad (id, nombre) VALUES (3, 'Dermatología');")
	m.SQL("INSERT INTO veterinarias.especialidad (id, nombre) VALUES (4, 'Cardiología');")
	m.SQL("INSERT INTO veterinarias.especialidad (id, nombre) VALUES (5, 'Neurología');")

}

// Reverse the migrations
func (m *InsertTableEspecialidad_20260521_171318) Down() {
	m.SQL("DELETE FROM veterinarias.especialidad WHERE id IN (1, 2, 3, 4, 5);")

}
