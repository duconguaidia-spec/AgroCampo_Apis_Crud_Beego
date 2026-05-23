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
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id_veterinaria, nombre_producto, descripcion, precio, activo) VALUES (1, 'Alimento para Perros', 'Alimento balanceado para perros de todas las edades.', 29.99, true);")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id_veterinaria, nombre_producto, descripcion, precio, activo) VALUES (1, 'Alimento para Gatos', 'Alimento balanceado para gatos de todas las edades.', 24.99, true);")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id_veterinaria, nombre_producto, descripcion, precio, activo) VALUES (1, 'Shampoo Antipulgas', 'Shampoo especialmente formulado para eliminar pulgas y garrapatas.', 19.99, true);")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id_veterinaria, nombre_producto, descripcion, precio, activo) VALUES (1, 'Juguete para Mascotas', 'Juguete interactivo para mantener a tu mascota entretenida.', 14.99, true);")
	m.SQL("INSERT INTO veterinarias.productoveterinaria (id_veterinaria, nombre_producto, descripcion, precio, activo) VALUES (1, 'Cama para Mascotas', 'Cama cómoda y acogedora para el descanso de tu mascota.', 39.99, true);")

}

// Reverse the migrations
func (m *InsertTableProductoveterinaria_20260522_130431) Down() {
	m.SQL("DELETE FROM veterinarias.productoveterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
