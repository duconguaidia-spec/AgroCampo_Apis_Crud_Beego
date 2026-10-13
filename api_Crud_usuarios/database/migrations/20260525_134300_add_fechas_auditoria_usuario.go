package main

import "github.com/beego/beego/v2/client/orm/migration"

// DO NOT MODIFY
type AddFechasAuditoriaUsuario_20260525_134300 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddFechasAuditoriaUsuario_20260525_134300{}
	m.Created = "20260525_134300"

	migration.Register("AddFechasAuditoriaUsuario_20260525_134300", m)
}

// Run the migrations
func (m *AddFechasAuditoriaUsuario_20260525_134300) Up() {
	m.SQL(`ALTER TABLE usuarios."AuditoriaUsuario"
		ADD COLUMN IF NOT EXISTS fecha_creacion timestamp without time zone NOT NULL DEFAULT now();`)
	m.SQL(`ALTER TABLE usuarios."AuditoriaUsuario"
		ADD COLUMN IF NOT EXISTS fecha_modificacion timestamp without time zone NOT NULL DEFAULT now();`)
}

// Reverse the migrations
func (m *AddFechasAuditoriaUsuario_20260525_134300) Down() {
	m.SQL(`ALTER TABLE usuarios."AuditoriaUsuario" DROP COLUMN IF EXISTS fecha_modificacion;`)
	m.SQL(`ALTER TABLE usuarios."AuditoriaUsuario" DROP COLUMN IF EXISTS fecha_creacion;`)
}
