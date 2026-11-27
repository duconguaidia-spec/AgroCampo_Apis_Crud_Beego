package controllers

import (
	"crypto/subtle"
	"net/http"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

func autorizarEdicionInterna(c *beego.Controller) bool {
	clave := os.Getenv("AGROCAMPO_CLAVE_INTERNA")
	if len(clave) >= 32 && subtle.ConstantTimeCompare([]byte(clave), []byte(c.Ctx.Input.Header("X-Agrocampo-Interno"))) == 1 {
		return true
	}
	c.Ctx.Output.SetStatus(http.StatusForbidden)
	c.Data["json"] = map[string]interface{}{"success": false, "message": "La edición de veterinarias requiere acceso interno"}
	c.ServeJSON()
	return false
}
