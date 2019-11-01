package main
import(
	"fmt"
	"os"
	"net/http"
	"strings"
	"strconv"
	"errors"
	"math/rand"

)
var dir string=""
var port int=-1
func main(){
	args:=os.Args 
	if len(args)>1{
		dir=args[1]
	}else{
		panic(errors.New("请指定目录！"))
	}
	if strings.TrimSpace(dir)==""{
		panic(errors.New("请指定目录！"))
	}
	http.Handle("/",http.FileServer(http.Dir(dir)))
	for ;;{
		port=rand.Intn(65535)
		fmt.Println(`使用`,port,`端口启动。。。`)
		err:=http.ListenAndServe(":"+strconv.Itoa(port),nil)
		fmt.Println("---->",err,"<----")
		if err!=nil{
			fmt.Printf("%T",err)
		}
	}

}