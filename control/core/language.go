package core

//该模块用于获取语言配置信息
//依赖内部模块：
// core.SendLog()
// core.LoadConfig()
//依赖外部库：无

//读取语言数据并返回
//param name string 语言文件名称部分
//return interface{},bool 配置信息，是否成功
func LoadLanguage(name string) (interface{},bool) {
	var res interface{}
	var src string
	src = "language" + PathSeparator + name + ".json"
	if IsFile(src) == false{
		return res,false
	}
	var b bool
	res,b = LoadConfig(src)
	return res,b
}