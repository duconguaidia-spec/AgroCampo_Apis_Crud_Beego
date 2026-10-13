package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableServicioveterinaria_20260522_130523 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableServicioveterinaria_20260522_130523{}
	m.Created = "20260522_130523"

	migration.Register("InsertTableServicioveterinaria_20260522_130523", m)
}

// Run the migrations
func (m *InsertTableServicioveterinaria_20260522_130523) Up() {
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id_veterinaria, nombre_servicio, descripcion, precio, activo) VALUES (1, 'Consulta General', 'Consulta médica general', 50000, TRUE);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id_veterinaria, nombre_servicio, descripcion, precio, activo) VALUES (1, 'Vacunación', 'Aplicación de vacunas', 30000, TRUE);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id_veterinaria, nombre_servicio, descripcion, precio, activo) VALUES (2, 'Desparasitación', 'Tratamiento contra parásitos', 40000, TRUE);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id_veterinaria, nombre_servicio, descripcion, precio, activo) VALUES (3, 'Corte de Uñas', 'Corte y limpieza de uñas', 20000, TRUE);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id_veterinaria, nombre_servicio, descripcion, precio, activo) VALUES (4, 'Limpieza Dental', 'Limpieza profesional de dientes', 60000, TRUE);")

}

// Reverse the migrations
func (m *InsertTableServicioveterinaria_20260522_130523) Down() {
	m.SQL("DELETE FROM veterinarias.servicioveterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
