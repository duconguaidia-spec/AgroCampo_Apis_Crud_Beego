package controllers

import (
	"API_CRUD_VETERINARIAS/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// ProductoVeterinariaController operations for ProductoVeterinaria
type ProductoVeterinariaController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductoVeterinariaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProductoVeterinaria
// @Param	body		body 	models.ProductoVeterinaria	true		"body for ProductoVeterinaria content"
// @Success 201 {int} models.ProductoVeterinaria
// @Failure 403 body is empty
// @router / [post]
func (c *ProductoVeterinariaController) Post() {
	var v models.ProductoVeterinaria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddProductoVeterinaria(&v); err == nil {
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
// @Description get ProductoVeterinaria by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProductoVeterinaria
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductoVeterinariaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProductoVeterinariaById(id)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor GetOne: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa GetOne", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ProductoVeterinaria
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProductoVeterinaria
// @Failure 403
// @router / [get]
func (c *ProductoVeterinariaController) GetAll() {
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

	l, err := models.GetAllProductoVeterinaria(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor GetAll: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
	} else {
		if l == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 400, "Message": "Peticion exitosa GetAll: No se encontraron registros"}
		} else {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa GetAll", "data": l}
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProductoVeterinaria
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProductoVeterinaria	true		"body for ProductoVeterinaria content"
// @Success 200 {object} models.ProductoVeterinaria
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductoVeterinariaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ProductoVeterinaria{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProductoVeterinariaById(&v); err == nil {
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
// @Description delete the ProductoVeterinaria
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductoVeterinariaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteProductoVeterinaria(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa Delete", "id": id}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Delete: La solicitud contiene un parametro incorrecto o no existe el recurso solicitado"}
	}
	c.ServeJSON()
}
