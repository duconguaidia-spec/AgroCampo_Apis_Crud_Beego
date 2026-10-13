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
	m.SQL("INSERT INTO veterinarias.especialidad (nombre_especialidad, activo) VALUES ('Medicina General', true);")
	m.SQL("INSERT INTO veterinarias.especialidad (nombre_especialidad, activo) VALUES ('Cirugía', true);")
	m.SQL("INSERT INTO veterinarias.especialidad (nombre_especialidad, activo) VALUES ('Dermatología', true);")
	m.SQL("INSERT INTO veterinarias.especialidad (nombre_especialidad, activo) VALUES ('Cardiología', true);")
	m.SQL("INSERT INTO veterinarias.especialidad (nombre_especialidad, activo) VALUES ('Neurología', true);")

}

// Reverse the migrations
func (m *InsertTableEspecialidad_20260521_171318) Down() {
	m.SQL("DELETE FROM veterinarias.especialidad WHERE id IN (1, 2, 3, 4, 5);")

}
