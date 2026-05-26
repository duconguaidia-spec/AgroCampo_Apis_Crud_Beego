package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaTemaForo_20260525_153313 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaTemaForo_20260525_153313{}
	m.Created = "20260525_153313"

	migration.Register("InsertTablaTemaForo_20260525_153313", m)
}

// Run the migrations
func (m *InsertTablaTemaForo_20260525_153313) Up() {
	m.SQL("CREATE TABLE TemasForo (id INT AUTO_INCREMENT PRIMARY KEY, titulo VARCHAR(255) NOT NULL, descripcion TEXT, id_autor INT, id_categoria INT, estado VARCHAR(50), FOREIGN KEY (id_autor) REFERENCES Usuarios(id), FOREIGN KEY (id_categoria) REFERENCES Categorias(id))")
	m.SQL("INSERT INTO TemasForo (titulo, descripcion, id_autor, id_categoria, estado) VALUES")
	m.SQL("    ('Mejoras en riego', 'Discusión sobre técnicas de riego eficientes', 1, 1, 'abierto'),")
	m.SQL("    ('Cuidado del ganado', 'Consejos para salud animal', 2, 2, 'abierto'),")
	m.SQL("    ('Herramientas digitales', 'Apps útiles para campo', 3, 3, 'abierto');")
}

// Reverse the migrations
func (m *InsertTablaTemaForo_20260525_153313) Down() {
	m.SQL("DROP TABLE TemasForo")
}