package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableResenaveterinaria_20260522_130500 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableResenaveterinaria_20260522_130500{}
	m.Created = "20260522_130500"

	migration.Register("InsertTableResenaveterinaria_20260522_130500", m)
}

// Run the migrations
func (m *InsertTableResenaveterinaria_20260522_130500) Up() {
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id, id_veterinaria, calificacion, comentario) VALUES (1, 1, 5, 'Excelente servicio y atención.');")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id, id_veterinaria, calificacion, comentario) VALUES (2, 1, 4, 'Muy buena experiencia, aunque el tiempo de espera fue un poco largo.');")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id, id_veterinaria, calificacion, comentario) VALUES (3, 2, 3, 'El servicio fue aceptable, pero podrían mejorar la comunicación con los clientes.');")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id, id_veterinaria, calificacion, comentario) VALUES (4, 3, 5, 'Mi mascota recibió un excelente cuidado y atención.');")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id, id_veterinaria, calificacion, comentario) VALUES (5, 4, 2, 'No quedé satisfecho con el servicio, el personal no fue muy amable.');")

}

// Reverse the migrations
func (m *InsertTableResenaveterinaria_20260522_130500) Down() {
	m.SQL("DELETE FROM veterinarias.resenaveterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
