package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:RespuestaforoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:TemaforoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"] = append(beego.GlobalControllerRouter["api_crud_contenido/controllers:VideoeducativoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
