package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type DropFkSubastaUsuario_20260525_120000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DropFkSubastaUsuario_20260525_120000{}
	m.Created = "20260525_120000"

	migration.Register("DropFkSubastaUsuario_20260525_120000", m)
}

// Run the migrations
func (m *DropFkSubastaUsuario_20260525_120000) Up() {
	m.SQL(`ALTER TABLE IF EXISTS agropecuario.subasta DROP CONSTRAINT IF EXISTS fk_subasta_usuario;`)
}

// Reverse the migrations
func (m *DropFkSubastaUsuario_20260525_120000) Down() {
	m.SQL(`ALTER TABLE IF EXISTS agropecuario.subasta
		ADD CONSTRAINT fk_subasta_usuario FOREIGN KEY (id_usuario_registra)
		REFERENCES usuarios."Usuario" (id_usuario)
		ON UPDATE CASCADE
		ON DELETE RESTRICT;`)
}
