package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type InsertAuditoriaUsuario_20260525_134200 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertAuditoriaUsuario_20260525_134200{}
	m.Created = "20260525_134200"

	migration.Register("InsertAuditoriaUsuario_20260525_134200", m)
}

// Run the migrations
func (m *InsertAuditoriaUsuario_20260525_134200) Up() {
	m.SQL(`INSERT INTO usuarios."AuditoriaUsuario" (id_auditoria, id_usuario, tipo_evento, descripcion, ip_origen, fecha_evento, fecha_creacion, fecha_modificacion)
		VALUES (1, 1, 'CREACION', 'Registro inicial de auditoria para usuario administrador', '127.0.0.1', now(), now(), now())
		ON CONFLICT (id_auditoria) DO UPDATE SET
			id_usuario = EXCLUDED.id_usuario,
			tipo_evento = EXCLUDED.tipo_evento,
			descripcion = EXCLUDED.descripcion,
			ip_origen = EXCLUDED.ip_origen,
			fecha_evento = EXCLUDED.fecha_evento,
			fecha_modificacion = EXCLUDED.fecha_modificacion;`)
	m.SQL(`SELECT setval('usuarios."AuditoriaUsuario_id_auditoria_seq"', GREATEST((SELECT COALESCE(MAX(id_auditoria), 0) FROM usuarios."AuditoriaUsuario"), 1), true);`)
}

// Reverse the migrations
func (m *InsertAuditoriaUsuario_20260525_134200) Down() {
	m.SQL(`DELETE FROM usuarios."AuditoriaUsuario" WHERE id_auditoria = 1;`)
}
