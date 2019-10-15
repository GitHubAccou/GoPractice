package main
import(
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/githubaccou/cssSelector"
	"golang.org/x/net/html"
	"github.com/gookit/config"
	"strings"
	"context"
	"os"
	"fmt"
)
func main(){
	e_ctx,e_cancel:=chromedp.NewExecAllocator(context.Background(),chromedp.NoSandbox,chromedp.Headless)
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
	var tableStr string
	if err := chromedp.Run(ctx,
	    chromedp.Emulate(device.Info{"PC","",1920,1080,1.0,true,false,false}),
	    chromedp.Navigate(`https://free-ss.site`),
	    chromedp.WaitReady(`//table[@id="tbss"]/tbody`,chromedp.BySearch),
	    chromedp.OuterHTML(`html`,&tableStr,chromedp.BySearch),
	); err != nil {
		fmt.Println("访问站点出错...")
	    panic(err)
	}
	reader:=strings.NewReader(tableStr)
	fmt.Println("解析数据...")
	tbodyNode,pErr:=html.Parse(reader)
	if pErr!=nil{
		fmt.Println("解析数据错误")
		panic(pErr.Error())
	}
	lines:=cssSelector.Query(tbodyNode,"#tbss tbody tr")
	result:=[]map[string]interface{}{}
	for _,line:=range lines{
		fields:=cssSelector.Query(line,"td")
		mp:=make(map[string]interface{})
		mp["server"]=fields[1].FirstChild.Data
		mp["server_port"]=fields[2].FirstChild.Data
		mp["password"]=fields[3].FirstChild.Data
		mp["method"]=fields[4].FirstChild.Data
		mp["plugin"]=""
		mp["plugin_opts"]=""
		mp["plugin_args"]=""
		mp["remarks"]=""
		mp["timeout"]=5
		fmt.Println(mp)
		result=append(result,mp)
	}
	cfg:=config.Default();
	loadErr:=config.LoadFiles("D:/ss/gui-config.json")
	fmt.Println("加载配置...")
	if loadErr!=nil{
		fmt.Println("加载配置文件错误")
		panic(loadErr.Error())
	}
	cfg.Set("configs",result,false)
	out,_:=os.Create("D:/ss/gui-config.json")
	cfg.WriteTo(out)
	fmt.Println("修改配置...")
}
