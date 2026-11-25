package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type SectorController struct{ beego.Controller }

func (c *SectorController) Cotizaciones() {
	db, err := orm.GetDB()
	var datos []byte
	if err == nil {
		err = db.QueryRowContext(c.Ctx.Request.Context(), `
			WITH ultimas AS (SELECT id,fecha_subasta FROM agropecuario.subastas WHERE activo ORDER BY fecha_subasta DESC,id DESC LIMIT 4)
			SELECT coalesce(json_agg(json_build_object('fecha',s.fecha_subasta,'categoria',cat.nombre,'precio',p.precio_por_kg)
				ORDER BY s.fecha_subasta DESC,s.id DESC,cat.nombre),'[]'::json)
			FROM ultimas s JOIN agropecuario.precios_subasta_ganado p ON p.subasta_id=s.id
			JOIN agropecuario.categorias_ganado cat ON cat.id=p.categoria_ganado_id WHERE cat.activo AND p.activo`).Scan(&datos)
	}
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]any{"success": false}
	} else {
		c.Data["json"] = map[string]any{"success": true, "data": json.RawMessage(datos)}
	}
	c.ServeJSON()
}

func (c *SectorController) Get() {
	db, err := orm.GetDB()
	var datos []byte
	if err == nil {
		err = db.QueryRowContext(c.Ctx.Request.Context(), `SELECT coalesce(json_agg(json_build_object('id',m.id,'indicador',i.nombre,'unidad',i.unidad,'fecha',m.fecha,'region',m.region,'valor',m.valor,'fuente',coalesce(m.fuente,i.fuente,'')) ORDER BY m.fecha DESC,i.nombre),'[]'::json) FROM agropecuario.mediciones_sector m JOIN agropecuario.indicadores_sector i ON i.id=m.indicador_id WHERE m.activo AND i.activo`).Scan(&datos)
	}
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]any{"success": false}
	} else {
		c.Data["json"] = map[string]any{"success": true, "data": json.RawMessage(datos)}
	}
	c.ServeJSON()
}

type PrecioSubastaController struct{ beego.Controller }

func (c *PrecioSubastaController) Post() {
	var datos struct {
		Subasta   int     `json:"subasta_id"`
		Categoria int     `json:"categoria_ganado_id"`
		Precio    float64 `json:"precio_por_kg"`
	}
	if json.Unmarshal(c.Ctx.Input.RequestBody, &datos) != nil || datos.Subasta <= 0 || datos.Categoria <= 0 || datos.Precio <= 0 || datos.Precio >= 1e10 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]any{"success": false, "message": "Revisa la subasta, categoría y precio"}
		c.ServeJSON()
		return
	}
	db, err := orm.GetDB()
	var id int
	if err == nil {
		err = db.QueryRowContext(c.Ctx.Request.Context(), `INSERT INTO agropecuario.precios_subasta_ganado(subasta_id,categoria_ganado_id,precio_por_kg) VALUES($1,$2,$3) ON CONFLICT(subasta_id,categoria_ganado_id) DO UPDATE SET precio_por_kg=excluded.precio_por_kg,activo=true RETURNING id`, datos.Subasta, datos.Categoria, datos.Precio).Scan(&id)
	}
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]any{"success": false}
	} else {
		c.Data["json"] = map[string]any{"success": true, "data": map[string]int{"id": id}}
	}
	c.ServeJSON()
}
