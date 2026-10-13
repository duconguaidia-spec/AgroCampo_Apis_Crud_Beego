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
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id_veterinaria, id_usuario, calificacion, comentario, activo) VALUES (1, 1, 5, 'Excelente servicio y atención.', TRUE);")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id_veterinaria, id_usuario, calificacion, comentario, activo) VALUES (1, 2, 4, 'Muy buena experiencia, aunque el tiempo de espera fue un poco largo.', TRUE);")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id_veterinaria, id_usuario, calificacion, comentario, activo) VALUES (2, 3, 3, 'El servicio fue aceptable, pero podrían mejorar la comunicación con los clientes.', TRUE);")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id_veterinaria, id_usuario, calificacion, comentario, activo) VALUES (3, 4, 5, 'Mi mascota recibió un excelente cuidado y atención.', TRUE);")
	m.SQL("INSERT INTO veterinarias.resenaveterinaria (id_veterinaria, id_usuario, calificacion, comentario, activo) VALUES (4, 5, 2, 'No quedé satisfecho con el servicio, el personal no fue muy amable.', TRUE);")

}

// Reverse the migrations
func (m *InsertTableResenaveterinaria_20260522_130500) Down() {
	m.SQL("DELETE FROM veterinarias.resenaveterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
