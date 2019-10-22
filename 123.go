package main
import(
	"fmt"
	"github.com/chromedp/chromedp"
	"context"
	"io/ioutil"
	"encoding/base64"
	"strings"
)
var url="https://kyfw.12306.cn/otn/resources/login.html"
func main(){
	dataurl:=getImgDataUrl()
	dataUrl2Img(dataurl,"img.png",true)
}

func getImgDataUrl() string{
	ctx,cancel:=chromedp.NewContext(context.Background())
	defer cancel()
	var screenBuf []byte
	var dataUrl string
	var getDataUrl bool
	err:=chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`//ul[@class="login-hd"]`,chromedp.BySearch),
		chromedp.Click(`//ul[@class="login-hd"]/li[@class="login-hd-account"]`,chromedp.BySearch),
		chromedp.CaptureScreenshot(&screenBuf),
		chromedp.AttributeValue(`//img[@id="J-loginImg"]`,"src",&dataUrl,&getDataUrl,chromedp.BySearch),
	)
	if err!=nil{
		fmt.Println("请求站点错误")
		panic(err.Error())
	}
	return dataUrl
}
func dataUrl2Img(dataurl string,file string,trimPrefix bool){
	if trimPrefix{
		dataurl=strings.Split(dataurl,",")[1]
	}
	imgBuf,err:=base64.StdEncoding.DecodeString(dataurl)
	if err!=nil{
		panic(err.Error())
	}
	ioutil.WriteFile(file,imgBuf,0666)

}