package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTableVeterinaria_20260521_171617 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTableVeterinaria_20260521_171617{}
	m.Created = "20260521_171617"

	migration.Register("InsertTableTrVeterinaria_20260521_171617", m)
}

// Run the migrations
func (m *InsertTableVeterinaria_20260521_171617) Up() {
	m.SQL("INSERT INTO veterinarias.veterinaria (nombre_clinica, direccion, telefono, correo_publico, horario_atencion, calificacion_promedio) VALUES ('Veterinaria San Roque', 'Calle 123 #45-67', '555-1234', 'info@veterinariasanroque.com', '08:00 - 20:00', 4.5);")
	m.SQL("INSERT INTO veterinarias.veterinaria (nombre_clinica, direccion, telefono, correo_publico, horario_atencion, calificacion_promedio) VALUES ('Clínica Veterinaria El Bosque', 'Avenida 456 #78-90', '555-5678', 'contacto@clinicaelbosque.com', '09:00 - 18:00', 4.2);")
	m.SQL("INSERT INTO veterinarias.veterinaria (nombre_clinica, direccion, telefono, correo_publico, horario_atencion, calificacion_promedio) VALUES ('Veterinaria La Mascota Feliz', 'Carrera 789 #12-34', '555-9012', 'hola@veterinarialamascotafeliz.com', '10:00 - 21:00', 4.8);")
	m.SQL("INSERT INTO veterinarias.veterinaria (nombre_clinica, direccion, telefono, correo_publico, horario_atencion, calificacion_promedio) VALUES ('Clínica Veterinaria El Arca de Noé', 'Calle 321 #54-76', '555-3456', 'servicio@clinicaelarcaconoé.com', '08:30 - 20:30', 4.6);")
	m.SQL("INSERT INTO veterinarias.veterinaria (nombre_clinica, direccion, telefono, correo_publico, horario_atencion, calificacion_promedio) VALUES ('Veterinaria Amigos Peludos', 'Avenida 654 #87-09', '555-7890', 'atencion@veterinariaamigospeludos.com', '09:30 - 19:30', 4.4);")

}

// Reverse the migrations
func (m *InsertTableVeterinaria_20260521_171617) Down() {
	m.SQL("DELETE FROM veterinarias.veterinaria WHERE id IN (1, 2, 3, 4, 5);")

}
