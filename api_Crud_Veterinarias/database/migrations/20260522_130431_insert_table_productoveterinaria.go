package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableProductoveterinaria_20260522_130431 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableProductoveterinaria_20260522_130431{}
	m.Created = "20260522_130431"

	migration.Register("InsertTableProductoveterinaria_20260522_130431", m)
}

// Run the migrations
func (m *InsertTableProductoveterinaria_20260522_130431) Up() {
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id, nombre, descripcion) VALUES (1, 'Alimento para Perros', 'Alimento balanceado para perros de todas las edades.');")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id, nombre, descripcion) VALUES (2, 'Alimento para Gatos', 'Alimento balanceado para gatos de todas las edades.');")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id, nombre, descripcion) VALUES (3, 'Shampoo Antipulgas', 'Shampoo especialmente formulado para eliminar pulgas y garrapatas.');")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id, nombre, descripcion) VALUES (4, 'Juguete para Mascotas', 'Juguete interactivo para mantener a tu mascota entretenida.');")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id, nombre, descripcion) VALUES (5, 'Cama para Mascotas', 'Cama cómoda y acogedora para el descanso de tu mascota.');")

}

// Reverse the migrations
func (m *InsertTableProductoveterinaria_20260522_130431) Down() {
	m.SQL("DELETE FROM veterinarias.productoveterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
