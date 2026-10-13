package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertContrasena_20260525_133800 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertContrasena_20260525_133800{}
	m.Created = "20260525_133800"

	migration.Register("InsertContrasena_20260525_133800", m)
}

// Run the migrations
func (m *InsertContrasena_20260525_133800) Up() {
	m.SQL(`INSERT INTO usuarios."Contrasena" (id_contrasena, id_usuario, contrasena_hash, activa)
		VALUES (1, 1, 'hash_admin_demo', true)
		ON CONFLICT (id_contrasena) DO UPDATE SET
			id_usuario = EXCLUDED.id_usuario,
			contrasena_hash = EXCLUDED.contrasena_hash,
			activa = EXCLUDED.activa;`)
	m.SQL(`SELECT setval('usuarios."Contrasena_id_contrasena_seq"', GREATEST((SELECT COALESCE(MAX(id_contrasena), 0) FROM usuarios."Contrasena"), 1), true);`)
}

// Reverse the migrations
func (m *InsertContrasena_20260525_133800) Down() {
	m.SQL(`DELETE FROM usuarios."Contrasena" WHERE id_contrasena = 1;`)
}
