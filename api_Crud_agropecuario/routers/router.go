// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"API_CRUD_AGROPECUARIO/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/categoria_ganado",
			beego.NSInclude(
				&controllers.CategoriaGanadoController{},
			),
		),

		beego.NSNamespace("/subasta",
			beego.NSInclude(
				&controllers.SubastaController{},
			),
		),

		beego.NSNamespace("/tr_precio_subasta_ganado",
			beego.NSInclude(
				&controllers.TrPrecioSubastaGanadoController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// Indicadores del sector
	beego.Router("/v1/indicador_sector", &controllers.IndicadorSectorController{}, "get:GetAll;post:Post")
	beego.Router("/v1/indicador_sector/:id([0-9]+)", &controllers.IndicadorSectorController{}, "get:GetOne;put:Put;delete:Delete")

	// Mediciones del sector
	beego.Router("/v1/medicion_sector", &controllers.MedicionSectorController{}, "get:GetAll;post:Post")
	beego.Router("/v1/medicion_sector/:id([0-9]+)", &controllers.MedicionSectorController{}, "get:GetOne;put:Put;delete:Delete")
}
