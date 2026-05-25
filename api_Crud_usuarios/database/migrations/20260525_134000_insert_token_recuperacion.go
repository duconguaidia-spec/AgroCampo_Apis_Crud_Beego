package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertTokenRecuperacion_20260525_134000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTokenRecuperacion_20260525_134000{}
	m.Created = "20260525_134000"

	migration.Register("InsertTokenRecuperacion_20260525_134000", m)
}

// Run the migrations
func (m *InsertTokenRecuperacion_20260525_134000) Up() {
	m.SQL(`INSERT INTO usuarios."TokenRecuperacion" (id_token, token, id_usuario, expiracion, usado)
		VALUES (1, 'token_demo_admin', 1, '2026-12-31 23:59:59', false)
		ON CONFLICT (id_token) DO UPDATE SET
			token = EXCLUDED.token,
			id_usuario = EXCLUDED.id_usuario,
			expiracion = EXCLUDED.expiracion,
			usado = EXCLUDED.usado;`)
	m.SQL(`SELECT setval('usuarios."TokenRecuperacion_id_token_seq"', GREATEST((SELECT COALESCE(MAX(id_token), 0) FROM usuarios."TokenRecuperacion"), 1), true);`)
}

// Reverse the migrations
func (m *InsertTokenRecuperacion_20260525_134000) Down() {
	m.SQL(`DELETE FROM usuarios."TokenRecuperacion" WHERE id_token = 1;`)
}
