package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertPerfilExtendido_20260525_133900 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertPerfilExtendido_20260525_133900{}
	m.Created = "20260525_133900"

	migration.Register("InsertPerfilExtendido_20260525_133900", m)
}

// Run the migrations
func (m *InsertPerfilExtendido_20260525_133900) Up() {
	m.SQL(`INSERT INTO usuarios."PerfilExtendido" (id_perfil, id_usuario, direccion, ciudad, intereses, redes_sociales)
		VALUES (1, 1, 'Calle 1 # 1-01', 'Bogota', 'Administracion de AgroCampo', '@agrocampo')
		ON CONFLICT (id_perfil) DO UPDATE SET
			id_usuario = EXCLUDED.id_usuario,
			direccion = EXCLUDED.direccion,
			ciudad = EXCLUDED.ciudad,
			intereses = EXCLUDED.intereses,
			redes_sociales = EXCLUDED.redes_sociales;`)
	m.SQL(`SELECT setval('usuarios."PerfilExtendido_id_perfil_seq"', GREATEST((SELECT COALESCE(MAX(id_perfil), 0) FROM usuarios."PerfilExtendido"), 1), true);`)
}

// Reverse the migrations
func (m *InsertPerfilExtendido_20260525_133900) Down() {
	m.SQL(`DELETE FROM usuarios."PerfilExtendido" WHERE id_perfil = 1;`)
}
