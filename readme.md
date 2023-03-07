# eample
```golang
package main

import (
	"fmt"
	"github.com/techidea8/restgo/core"
	"github.com/techidea8/restgo/pkg/wxapp"
	"linker/config"
	"linker/initial"
	"linker/logic"
	_ "linker/rest"
)

func main() {
	fmt.Println(config.Banner)
	// Create an instance of the app structure
	appConfig, err := config.Configure()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	// 初始化
	initial.InitializeLogger(appConfig)

	// 初始化数据库
	initial.InitializeDataBase(appConfig, initial.AutoMigrate(appConfig), initial.InitAdmin(appConfig))

	logic.InitDysmsService(appConfig.Dysms)
	initial.InitializeCache(appConfig)
	initial.InitializeClearResTask(appConfig)
	// 初始化微信
	wxapp.InitializeMiniapp(appConfig.Miniapp)
	// 静态资源
	restapp := core.NewRestApp(&appConfig.App)
	// 启用中间件
	restapp.UseMiddleware(core.Cros, core.AccessLog, core.Authorize)
	restapp.Start()
}
```