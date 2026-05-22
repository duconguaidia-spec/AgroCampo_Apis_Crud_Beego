package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableVeterinaria_20260521_171617 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableVeterinaria_20260521_171617{}
	m.Created = "20260521_171617"

	migration.Register("InsertTableTrVeterinaria_20260521_171617", m)
}

// Run the migrations
func (m *InsertTableVeterinaria_20260521_171617) Up() {
	m.SQL("INSERT INTO veterinarias.veterinaria (id, nombre, direccion, telefono) VALUES (1, 'Veterinaria San Roque', 'Calle 123 #45-67', '555-1234');")
	m.SQL("INSERT INTO veterinarias.veterinaria (id, nombre, direccion, telefono) VALUES (2, 'Clínica Veterinaria El Bosque', 'Avenida 456 #78-90', '555-5678');")
	m.SQL("INSERT INTO veterinarias.veterinaria (id, nombre, direccion, telefono) VALUES (3, 'Veterinaria La Mascota Feliz', 'Carrera 789 #12-34', '555-9012');")
	m.SQL("INSERT INTO veterinarias.veterinaria (id, nombre, direccion, telefono) VALUES (4, 'Clínica Veterinaria El Arca de Noé', 'Calle 321 #54-76', '555-3456');")
	m.SQL("INSERT INTO veterinarias.veterinaria (id, nombre, direccion, telefono) VALUES (5, 'Veterinaria Amigos Peludos', 'Avenida 654 #87-09', '555-7890');")

}

// Reverse the migrations
func (m *InsertTableVeterinaria_20260521_171617) Down() {
	m.SQL("DELETE FROM veterinarias.veterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
