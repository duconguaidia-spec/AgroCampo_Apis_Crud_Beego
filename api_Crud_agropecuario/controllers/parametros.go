package controllers

import (
	"errors"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// leerParametrosLista lee lo que viene en la URL de un GET de lista:
//
//	?query=nombre:Feria&sortby=id&order=desc&limit=20&offset=0&fields=Id,Nombre
//
// si no mandan nada, trae 10 registros desde el primero
func leerParametrosLista(c *beego.Controller) (query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64, err error) {
	query = make(map[string]string)
	limit = 10

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, e := c.GetInt64("limit"); e == nil {
		limit = v
	}
	if v, e := c.GetInt64("offset"); e == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, condicion := range strings.Split(v, ",") {
			partes := strings.SplitN(condicion, ":", 2)
			if len(partes) != 2 {
				return nil, nil, nil, nil, 0, 0, errors.New("query debe verse asi: campo:valor,campo:valor")
			}
			query[partes[0]] = partes[1]
		}
	}
	return query, fields, sortby, order, offset, limit, nil
}
