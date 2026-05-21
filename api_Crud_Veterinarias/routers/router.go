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

		beego.NSNamespace("/Veterinaria",
			beego.NSInclude(
				&controllers.VeterinariaController{},
			),
		),

		beego.NSNamespace("/tr_VeterinariaEspecialidad",
			beego.NSInclude(
				&controllers.TrVeterinariaEspecialidadController{},
			),
		),

		beego.NSNamespace("/Especialidad",
			beego.NSInclude(
				&controllers.EspecialidadController{},
			),
		),

		beego.NSNamespace("/tr_VeterinariaServicio",
			beego.NSInclude(
				&controllers.TrVeterinariaServicioController{},
			),
		),

		beego.NSNamespace("/ServicioGeneral",
			beego.NSInclude(
				&controllers.ServicioGeneralController{},
			),
		),

		beego.NSNamespace("/ServicioVeterinaria",
			beego.NSInclude(
				&controllers.ServicioVeterinariaController{},
			),
		),

		beego.NSNamespace("/ProductoVeterinaria",
			beego.NSInclude(
				&controllers.ProductoVeterinariaController{},
			),
		),

		beego.NSNamespace("/ResenaVeterinaria",
			beego.NSInclude(
				&controllers.ResenaVeterinariaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
