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
	"time"
	"strconv"
	"errors"
)
var url="https://kyfw.12306.cn/otn/resources/login.html"
var cpatchaURL="http://littlebigluo.qicp.net:47720"
var buyURL="https://kyfw.12306.cn/otn/leftTicket/init?linktypeid=dc"
var stationMap map[string]string=map[string]string{"北京":"BJP","定州":"DXP"}
var postionRes string
func main(){
	ctx,err:=Login("624867029@qq.com","m127906pA4")
	if err==nil{
		fmt.Println("登录成功,开始查询车票。。。")
		queryAndBook(ctx,"北京","定州","2019-11-05")
	}else{
		fmt.Println("登录失败")
	}
}

func queryAndBook(ctx1 context.Context,from,to,date string){
	var useless interface{}
	fmt.Println(from,"\t",to,"\t",date)
	err:=chromedp.Run(ctx1,
		// chromedp.EvaluateAsDevTools(`$("#J-chepiao > div").show()`,&useless),
		chromedp.EvaluateAsDevTools(`document.querySelector("#J-chepiao > div").style.display="block"`,&useless),
		chromedp.WaitVisible(`//li[@id="J-chepiao"]//li[a="单程"]/a`,chromedp.BySearch),
		chromedp.Sleep(time.Second*3),
		chromedp.Click(`//li[@id="J-chepiao"]//li[a="单程"]/a`,chromedp.BySearch),
		chromedp.WaitVisible(`//form[@id="queryLeftForm"]`,chromedp.BySearch),
		chromedp.EvaluateAsDevTools(`document.getElementById("fromStationText").value="`+from+`"`,&useless),
		chromedp.EvaluateAsDevTools(`document.getElementById("fromStation").value="`+stationMap[from]+`"`,&useless),
		chromedp.EvaluateAsDevTools(`document.getElementById("toStationText").value="`+to+`"`,&useless),
		chromedp.EvaluateAsDevTools(`document.getElementById("toStation").value="`+stationMap[to]+`"`,&useless),
		chromedp.EvaluateAsDevTools(`document.getElementById("train_date").value="`+date+`"`,&useless),
		chromedp.Click(`//a[@id="query_ticket"]`,chromedp.BySearch),
		chromedp.WaitVisible(`//div[@id="t-list"]//tbody[1]//tr`,chromedp.BySearch),
		chromedp.Sleep(time.Second*60),
	)
	if err!=nil{
		fmt.Println("查询车票出错")
		fmt.Printf("%V",err)
		panic(err.Error())
	}
}



func Login(user,pass string)(cxtt context.Context,err error){
    e_ctx,_:=chromedp.NewExecAllocator(context.Background(),chromedp.NoSandbox)
	ctx,_:=chromedp.NewContext(e_ctx,chromedp.WithErrorf(func (x string,ops ...interface{}){
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
	var dataUrl string
	var getDataUrl bool
	var useless interface{}
	err=chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Emulate(device.Info{"PC","",1920,1080,1.0,true,false,false}),
		chromedp.WaitVisible(`//ul[@class="login-hd"]`,chromedp.BySearch),
		chromedp.Click(`//ul[@class="login-hd"]/li[@class="login-hd-account"]//a`,chromedp.BySearch),
		chromedp.WaitVisible(`//img[@id="J-loginImg"]`,chromedp.BySearch),
		chromedp.AttributeValue(`//img[@id="J-loginImg"]`,`src`,&dataUrl,&getDataUrl,chromedp.BySearch),
	)
	if err!=nil{
		fmt.Println("请求站点错误")
		panic(err.Error())
	}
	if getDataUrl{
		fmt.Println("下载验证码。。。")
		file:=`E:\EasterGitRepositorys\GoPractice\code.png`
		dataUrl2Img(dataUrl,file,true)
		fmt.Println("解析验证码。。。")
		positions:=analyseCapatcha(file)
		fmt.Println(positions)
		codeHTML:=position2HTML(positions)
		fmt.Println(codeHTML)
		err=chromedp.Run(ctx,
			chromedp.SendKeys(`//input[@id="J-userName"]`,user,chromedp.BySearch),
			chromedp.SendKeys(`//input[@id="J-password"]`,pass,chromedp.BySearch),
			chromedp.EvaluateAsDevTools(`$("#J-passCodeCoin").append("`+codeHTML+`")`,&useless),
			chromedp.Click(`//a[@id="J-login"]`,chromedp.BySearch),
			chromedp.WaitVisible(`//div[@class="center-welcome"]`,chromedp.BySearch),
		)
		if err!=nil{
			cxtt=ctx
			return cxtt,err
		}else{
			cxtt=ctx
			return cxtt,nil
		}
	}else{
		cxtt=ctx
		return cxtt,errors.New("验证码处理错误")
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
	return res
}

func position2HTML(positionStr string) string{
	positions:=strings.Split(positionStr," ")
	res:=""
	for _,v:=range positions{
		p,_:=strconv.Atoi(v)
		x,y:=35,41
		if p>4{
			y+=67
			p-=4
		}
		x+=(p-1)*69
		res+=`<div randcode='`+strconv.Itoa(x)+`,`+strconv.Itoa(y)+`' class='lgcode-active'></div>`
	}
	return res
}
