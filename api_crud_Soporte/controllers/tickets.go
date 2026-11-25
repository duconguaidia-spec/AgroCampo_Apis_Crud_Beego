package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type TicketsController struct{ beego.Controller }

func (c *TicketsController) Post() {
	var datos struct {
		Usuario     int `json:"usuario_id"`
		Asunto      string
		Descripcion string
	}
	if json.Unmarshal(c.Ctx.Input.RequestBody, &datos) != nil || datos.Usuario <= 0 || strings.TrimSpace(datos.Descripcion) == "" || len(datos.Descripcion) > 10000 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]any{"success": false, "message": "Describe la incidencia (máximo 10 000 caracteres)"}
		c.ServeJSON()
		return
	}
	db, err := orm.GetDB()
	var id int
	if err == nil {
		err = db.QueryRowContext(c.Ctx.Request.Context(), `INSERT INTO soporte.tickets_soporte(usuario_id,asunto,descripcion) VALUES($1,$2,$3) RETURNING id`, datos.Usuario, datos.Asunto, datos.Descripcion).Scan(&id)
	}
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]any{"success": false}
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = map[string]any{"success": true, "data": map[string]int{"id": id}}
	}
	c.ServeJSON()
}
