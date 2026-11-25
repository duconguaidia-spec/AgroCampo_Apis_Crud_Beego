package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type CatalogoController struct{ beego.Controller }

func (c *CatalogoController) Get() {
	db, err := orm.GetDB()
	var datos []byte
	if err == nil {
		err = db.QueryRowContext(c.Ctx.Request.Context(), `SELECT coalesce(json_agg(json_build_object('id',id,'nombre',nombre) ORDER BY nombre),'[]'::json) FROM noticias.categorias_noticia WHERE activo`).Scan(&datos)
	}
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]any{"success": false}
	} else {
		c.Data["json"] = map[string]any{"success": true, "data": json.RawMessage(datos)}
	}
	c.ServeJSON()
}
