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
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var tableStr string
	if err := chromedp.Run(ctx,
	    chromedp.Emulate(device.Info{"PC","",1920,1080,1.0,true,false,false}),
	    chromedp.Navigate(`https://www.youneed.win/free-ss`),
	    chromedp.WaitVisible(`#post-box`,chromedp.ByID),
	    chromedp.OuterHTML(`html`,&tableStr,chromedp.BySearch),
	); err != nil {
	    panic(err)
	}
	reader:=strings.NewReader(tableStr)
	tbodyNode,pErr:=html.Parse(reader)
	if pErr!=nil{
		fmt.Println("解析数据错误")
		panic(pErr.Error())
	}
	lines:=cssSelector.Query(tbodyNode,"#post-box tbody tr")
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
		result=append(result,mp)
	}
	cfg:=config.Default();
	loadErr:=config.LoadFiles("./gui-config.json")
	if loadErr!=nil{
		fmt.Println("加载配置文件错误")
		panic(loadErr.Error())
	}
	cfg.Set("configs",result,false)
	out,_:=os.Create("./gui-config.json")
	cfg.WriteTo(out)
}
