package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableNoticia_20260521_133356 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableNoticia_20260521_133356{}
	m.Created = "20260521_133356"

	migration.Register("InsertTableNoticia_20260521_133356", m)
}

// Run the migrations
func (m *InsertTableNoticia_20260521_133356) Up() {
	m.SQL("INSERT INTO Noticias.Noticias (titulo, tipo, cuerpo, url_video, imagen_destacada, fuente, fecha_noticia, id_usuario, acceso_limitado, estado, activo) VALUES ('Título de la noticia', 'Tipo de noticia', 'Cuerpo de la noticia', 'https://example.com/video', 'imagen.jpg', 'Fuente de la noticia', '2023-10-01', 1, false, 'Borrador', true)")

}

// Reverse the migrations
func (m *InsertTableNoticia_20260521_133356) Down() {
	m.SQL("DELETE FROM Noticias.Noticias WHERE titulo = 'Título de la noticia' AND tipo = 'Tipo de noticia' AND cuerpo = 'Cuerpo de la noticia' AND url_video = 'https://example.com/video' AND imagen_destacada = 'imagen.jpg' AND fuente = 'Fuente de la noticia' AND fecha_noticia = '2023-10-01' AND id_usuario = 1 AND acceso_limitado = false AND estado = 'Borrador' AND activo = true")

}
