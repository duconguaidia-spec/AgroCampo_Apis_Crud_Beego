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
	m.SQL("INSERT INTO tr_precio_subasta_ganado (id_subasta, id_ganado, precio) VALUES (1, 1, 6800);")

}

// Reverse the migrations
func (m *InsertTrPrecioSubastaGanado_20260521_143520) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM tr_precio_subasta_ganado WHERE id_subasta = 1 AND id_ganado = 1;")
	

}
