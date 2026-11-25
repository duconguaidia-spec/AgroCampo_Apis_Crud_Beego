package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type GuardadosController struct{ beego.Controller }

func (c *GuardadosController) Get() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":usuario"))
	if err != nil || id <= 0 {
		c.responder(400, nil)
		return
	}
	c.listar(id)
}

func (c *GuardadosController) Post() {
	var datos struct {
		Usuario  int   `json:"usuario_id"`
		Tema     int   `json:"tema_id"`
		Guardado *bool `json:"guardado"`
	}
	if json.Unmarshal(c.Ctx.Input.RequestBody, &datos) != nil || datos.Usuario <= 0 || datos.Tema <= 0 || datos.Guardado == nil {
		c.responder(400, nil)
		return
	}
	db, err := orm.GetDB()
	if err == nil {
		if *datos.Guardado {
			_, err = db.ExecContext(c.Ctx.Request.Context(), `INSERT INTO contenido.temas_guardados(usuario_id,tema_id) VALUES($1,$2) ON CONFLICT(usuario_id,tema_id) DO UPDATE SET activo=true`, datos.Usuario, datos.Tema)
		} else {
			_, err = db.ExecContext(c.Ctx.Request.Context(), `DELETE FROM contenido.temas_guardados WHERE usuario_id=$1 AND tema_id=$2`, datos.Usuario, datos.Tema)
		}
	}
	if err != nil {
		c.responder(500, nil)
		return
	}
	c.listar(datos.Usuario)
}

func (c *GuardadosController) listar(id int) {
	db, err := orm.GetDB()
	var datos []byte
	if err == nil {
		err = db.QueryRowContext(c.Ctx.Request.Context(), `SELECT coalesce(json_agg(g.tema_id ORDER BY g.tema_id),'[]'::json) FROM contenido.temas_guardados g JOIN contenido.temas_foro t ON t.id=g.tema_id WHERE g.usuario_id=$1 AND g.activo AND t.activo`, id).Scan(&datos)
	}
	if err != nil {
		c.responder(500, nil)
		return
	}
	c.responder(200, json.RawMessage(datos))
}

func (c *GuardadosController) responder(estado int, datos any) {
	c.Ctx.Output.SetStatus(estado)
	c.Data["json"] = map[string]any{"success": estado < 400, "data": datos}
	c.ServeJSON()
}
