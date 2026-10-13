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
	m.SQL("INSERT INTO veterinarias.tr_VeterinariaEspecialidad (id_veterinaria, id_especialidad, activo) VALUES (1, 1, TRUE);")
	m.SQL("INSERT INTO veterinarias.tr_VeterinariaEspecialidad (id_veterinaria, id_especialidad, activo) VALUES (1, 2, TRUE);")
	m.SQL("INSERT INTO veterinarias.tr_VeterinariaEspecialidad (id_veterinaria, id_especialidad, activo) VALUES (2, 1, TRUE);")
	m.SQL("INSERT INTO veterinarias.tr_VeterinariaEspecialidad (id_veterinaria, id_especialidad, activo) VALUES (3, 3, TRUE);")
	m.SQL("INSERT INTO veterinarias.tr_VeterinariaEspecialidad (id_veterinaria, id_especialidad, activo) VALUES (4, 4, TRUE);")

}

// Reverse the migrations
func (m *InsertTableTrVeterinariasespecialidad_20260522_130557) Down() {
	m.SQL("DELETE FROM veterinarias.tr_VeterinariaEspecialidad WHERE id IN (1, 2, 3, 4, 5);")

}
