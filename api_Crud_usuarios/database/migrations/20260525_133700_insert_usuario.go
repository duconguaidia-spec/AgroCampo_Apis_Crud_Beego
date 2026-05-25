package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertUsuario_20260525_133700 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertUsuario_20260525_133700{}
	m.Created = "20260525_133700"

	migration.Register("InsertUsuario_20260525_133700", m)
}

// Run the migrations
func (m *InsertUsuario_20260525_133700) Up() {
	m.SQL(`INSERT INTO usuarios."Usuario" (id_usuario, nombre_completo, correo, telefono, id_rol, verificacion_dos_pasos, avatar, activo)
		VALUES (1, 'Usuario Administrador', 'admin@agrocampo.com', '3000000001', 1, false, 'admin.png', true)
		ON CONFLICT (id_usuario) DO UPDATE SET
			nombre_completo = EXCLUDED.nombre_completo,
			correo = EXCLUDED.correo,
			telefono = EXCLUDED.telefono,
			id_rol = EXCLUDED.id_rol,
			verificacion_dos_pasos = EXCLUDED.verificacion_dos_pasos,
			avatar = EXCLUDED.avatar,
			activo = EXCLUDED.activo;`)
	m.SQL(`SELECT setval('usuarios."Usuario_id_usuario_seq"', GREATEST((SELECT COALESCE(MAX(id_usuario), 0) FROM usuarios."Usuario"), 1), true);`)
}

// Reverse the migrations
func (m *InsertUsuario_20260525_133700) Down() {
	m.SQL(`DELETE FROM usuarios."Usuario" WHERE id_usuario = 1;`)
}
