package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaVideoEducativo_20260525_153409 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaVideoEducativo_20260525_153409{}
	m.Created = "20260525_153409"

	migration.Register("InsertTablaVideoEducativo_20260525_153409", m)
}

// Run the migrations
func (m *InsertTablaVideoEducativo_20260525_153409) Up() {
	m.SQL("CREATE TABLE VideosEducativos (id INT AUTO_INCREMENT PRIMARY KEY, titulo VARCHAR(255) NOT NULL, descripcion TEXT, url_video VARCHAR(255) NOT NULL, id_usuario INT, FOREIGN KEY (id_usuario) REFERENCES Usuarios(id))")
	m.SQL("INSERT INTO VideosEducativos (titulo, descripcion, url_video, id_usuario) VALUES")
	m.SQL("    ('Uso de drones', 'Introducción al uso de drones en agricultura', 'https://example.com/video1', 1),")
	m.SQL("    ('Manejo de suelos', 'Técnicas para mejorar la fertilidad', 'https://example.com/video2', 2);")

}

// Reverse the migrations
func (m *InsertTablaVideoEducativo_20260525_153409) Down() {
	m.SQL("DROP TABLE VideosEducativos") 

}
