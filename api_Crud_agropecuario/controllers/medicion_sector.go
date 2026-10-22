package controllers

import (
	"API_CRUD_AGROPECUARIO/models"
	"encoding/json"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// MedicionSectorController operations for MedicionSector
type MedicionSectorController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description crear la medicion
// @router / [post]
func (c *MedicionSectorController) Post() {
	var dato models.MedicionSector
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dato)
	if err != nil {
		responder(&c.Controller, 400, false, "El json enviado no es valido", err.Error())
		return
	}

	_, err = models.AddMedicionSector(&dato)
	if err != nil {
		responder(&c.Controller, 400, false, "No se pudo crear la medicion", err.Error())
		return
	}
	responder(&c.Controller, 201, true, "Medicion creada", dato)
}

// GetOne ...
// @Title GetOne
// @Description consultar la medicion por id
// @router /:id [get]
func (c *MedicionSectorController) GetOne() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		responder(&c.Controller, 400, false, "El id no es valido", nil)
		return
	}

	dato, err := models.GetMedicionSectorById(id)
	if err != nil {
		responder(&c.Controller, 404, false, "No existe la medicion", err.Error())
		return
	}
	responder(&c.Controller, 200, true, "Medicion consultada", dato)
}

// GetAll ...
// @Title GetAll
// @Description consultar las mediciones
// @router / [get]
func (c *MedicionSectorController) GetAll() {
	query, fields, sortby, order, offset, limit, err := leerParametrosLista(&c.Controller)
	if err != nil {
		responder(&c.Controller, 400, false, "Los parametros de la consulta no son validos", err.Error())
		return
	}

	lista, err := models.GetAllMedicionSector(query, fields, sortby, order, offset, limit)
	if err != nil {
		responder(&c.Controller, 400, false, "Error consultando las mediciones", err.Error())
		return
	}

	// si no hay registros se devuelve una lista vacia, no un error
	if lista == nil {
		lista = []interface{}{}
	}
	responder(&c.Controller, 200, true, "Mediciones consultadas", lista)
}

// Put ...
// @Title Put
// @Description modificar la medicion
// @router /:id [put]
func (c *MedicionSectorController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		responder(&c.Controller, 400, false, "El id no es valido", nil)
		return
	}

	dato := models.MedicionSector{}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &dato)
	if err != nil {
		responder(&c.Controller, 400, false, "El json enviado no es valido", err.Error())
		return
	}
	dato.Id = id

	err = models.UpdateMedicionSectorById(&dato)
	if err == orm.ErrNoRows {
		responder(&c.Controller, 404, false, "No existe la medicion", nil)
		return
	}
	if err != nil {
		responder(&c.Controller, 400, false, "No se pudo modificar la medicion", err.Error())
		return
	}
	responder(&c.Controller, 200, true, "Medicion modificada", dato)
}

// Delete ...
// @Title Delete
// @Description eliminar la medicion (solo lo desactiva)
// @router /:id [delete]
func (c *MedicionSectorController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		responder(&c.Controller, 400, false, "El id no es valido", nil)
		return
	}

	err = models.DeleteMedicionSector(id)
	if err == orm.ErrNoRows {
		responder(&c.Controller, 404, false, "No existe la medicion", nil)
		return
	}
	if err != nil {
		responder(&c.Controller, 400, false, "No se pudo eliminar la medicion", err.Error())
		return
	}
	responder(&c.Controller, 200, true, "Medicion eliminada", nil)
}
