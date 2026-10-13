package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:AuditoriaUsuarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:ContrasenaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:PerfilExtendidoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:RolController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:TokenRecuperacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:UsuarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"] = append(beego.GlobalControllerRouter["api_Crud_usuarios/controllers:VerificacionDosPasosController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
