// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"API_CRUD_VETERINARIAS/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/veterinaria",
			beego.NSInclude(
				&controllers.VeterinariaController{},
			),
		),

		beego.NSNamespace("/tr_veterinariaespecialidad",
			beego.NSInclude(
				&controllers.TrVeterinariaEspecialidadController{},
			),
		),

		beego.NSNamespace("/especialidad",
			beego.NSInclude(
				&controllers.EspecialidadController{},
			),
		),

		beego.NSNamespace("/tr_veterinariaservicio",
			beego.NSInclude(
				&controllers.TrVeterinariaServicioController{},
			),
		),

		beego.NSNamespace("/serviciogeneral",
			beego.NSInclude(
				&controllers.ServicioGeneralController{},
			),
		),

		beego.NSNamespace("/servicioveterinaria",
			beego.NSInclude(
				&controllers.ServicioVeterinariaController{},
			),
		),

		beego.NSNamespace("/producto_veterinaria",
			beego.NSInclude(
				&controllers.ProductoVeterinariaController{},
			),
		),

		beego.NSNamespace("/resena_veterinaria",
			beego.NSInclude(
				&controllers.ResenaVeterinariaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
