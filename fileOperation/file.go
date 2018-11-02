package main 
import(
	"fmt"
	// "os"
	// "io"
	"path/filepath"
)
func main(){
	// //1.列出指定目录下的所有文件和目录,并显示文件信息
	// dic:="E:/"
	// file,err:=os.Open(dic)
	// if err!=nil{
	// 	panic(err.Error())
	// }else{
	// 	fis,err1:=file.Readdir(-1)//参数用来限制返回的结果数量，参数<0时不做限制
	// 	if err1!=nil{
	// 		panic(err1.Error())
	// 	}else{
	// 		fmt.Println("INDEX\t,NAME\tLENGTH\tMODE\tMODIFY_TIME\tFILE/DIC")
	// 		for i,v:=range fis{
	// 			typeStr:="FILE"
	// 			if v.IsDir() {
	// 				typeStr="DIC"
	// 			}
	// 			fmt.Println(i,"\t",v.Name(),"\t",v.Size(),"\t",v.Mode(),"\t",v.ModTime(),"\t",typeStr)
	// 		}
	// 	}
	// }
	// // 2.显示指定文件的信息
	// fileP:="E:/EasterGitRepositorys/GoPractice/fileOperation/file.go"
	// fileF,err:=os.Open(fileP)
	// if err!=nil{
	// 	panic(err.Error())
	// }else{
	// 	v,err:=fileF.Stat()
	// 	if err!=nil{
	// 		panic(err.Error())
	// 	}else{
	// 		fmt.Println("NAME\tLENGTH\tMODE\tMODIFY_TIME\tIsDic")
	// 		fmt.Println(v.Name(),"\t",v.Size(),"\t",v.Mode(),"\t",v.ModTime(),"\t",v.IsDir())
	// 	}
	// }
	// // 3.获取文件所在目录
	// dirStr:=filepath.Dir("/s_txf\\zzz/ddd.doc")// 该方法不会检查参数代表的路径是否存在,只是做计算
	// fmt.Println(dirStr)
	// 4.判断字符串路径是否是绝对路径

}