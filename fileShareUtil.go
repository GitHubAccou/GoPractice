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
	var dirPrefix="dir="
	var portPrefix="port="
	args:=os.Args 
	for _,v:=range args{
		if strings.HasPrefix(v,dirPrefix){
			dir=strings.TrimPrefix(v,dirPrefix)
		}else if strings.HasPrefix(v,portPrefix){
			port,_=strconv.Atoi(strings.TrimPrefix(v,portPrefix))
		}
	}
	if dir==""{
		panic(errors.New("请用dir=XXX 参数指定目录！"))
	}
	if port==-1{
		panic(errors.New("请用port=xxx 参数指定端口！"))
	}

	http.Handle("/",http.FileServer(http.Dir(dir)))
	var err error
	try: err=http.ListenAndServe(":"+strconv.Itoa(port),nil)
	if err!=nil{
		fmt.Printf("%T",err)
		fmt.Println(port,"端口被占用，正在选用随机端口尝试启动。。。")
		port=rand.Intn(65535)
		goto try
	}else{
		fmt.Println("启动成功:\n\t使用端口：",port,"\n\t共享目录:",dir)
	}

}