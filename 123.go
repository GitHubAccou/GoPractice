package main
import(
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	// "github.com/chromedp/cdproto/cdp"
	"context"
	"io/ioutil"
	"encoding/base64"
	"strings"
)
var url="https://kyfw.12306.cn/otn/resources/login.html"
var cpatchaURL="http://littlebigluo.qicp.net:47720"
var postionRes string
func main(){
	Login("624867@qq.com","m11506pass")
}

func Login(user,pass string){
    e_ctx,e_cancel:=chromedp.NewExecAllocator(context.Background(),chromedp.NoSandbox)
	defer e_cancel()
	ctx,cancel:=chromedp.NewContext(e_ctx,chromedp.WithErrorf(func (x string,ops ...interface{}){
		fmt.Println("Error:-------------------------------->")
		fmt.Printf(x,ops)
		fmt.Println("\n<--------------------------------------")
	}),chromedp.WithDebugf(func (x string,ops ...interface{}){
		for _,v:=range ops{
			str:=fmt.Sprintf(x,v)
			if strings.Contains(str,"www.recaptcha.net")||strings.Contains(str,`recaptcha.google.cn`)||strings.Contains(str,`www.google.com/recaptcha`){
				fmt.Println(str)
				fmt.Println("Warning:This site may use google's Recaptcha to stop robot,to Crwal this site you try without Headless mode.")
			}
		}
	}),chromedp.WithLogf(func (x string,ops ...interface{}){
		fmt.Println("Log:-------------------------------->")
		fmt.Printf(x,ops)
		fmt.Println("\n<--------------------------------------")
	}))
	defer cancel()
	var dataUrl string
	var getDataUrl bool
	var useless interface{}
	err:=chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Emulate(device.Info{"PC","",1920,1080,1.0,true,false,false}),
		chromedp.WaitVisible(`//ul[@class="login-hd"]`,chromedp.BySearch),
		chromedp.Click(`//ul[@class="login-hd"]/li[@class="login-hd-account"]`,chromedp.BySearch),
		chromedp.WaitVisible(`//img[@id="J-loginImg"]`,chromedp.BySearch),
		chromedp.AttributeValue(`//img[@id="J-loginImg"]`,`src`,&dataUrl,&getDataUrl,chromedp.BySearch),
	)
	if err!=nil{
		fmt.Println("请求站点错误")
		panic(err.Error())
	}
	if getDataUrl{
		fmt.Println("下载验证码。。。")
		file:=`code.png`
		dataUrl2Img(dataUrl,file,true)
		fmt.Println("解析验证码。。。")
		result:=analyseCapatcha(file)
		fmt.Println(result)
		err=chromedp.Run(ctx,
			chromedp.SendKeys(`//input[@id="J-userName"]`,user,chromedp.BySearch),
			chromedp.SendKeys(`//input[@id="J-password"]`,pass,chromedp.BySearch),
			chromedp.EvaluateAsDevTools(`$("#J-passCodeCoin").append("<p>test</p><p>test1</p>")`,&useless),
		)
		if err!=nil{
			fmt.Println("请求站点错误")
			panic(err.Error())
		}
	}
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

func analyseCapatcha(filePath string) string{
    e_ctx,e_cancel:=chromedp.NewExecAllocator(context.Background(),chromedp.NoSandbox)
	defer e_cancel()
	ctx,cancel:=chromedp.NewContext(e_ctx,chromedp.WithErrorf(func (x string,ops ...interface{}){
		fmt.Println("Error:-------------------------------->")
		fmt.Printf(x,ops)
		fmt.Println("\n<--------------------------------------")
	}),chromedp.WithDebugf(func (x string,ops ...interface{}){
		for _,v:=range ops{
			str:=fmt.Sprintf(x,v)
			if strings.Contains(str,"www.recaptcha.net")||strings.Contains(str,`recaptcha.google.cn`)||strings.Contains(str,`www.google.com/recaptcha`){
				fmt.Println(str)
				fmt.Println("Warning:This site may use google's Recaptcha to stop robot,to Crwal this site you try without Headless mode.")
			}
		}
	}),chromedp.WithLogf(func (x string,ops ...interface{}){
		fmt.Println("Log:-------------------------------->")
		fmt.Printf(x,ops)
		fmt.Println("\n<--------------------------------------")
	}))
	defer cancel()
	var res string
	err:=chromedp.Run(ctx,
		chromedp.Navigate(cpatchaURL),
		chromedp.WaitVisible("form",chromedp.BySearch),
		chromedp.SetUploadFiles(`//form/input[@type='file']`,[]string{filePath},chromedp.BySearch),
		chromedp.Click(`//form/input[@type="submit"]`),
		chromedp.WaitVisible("//img",chromedp.BySearch),
		chromedp.InnerHTML(`//p//font//b`,&res,chromedp.BySearch),
	)
	if err!=nil{
		fmt.Println("请求站点错误")
		panic(err.Error())
	}
	fmt.Println(res)
	return res
}
