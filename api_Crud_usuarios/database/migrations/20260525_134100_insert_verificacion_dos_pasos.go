package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertVerificacionDosPasos_20260525_134100 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertVerificacionDosPasos_20260525_134100{}
	m.Created = "20260525_134100"

	migration.Register("InsertVerificacionDosPasos_20260525_134100", m)
}

// Run the migrations
func (m *InsertVerificacionDosPasos_20260525_134100) Up() {
	m.SQL(`INSERT INTO usuarios."VerificacionDosPasos" (id_verificacion, id_usuario, tipo_metodo, codigo_otp, expiracion_otp, activo)
		VALUES (1, 1, 'correo', '123456', '2026-12-31 23:59:59', true)
		ON CONFLICT (id_verificacion) DO UPDATE SET
			id_usuario = EXCLUDED.id_usuario,
			tipo_metodo = EXCLUDED.tipo_metodo,
			codigo_otp = EXCLUDED.codigo_otp,
			expiracion_otp = EXCLUDED.expiracion_otp,
			activo = EXCLUDED.activo;`)
	m.SQL(`SELECT setval('usuarios."VerificacionDosPasos_id_verificacion_seq"', GREATEST((SELECT COALESCE(MAX(id_verificacion), 0) FROM usuarios."VerificacionDosPasos"), 1), true);`)
}

// Reverse the migrations
func (m *InsertVerificacionDosPasos_20260525_134100) Down() {
	m.SQL(`DELETE FROM usuarios."VerificacionDosPasos" WHERE id_verificacion = 1;`)
}
