package main
import(
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"strings"
)

func main(){
	res,err:=http.Get("https://www.baidu.com")
	defer res.Body.Close()
	if err!=nil{
		panic(err.Error())
	}
	node,err1:=html.Parse(res.Body)
	if err1!=nil{
		panic(err1.Error())
	}
	Echo(0,node)
}
func Echo(deep int,node *html.Node){
	if node==nil{
		return
	}
	if node.Type==html.ElementNode{
		fmt.Println(strings.Repeat("--|",deep),node.DataAtom.String())
	}
	Echo(deep+1,node.FirstChild)
	for next:=node.NextSibling;next!=nil;{
		if next.Type==html.ElementNode{
			fmt.Println(strings.Repeat("--|",deep),next.DataAtom.String())
			Echo(deep+1,next.FirstChild)
		}
		next=next.NextSibling
	}
}
