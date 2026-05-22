package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTrVeterinariasespecialidad_20260522_130557 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTrVeterinariasespecialidad_20260522_130557{}
	m.Created = "20260522_130557"

	migration.Register("InsertTableTrVeterinariasespecialidad_20260522_130557", m)
}

// Run the migrations
func (m *InsertTableTrVeterinariasespecialidad_20260522_130557) Up() {
	m.SQL("INSERT INTO veterinarias.tr_veterinariasespecialidad (id, id_veterinaria, id_especialidad) VALUES (1, 1, 1);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariasespecialidad (id, id_veterinaria, id_especialidad) VALUES (2, 1, 2);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariasespecialidad (id, id_veterinaria, id_especialidad) VALUES (3, 2, 1);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariasespecialidad (id, id_veterinaria, id_especialidad) VALUES (4, 3, 3);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariasespecialidad (id, id_veterinaria, id_especialidad) VALUES (5, 4, 4);")

}

// Reverse the migrations
func (m *InsertTableTrVeterinariasespecialidad_20260522_130557) Down() {
	m.SQL("DELETE FROM veterinarias.tr_veterinariasespecialidad WHERE id IN (1, 2, 3, 4, 5);")

}
