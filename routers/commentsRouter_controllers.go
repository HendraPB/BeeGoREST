package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"] = append(beego.GlobalControllerRouter["BeeGoREST/controllers:CountryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
