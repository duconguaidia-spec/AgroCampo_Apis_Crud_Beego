package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:EspecialidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ProductoVeterinariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ResenaVeterinariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioGeneralController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:ServicioVeterinariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaEspecialidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:TrVeterinariaServicioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"] = append(beego.GlobalControllerRouter["API_CRUD_VETERINARIAS/controllers:VeterinariaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
