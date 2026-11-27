package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func TestEdicionVeterinariaRequiereClaveInterna(t *testing.T) {
	t.Setenv("AGROCAMPO_CLAVE_INTERNA", strings.Repeat("a", 32))
	for _, ruta := range []string{"/v1/veterinaria/1", "/v1/producto_veterinaria/1"} {
		r := httptest.NewRequest(http.MethodPut, ruta, strings.NewReader(`{}`))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		if w.Code != http.StatusForbidden {
			t.Errorf("%s: se esperaba HTTP 403, se obtuvo %d (%s)", ruta, w.Code, w.Body)
		}
	}
}
