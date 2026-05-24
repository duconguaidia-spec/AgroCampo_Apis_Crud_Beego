package controllers

import (
	"API_CRUD_VETERINARIAS/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// ResenaVeterinariaController operations for ResenaVeterinaria
type ResenaVeterinariaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ResenaVeterinariaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Resenaveterinaria
// @Param	body		body 	models.Resenaveterinaria	true		"body for Resenaveterinaria content"
// @Success 201 {int} models.Resenaveterinaria
// @Failure 403 body is empty
// @router / [post]
func (c *ResenaVeterinariaController) Post() {
	var v models.ResenaVeterinaria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddResenaVeterinaria(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "Message": "Peticion exitosa Post", "data": v}
		} else {
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Post: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Post: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Resenaveterinaria by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Resenaveterinaria
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ResenaVeterinariaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetResenaVeterinariaById(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 400, "Message": "Error en el servidor GetOne: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}//Se agrega como mejora a la api, para crear una interfas mejor con el map
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa GetOne", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Resenaveterinaria
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Resenaveterinaria
// @Failure 403
// @router / [get]
func (c *ResenaVeterinariaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllResenaVeterinaria(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		if l==nil{
			c.Data["json"] = map[string]interface{}{"success": true, "status": 400, "Message": "Peticion exitosa GetAll: No se encontraron registros"}
		} else {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa GetAll", "data": l}
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Resenaveterinaria
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Resenaveterinaria	true		"body for Resenaveterinaria content"
// @Success 200 {object} models.Resenaveterinaria
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ResenaVeterinariaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ResenaVeterinaria{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateResenaVeterinariaById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa Put"}
		} else {
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Put: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Put: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Resenaveterinaria
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ResenaVeterinariaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteResenaVeterinaria(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa Delete", "id": id}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Delete: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
	}
	c.ServeJSON()
}