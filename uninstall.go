package main
import(
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/cdp"
	"context"
	"fmt"
	"strings"
	// "io/ioutil"
)
var base=`http://www.dheaster208mail.com`
func main(){
	// base=`https://www.hao123.com/`
	base=strings.ReplaceAll(base,"easter","")
	sub:=strings.Split(base,`mail`)
	base=sub[0]+sub[1]
	fmt.Println(base)
	Find(base,``)
}

func Find(website,selector string)map[string]string{
	ectx,ecancel:=chromedp.NewExecAllocator(context.Background(),chromedp.Headless)
	defer ecancel()
	ctx,cancel:=chromedp.NewContext(ectx)
	defer cancel()
	// var imgBuf []byte
	var nodes []*cdp.Node
	err:=chromedp.Run(ctx,
		chromedp.Navigate(website),
		// chromedp.WaitVisible(`//body`,chromedp.BySearch),
		chromedp.Nodes(`//a`,&nodes,chromedp.BySearch),
		// chromedp.CaptureScreenshot(&imgBuf),
	)
	if err!=nil{
		panic(err.Error())
	}
	// ioutil.WriteFile(`save.png`,imgBuf,0666)
	res:=make(map[string]string)
	for _,node:=range nodes{
		fmt.Println(node.AttributeValue(`href`))
	}
	return res
}