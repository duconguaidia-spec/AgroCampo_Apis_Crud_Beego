// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_Crud_usuarios/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/VerificacionDosPasos",
			beego.NSInclude(
				&controllers.VerificacionDosPasosController{},
			),
		),

		beego.NSNamespace("/AuditoriaUsuario",
			beego.NSInclude(
				&controllers.AuditoriaUsuarioController{},
			),
		),

		beego.NSNamespace("/Contrasena",
			beego.NSInclude(
				&controllers.ContrasenaController{},
			),
		),

		beego.NSNamespace("/PerfilExtendido",
			beego.NSInclude(
				&controllers.PerfilExtendidoController{},
			),
		),

		beego.NSNamespace("/Rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/TokenRecuperacion",
			beego.NSInclude(
				&controllers.TokenRecuperacionController{},
			),
		),

		beego.NSNamespace("/Usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
