package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableTrVeterinariaservicio_20260522_130609 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableTrVeterinariaservicio_20260522_130609{}
	m.Created = "20260522_130609"

	migration.Register("InsertTableTrVeterinariaservicio_20260522_130609", m)
}

// Run the migrations
func (m *InsertTableTrVeterinariaservicio_20260522_130609) Up() {
	m.SQL("INSERT INTO veterinarias.tr_veterinariaservicio (id, id_veterinaria, id_serviciogeneral) VALUES (1, 1, 1);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariaservicio (id, id_veterinaria, id_serviciogeneral) VALUES (2, 1, 2);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariaservicio (id, id_veterinaria, id_serviciogeneral) VALUES (3, 2, 1);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariaservicio (id, id_veterinaria, id_serviciogeneral) VALUES (4, 3, 3);")
	m.SQL("INSERT INTO veterinarias.tr_veterinariaservicio (id, id_veterinaria, id_serviciogeneral) VALUES (5, 4, 4);")

}

// Reverse the migrations
func (m *InsertTableTrVeterinariaservicio_20260522_130609) Down() {
	m.SQL("DELETE FROM veterinarias.tr_veterinariaservicio WHERE id IN (1, 2, 3, 4, 5);")

}
