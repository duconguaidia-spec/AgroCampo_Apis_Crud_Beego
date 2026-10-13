package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaCategoria_20260525_153157 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaCategoria_20260525_153157{}
	m.Created = "20260525_153157"

	migration.Register("InsertTablaCategoria_20260525_153157", m)
}

// Run the migrations
func (m *InsertTablaCategoria_20260525_153157) Up() {
	m.SQL("INSERT INTO public.categoria (nombre_categoria, activo) VALUES")
	m.SQL("    ('Agricultura', true),")
	m.SQL("    ('Ganadería', true),")
	m.SQL("    ('Tecnología', true);")
	

}

// Reverse the migrations
func (m *InsertTablaCategoria_20260525_153157) Down() {
	m.SQL("DELETE FROM public.categoria WHERE nombre_categoria IN ('Agricultura', 'Ganadería', 'Tecnología')")


}
