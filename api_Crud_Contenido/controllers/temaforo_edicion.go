package controllers

import (
	"api_crud_contenido/models"
	"crypto/subtle"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (c *TemaforoController) autorizarEdicionInterna() bool {
	clave := os.Getenv("AGROCAMPO_CLAVE_INTERNA")
	if len(clave) >= 32 && subtle.ConstantTimeCompare([]byte(clave), []byte(c.Ctx.Input.Header("X-Agrocampo-Interno"))) == 1 {
		return true
	}
	c.responderEdicion(http.StatusForbidden, "La edición de temas requiere acceso interno", nil)
	return false
}

func (c *TemaforoController) Editar() {
	if !c.autorizarEdicionInterna() {
		return
	}
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil || id <= 0 {
		c.responderEdicion(http.StatusBadRequest, "ID de tema inválido", nil)
		return
	}
	var datos models.EdicionTemaForo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err != nil {
		c.responderEdicion(http.StatusBadRequest, "Envía un objeto JSON válido", nil)
		return
	}
	datos.Titulo = strings.TrimSpace(datos.Titulo)
	datos.Descripcion = strings.TrimSpace(datos.Descripcion)
	if datos.Titulo == "" || datos.Descripcion == "" || datos.CategoriaID <= 0 {
		c.responderEdicion(http.StatusBadRequest, "Completa el título, la categoría y la descripción", nil)
		return
	}
	if utf8.RuneCountInString(datos.Titulo) > 240 {
		c.responderEdicion(http.StatusBadRequest, "El título admite hasta 240 caracteres", nil)
		return
	}
	edicion, err := models.EditarTemaForo(c.Ctx.Request.Context(), id, datos)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		c.responderEdicion(http.StatusNotFound, "El tema o la categoría del foro ya no está disponible", nil)
	case err != nil:
		c.responderEdicion(http.StatusInternalServerError, "No fue posible guardar el tema", nil)
	default:
		c.responderEdicion(http.StatusOK, "", edicion)
	}
}

func (c *TemaforoController) responderEdicion(status int, mensaje string, datos interface{}) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = map[string]interface{}{"success": status < 400, "message": mensaje, "data": datos}
	c.ServeJSON()
}
