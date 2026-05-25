package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaCategoriaPregunta_20260525_162906 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaCategoriaPregunta_20260525_162906{}
	m.Created = "20260525_162906"

	migration.Register("InsertTablaCategoriaPregunta_20260525_162906", m)
}

// Run the migrations
func (m *InsertTablaCategoriaPregunta_20260525_162906) Up() {

	m.SQL("INSERT INTO CategoriaPreguntas (nombre_categoria, activo) VALUES")
	m.SQL("    ('General', true),")
	m.SQL("    ('Técnica', true),")
	m.SQL("    ('Administrativa', true);")


}

// Reverse the migrations
func (m *InsertTablaCategoriaPregunta_20260525_162906) Down() {
	m.SQL("DELETE FROM CategoriaPreguntas WHERE nombre_categoria IN ('General', 'Técnica', 'Administrativa')")

}
