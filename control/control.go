package control

import (
	"./core"
	"./router"
)

//控制器主程序
//该函数用于启动整个项目
func Control(){
	//读取配置文件信息
	var configSrc string
	configSrc = "config" + core.PathSeparator + "config.json"
	var configData interface{}
	var b bool
	configData,b = core.LoadConfig(configSrc)
	if b == false{
		core.SendLog("无法读取config.json配置数据。")
		return
	}
	//连接数据库
	var db core.Database
	b = db.Connect(configData["database"]["type"].(string),configData["database"]["dns"].(string))
	if b == false{
		core.SendLog("无法连接到数据库。")
		return
	}
	//启动服务器
	router.RunSever(configData["server"]["host"].(string))
}
