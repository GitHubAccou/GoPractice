package main
import(
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"strings"
)

func main(){
	res,err:=http.Get("https://www.pptv.com")
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
	fmt.Println(strings.Repeat("-->",deep),node.DataAtom.String())
	Echo(deep+1,node.FirstChild)
	for cu,next:=node,node.NextSibling;cu.NextSibling!=nil;{
		fmt.Println(strings.Repeat("-->",deep),next.DataAtom.String())
		cu=next
	}
}
