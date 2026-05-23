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
	m.SQL("INSERT INTO veterinarias.serviciogeneral (nombre_servicio, descripcion, activo) VALUES ('Consulta General', 'Consulta médica general', true);")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (nombre_servicio, descripcion, activo) VALUES ('Vacunación', 'Aplicación de vacunas', true);")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (nombre_servicio, descripcion, activo) VALUES ('Desparasitación', 'Tratamiento contra parásitos', true);")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (nombre_servicio, descripcion, activo) VALUES ('Corte de Uñas', 'Corte y limpieza de uñas', true);")
	m.SQL("INSERT INTO veterinarias.serviciogeneral (nombre_servicio, descripcion, activo) VALUES ('Limpieza Dental', 'Limpieza profesional de dientes', true);")

}

// Reverse the migrations
func (m *InsertTableServiciogeneral_20260521_171502) Down() {
	m.SQL("DELETE FROM veterinarias.serviciogeneral WHERE id IN (1, 2, 3, 4, 5);")

}
