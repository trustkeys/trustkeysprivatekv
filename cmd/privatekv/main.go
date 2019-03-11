// package main

// import (
// 	_ "github.com/trustkeys/trustkeysprivatekv/routers"
// 	"github.com/trustkeys/trustkeysprivatekv/appconfig"
// 	"github.com/trustkeys/trustkeysprivatekv/controllers"
// 	"os"
// 	"github.com/astaxie/beego"
// 	"strconv"
// )

// func main() {
// 	// if beego.BConfig.RunMode == "dev" {
// 		beego.BConfig.WebConfig.DirectoryIndex = true
// 		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
// 	// }
// 	appconfig.InitConfig()
// 	controllers.ConfigSetBSHostPort(appconfig.BIGSETKV_HOST, strconv.Itoa(appconfig.BIGSETKV_PORT) )
	
// 	os.Setenv("HOST", appconfig.RunningHost)
// 	os.Setenv("PORT", appconfig.ListenPort)
	
// 	beego.Run()
// }

package main

import (
	_ "github.com/trustkeys/trustkeysprivatekv/routers"
	"github.com/trustkeys/trustkeysprivatekv/appconfig"
	// "github.com/trustkeys/trustkeysprivatekv/controllers"
	"os"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/trustkeys/trustkeysprivatekv/controllers"
	"github.com/trustkeys/trustkeysprivatekv/models"
	
)


func InitWithBSHostPort(bsHost, bsPort string) {
	controllers.SetPrivateModel(models.NewtrustkeysprivatekvAcceptAllModel(bsHost, bsPort) , true)
}


func main() {
	// if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.WebConfig.StaticDir["/v1/privatekv/swagger"] = "swagger"
	// }
	appconfig.InitConfig()
	InitWithBSHostPort(appconfig.BIGSETKV_HOST, strconv.Itoa(appconfig.BIGSETKV_PORT) )
	
	os.Setenv("HOST", appconfig.RunningHost)
	os.Setenv("PORT", appconfig.ListenPort)
	
	beego.Run()
}
