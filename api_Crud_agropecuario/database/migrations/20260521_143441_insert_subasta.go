package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertSubasta_20260521_143441 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertSubasta_20260521_143441{}
	m.Created = "20260521_143441"

	migration.Register("InsertSubasta_20260521_143441", m)
}

// Run the migrations
func (m *InsertSubasta_20260521_143441) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO agropecuario.subasta (id_subasta, fecha_subasta, ubicacion, precio, id_usuario_registra) VALUES (1, '2025-03-10', 'Feria ganadera de Villavicencio, Meta', 2500000.00, 1) ON CONFLICT (id_subasta) DO NOTHING;")
	m.SQL("SELECT setval('agropecuario.\"Subasta_id_subasta_seq\"', GREATEST((SELECT COALESCE(MAX(id_subasta), 0) FROM agropecuario.subasta), 1), true);")

}

// Reverse the migrations
func (m *InsertSubasta_20260521_143441) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM agropecuario.subasta WHERE id_subasta = 1;")

}
