package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func TestEdicionForoRequiereClaveInterna(t *testing.T) {
	clave := strings.Repeat("a", 32)
	t.Setenv("AGROCAMPO_CLAVE_INTERNA", clave)
	for _, caso := range []struct {
		metodo string
		clave  string
		codigo int
	}{
		{http.MethodPut, "", http.StatusForbidden},
		{http.MethodPatch, "", http.StatusForbidden},
		{http.MethodPatch, clave, http.StatusBadRequest},
	} {
		r := httptest.NewRequest(caso.metodo, "/v1/temaforo/invalido", strings.NewReader(`{}`))
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("X-Agrocampo-Interno", caso.clave)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)
		if w.Code != caso.codigo {
			t.Errorf("%s: se esperaba HTTP %d, se obtuvo %d (%s)", caso.metodo, caso.codigo, w.Code, w.Body)
		}
	}
}
