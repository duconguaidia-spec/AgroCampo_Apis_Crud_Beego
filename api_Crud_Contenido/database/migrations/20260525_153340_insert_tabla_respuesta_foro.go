package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaRespuestaForo_20260525_153340 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaRespuestaForo_20260525_153340{}
	m.Created = "20260525_153340"

	migration.Register("InsertTablaRespuestaForo_20260525_153340", m)
}

// Run the migrations
func (m *InsertTablaRespuestaForo_20260525_153340) Up() {
	m.SQL("CREATE TABLE RespuestasForo (id INT AUTO_INCREMENT PRIMARY KEY, id_tema INT, id_autor INT, descripcion TEXT NOT NULL, FOREIGN KEY (id_tema) REFERENCES TemasForo(id), FOREIGN KEY (id_autor) REFERENCES Usuarios(id))")
	m.SQL("INSERT INTO RespuestasForo (id_tema, id_autor, descripcion) VALUES")
	m.SQL("    (1, 2, 'Excelente aporte, gracias por compartir'),")
	m.SQL("    (1, 3, '¿Podrías compartir fuentes?'),")
	m.SQL("    (2, 1, 'Muy útil, aplicaré esto');")
}

// Reverse the migrations
func (m *InsertTablaRespuestaForo_20260525_153340) Down() {
	m.SQL("DROP TABLE RespuestasForo") 

}
