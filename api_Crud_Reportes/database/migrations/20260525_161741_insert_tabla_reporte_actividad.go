package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaReporteActividad_20260525_161741 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaReporteActividad_20260525_161741{}
	m.Created = "20260525_161741"

	migration.Register("InsertTablaReporteActividad_20260525_161741", m)
}

// Run the migrations
func (m *InsertTablaReporteActividad_20260525_161741) Up() {
	m.SQL("INSERT INTO public.reporte_actividad (id_usuario, id_actividad, descripcion_reporte, fecha_reporte) VALUES")
	m.SQL("    (1, 1, 'Reporte de actividad 1', '2026-05-25 16:17:41'),")
	m.SQL("    (2, 2, 'Reporte de actividad 2', '2026-05-25 16:17:41'),")
	m.SQL("    (3, 3, 'Reporte de actividad 3', '2026-05-25 16:17:41');")
}

// Reverse the migrations
func (m *InsertTablaReporteActividad_20260525_161741) Down() {
	m.SQL("DELETE FROM public.reporte_actividad WHERE descripcion_reporte IN ('Reporte de actividad 1', 'Reporte de actividad 2', 'Reporte de actividad 3')")

}
