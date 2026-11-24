package controllers

import (
	"API_CRUD_AGROPECUARIO/models"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// IndicadorSectorController operations for IndicadorSector
type IndicadorSectorController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description crear el indicador
// @router / [post]
func (c *IndicadorSectorController) Post() {
	var dato models.IndicadorSector
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dato)
	if err != nil {
		responder(&c.Controller, 400, false, "El json enviado no es valido", err.Error())
		return
	}

	_, err = models.AddIndicadorSector(&dato)
	if err != nil {
		responder(&c.Controller, 400, false, "No se pudo crear el indicador", err.Error())
		return
	}
	responder(&c.Controller, 201, true, "Indicador creado", dato)
}

// GetOne ...
// @Title GetOne
// @Description consultar el indicador por id
// @router /:id [get]
func (c *IndicadorSectorController) GetOne() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		responder(&c.Controller, 400, false, "El id no es valido", nil)
		return
	}

	dato, err := models.GetIndicadorSectorById(id)
	if err != nil {
		responder(&c.Controller, 404, false, "No existe el indicador", err.Error())
		return
	}
	responder(&c.Controller, 200, true, "Indicador consultado", dato)
}

// GetAll ...
// @Title GetAll
// @Description consultar los indicadores
// @router / [get]
func (c *IndicadorSectorController) GetAll() {
	query, fields, sortby, order, offset, limit, err := leerParametrosLista(&c.Controller)
	if err != nil {
		responder(&c.Controller, 400, false, "Los parametros de la consulta no son validos", err.Error())
		return
	}

	lista, err := models.GetAllIndicadorSector(query, fields, sortby, order, offset, limit)
	if err != nil {
		responder(&c.Controller, 400, false, "Error consultando los indicadores", err.Error())
		return
	}

	// si no hay registros se devuelve una lista vacia, no un error
	if lista == nil {
		lista = []interface{}{}
	}
	responder(&c.Controller, 200, true, "Indicadores consultados", lista)
}

// Put ...
// @Title Put
// @Description modificar el indicador
// @router /:id [put]
func (c *IndicadorSectorController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		responder(&c.Controller, 400, false, "El id no es valido", nil)
		return
	}

	dato := models.IndicadorSector{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &dato)
	if err != nil {
		responder(&c.Controller, 400, false, "El json enviado no es valido", err.Error())
		return
	}
	dato.Id = id

	err = models.UpdateIndicadorSectorById(&dato)
	if err == orm.ErrNoRows {
		responder(&c.Controller, 404, false, "No existe el indicador", nil)
		return
	}
	if err != nil {
		responder(&c.Controller, 400, false, "No se pudo modificar el indicador", err.Error())
		return
	}
	responder(&c.Controller, 200, true, "Indicador modificado", dato)
}

// Delete ...
// @Title Delete
// @Description eliminar el indicador (solo lo desactiva)
// @router /:id [delete]
func (c *IndicadorSectorController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		responder(&c.Controller, 400, false, "El id no es valido", nil)
		return
	}

	err = models.DeleteIndicadorSector(id)
	if err == orm.ErrNoRows {
		responder(&c.Controller, 404, false, "No existe el indicador", nil)
		return
	}
	if err != nil {
		responder(&c.Controller, 400, false, "No se pudo eliminar el indicador", err.Error())
		return
	}
	responder(&c.Controller, 200, true, "Indicador eliminado", nil)
}
