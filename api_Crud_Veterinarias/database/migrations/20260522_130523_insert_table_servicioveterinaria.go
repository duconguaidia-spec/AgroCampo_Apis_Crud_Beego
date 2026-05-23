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
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id, id_veterinaria, id_serviciogeneral) VALUES (1, 1, 1);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id, id_veterinaria, id_serviciogeneral) VALUES (2, 1, 2);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id, id_veterinaria, id_serviciogeneral) VALUES (3, 2, 1);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id, id_veterinaria, id_serviciogeneral) VALUES (4, 3, 3);")
	m.SQL("INSERT INTO veterinarias.servicioveterinaria (id, id_veterinaria, id_serviciogeneral) VALUES (5, 4, 4);")

}

// Reverse the migrations
func (m *InsertTableServicioveterinaria_20260522_130523) Down() {
	m.SQL("DELETE FROM veterinarias.servicioveterinaria WHERE id IN (1, 2, 3, 4, 5);")

}