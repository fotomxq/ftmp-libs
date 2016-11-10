//网络连接包
//建议使用github.com/PuerkitoBio/goquery采集数据
package ftmplibs

import (
    "io/ioutil"
    "net/http"
    "net/url"
)

//网络通讯类构建
type SimpleHttp struct{
	sendUrl string
	sendParams map[string][]string
	proxyOn bool
}

//设定URL地址
<<<<<<< HEAD
func (simpleHttp *SimpleHttp) SetSendUrl(sendUrl string){
	simpleHttp.sendUrl = sendUrl
}

//设定参数
func (simpleHttp *SimpleHttp) SetSendParams(sendParams map[string][]string) {
	simpleHttp.sendParams = sendParams
}

//设定是否启动代理
func (simpleHttp *SimpleHttp) SetProxy(setOn bool){
	simpleHttp.proxyOn = setOn
}

//Get数据
// url - 网络地址 ; param - 参数 (url.value)
func (simpleHttp *SimpleHttp) Get() ([]byte,error){
	var Url *url.URL
	var err error
	Url,err = url.Parse(simpleHttp.sendUrl)
=======
func (this *SimpleHttp) SetSendUrl(sendUrl string){
	this.sendUrl = sendUrl
}

//设定参数
func (this *SimpleHttp) SetSendParams(sendParams map[string][]string) {
	this.sendParams = sendParams
}

//设定是否启动代理
func (this *SimpleHttp) SetProxy(setOn bool){
	this.proxyOn = setOn
}



//Get数据
// url - 网络地址 ; param - 参数 (url.value)
func (this *SimpleHttp) Get() (res []byte, err error){
	var Url *url.URL
	Url,err = url.Parse(this.sendUrl)
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	if err != nil{
		return nil, err
	}
	//转换格式
<<<<<<< HEAD
	var urlParams url.Values = simpleHttp.sendParams
=======
	var urlParams url.Values = this.sendParams
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = urlParams.Encode()
	resp,err := http.Get(Url.String())
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//post数据
<<<<<<< HEAD
func (simpleHttp *SimpleHttp) Post() ([]byte,error){
	var urlParams url.Values = simpleHttp.sendParams
	resp,err := http.PostForm(simpleHttp.sendUrl, urlParams)
=======
func (this *SimpleHttp) Post() (res []byte, err error){
	var urlParams url.Values = this.sendParams
	resp,err := http.PostForm(this.sendUrl, urlParams)
>>>>>>> 2865ccf63334543b43a4ded793e8ed3dd3a456b2
	if err != nil{
		return nil ,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}