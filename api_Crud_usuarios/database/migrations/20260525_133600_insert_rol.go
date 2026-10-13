package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertRol_20260525_133600 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertRol_20260525_133600{}
	m.Created = "20260525_133600"

	migration.Register("InsertRol_20260525_133600", m)
}

// Run the migrations
func (m *InsertRol_20260525_133600) Up() {
	m.SQL(`INSERT INTO usuarios."Rol" (id_rol, nombre_rol, descripcion, activo) VALUES
		(1, 'Administrador', 'Control total del sistema', true),
		(2, 'Ganadero', 'Usuario propietario de ganado y fincas', true),
		(3, 'Veterinario', 'Profesional de salud animal registrado', true)
		ON CONFLICT (id_rol) DO UPDATE SET
			nombre_rol = EXCLUDED.nombre_rol,
			descripcion = EXCLUDED.descripcion,
			activo = EXCLUDED.activo;`)
	m.SQL(`SELECT setval('usuarios."Rol_id_rol_seq"', GREATEST((SELECT COALESCE(MAX(id_rol), 0) FROM usuarios."Rol"), 1), true);`)
}

// Reverse the migrations
func (m *InsertRol_20260525_133600) Down() {
	m.SQL(`DELETE FROM usuarios."Rol" WHERE id_rol IN (1, 2, 3);`)
}
