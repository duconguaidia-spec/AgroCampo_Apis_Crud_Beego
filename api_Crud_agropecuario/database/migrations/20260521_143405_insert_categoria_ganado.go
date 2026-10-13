package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertCategoriaGanado_20260521_143405 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertCategoriaGanado_20260521_143405{}
	m.Created = "20260521_143405"

	migration.Register("InsertCategoriaGanado_20260521_143405", m)
}

// Run the migrations
func (m *InsertCategoriaGanado_20260521_143405) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO agropecuario.categoria_ganado (codigo, descripcion) VALUES ('VG', 'Vaca Gorda') ON CONFLICT (codigo) DO UPDATE SET descripcion = EXCLUDED.descripcion;")

}

// Reverse the migrations
func (m *InsertCategoriaGanado_20260521_143405) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM agropecuario.categoria_ganado WHERE codigo = 'VG';")

}
