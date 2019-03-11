package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"],
        beego.ControllerComments{
            Method: "GetSafeSliceFrom",
            Router: `/GetSafeSliceFrom/:appID/:pubKey`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("pubKey", param.IsRequired),
				param.New("appID", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"],
        beego.ControllerComments{
            Method: "GetSliceFrom",
            Router: `/GetSliceFrom/:appID/:pubKey`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("pubKey", param.IsRequired),
				param.New("appID", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"],
        beego.ControllerComments{
            Method: "GetSafe",
            Router: `/getsafe`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"],
        beego.ControllerComments{
            Method: "PutItem",
            Router: `/putitem/:appID/:pubKey`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("pubKey", param.IsRequired, param.InPath),
				param.New("appID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"] = append(beego.GlobalControllerRouter["github.com/trustkeys/trustkeysprivatekv/controllers: PrivateKVController"],
        beego.ControllerComments{
            Method: "PutSafeItem",
            Router: `/putsafeitem/:appID/:pubKey`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("pubKey", param.IsRequired, param.InPath),
				param.New("appID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

}
