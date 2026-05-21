// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	":/Users/guald/OneDrive/Escritorio/go/AgroCampo_Apis_Crud_Beego/api_Crud_Contenido/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/categoria",
			beego.NSInclude(
				&controllers.CategoriaController{},
			),
		),

		beego.NSNamespace("/respuestaforo",
			beego.NSInclude(
				&controllers.RespuestaforoController{},
			),
		),

		beego.NSNamespace("/temaforo",
			beego.NSInclude(
				&controllers.TemaforoController{},
			),
		),

		beego.NSNamespace("/videoeducativo",
			beego.NSInclude(
				&controllers.VideoeducativoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
