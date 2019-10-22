package main
import(
	"fmt"
	"github.com/chromedp/chromedp"
	"context"
)
var url="http://littlebigluo.qicp.net:47720/"
func main(){
	res:=analyseCapatcha(`C:\Users\Administrator\Desktop\p.png`)
	fmt.Println("解析结果："+res)
}

func analyseCapatcha(filePath string) string{
	ctx,cancel:=chromedp.NewContext(context.Background())
	defer cancel()
	var res string
	err:=chromedp.Run(ctx,
		chromedp.Navigate(url),
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