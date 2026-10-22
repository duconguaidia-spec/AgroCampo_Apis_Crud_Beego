package controllers

import beego "github.com/beego/beego/v2/server/web"

// responder arma la respuesta JSON igual para todos los controllers
func responder(c *beego.Controller, codigo int, exito bool, mensaje string, datos interface{}) {
	c.Ctx.Output.SetStatus(codigo)
	c.Data["json"] = map[string]interface{}{
		"Success": exito,
		"Status":  codigo,
		"Message": mensaje,
		"Data":    datos,
	}
	c.ServeJSON()
}
