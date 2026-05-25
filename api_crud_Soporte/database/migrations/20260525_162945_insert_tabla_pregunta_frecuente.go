package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type InsertTablaPreguntaFrecuente_20260525_162945 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaPreguntaFrecuente_20260525_162945{}
	m.Created = "20260525_162945"

	migration.Register("InsertTablaPreguntaFrecuente_20260525_162945", m)
}

// Run the migrations
func (m *InsertTablaPreguntaFrecuente_20260525_162945) Up() {
	m.SQL("INSERT INTO PreguntaFrecuente (pregunta, respuesta, id_categoria_faq) VALUES")
	m.SQL("    ('¿Cómo puedo restablecer mi contraseña?', 'Para restablecer tu contraseña, haz clic en \"Olvidé mi contraseña\" en la página de inicio de sesión y sigue las instrucciones.', 1),")
	m.SQL("    ('¿Cómo puedo contactar al soporte técnico?', 'Puedes contactar al soporte técnico enviando un correo a soporte@empresa.com', 2);")
}

// Reverse the migrations
func (m *InsertTablaPreguntaFrecuente_20260525_162945) Down() {
	m.SQL("DELETE FROM PreguntaFrecuente WHERE pregunta IN ('¿Cómo puedo restablecer mi contraseña?', '¿Cómo puedo contactar al soporte técnico?')")
}
