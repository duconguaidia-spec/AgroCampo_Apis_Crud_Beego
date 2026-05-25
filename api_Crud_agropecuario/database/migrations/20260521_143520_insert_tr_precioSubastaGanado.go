package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTrPrecioSubastaGanado_20260521_143520 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTrPrecioSubastaGanado_20260521_143520{}
	m.Created = "20260521_143520"

	migration.Register("InsertTrPrecioSubastaGanado_20260521_143520", m)
}

// Run the migrations
func (m *InsertTrPrecioSubastaGanado_20260521_143520) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO agropecuario.\"tr_precioSubastaGanado\" (id_subasta, id_categoria_ganado) VALUES (1, 1) ON CONFLICT (id_subasta, id_categoria_ganado) DO NOTHING;")
	m.SQL("SELECT setval('agropecuario.\"Tr_PrecioSubastaGanado_id_precio_subasta_seq\"', GREATEST((SELECT COALESCE(MAX(id_precio_subasta), 0) FROM agropecuario.\"tr_precioSubastaGanado\"), 1), true);")

}

// Reverse the migrations
func (m *InsertTrPrecioSubastaGanado_20260521_143520) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM agropecuario.\"tr_precioSubastaGanado\" WHERE id_subasta = 1 AND id_categoria_ganado = 1;")

}
