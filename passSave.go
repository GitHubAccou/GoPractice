package main
import (
	"math/rand"
	"encoding/base64"
	"strings"
	"io/ioutil"
	"errors"
	"os"
	"fmt"
	"flag"
)

var cs []rune=[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
func main(){
	act:=flag.String("act","e","-act=e表示加密 -act=d表示解密 默认解密")
	pass:=flag.String("pass","","-pass 表示密码，解密会用到")
	sF:=flag.String("sF","","-sF 源文件")
	dF:=flag.String("dF","","-dF 处理后的文件")
	flag.Parse()
	// fmt.Printf("act\t:\t%s\tpass\t:\t%s\tsF\t:\t%s\tdF\t:\t%s\n",*act,*pass,*sF,*dF)
	params:=map[string]string{"act":*act,"pass":*pass,"sF":*sF,"dF":*dF}
	for k,v:=range params{
		if len(v)<1{
			fmt.Println("缺少参数:",k)
		}
	}
	err:=error(nil)
	switch(*act){
		case "e":{
			err=encFile(*sF,*pass,*dF)
		}
		case "d":{
			err=decFile(*sF,*pass,*dF)
		}
		case "help":{
			fmt.Println(`命令格式：passSave(.exe) -act=act -pass=pass -sF=sF -dF=dF`)
		}
		default:{
			err=errors.New("未识别的指令 调用 -help 查看帮助")
		}
	}
	if err!=nil{
		fmt.Println(err.Error())
	}
}

//pick a Seed String
func randSeed()string{
	l:=len(cs)
	cs1:=make([]rune,l)
	rcs:=make([]rune,l)
	copy(cs1,cs)
	for i:=0;i<l;i++{
		ll:=len(cs1)
		idx:=rand.Intn(ll)
		rcs[i]=cs1[idx]
		cs1_c:=make([]rune,ll-1)
		copy(cs1_c,cs1[:idx])
		if idx<ll-1{
			for x:=idx+1;x<ll;x++{
				cs1_c[x-1]=cs1[x]
			}
		}
		cs1=cs1_c
		if len(cs1_c)>200{
			break;
		}
	}
	return string(rcs)
}

func readAllLines(file string)(res []string, err error){
	if data,err1:=ioutil.ReadFile(file);err!=nil{
		res,err=nil,err1
	}else{
		str:=string(data)
		lines:=strings.Split(str,"\n")
		res,err=lines,nil
	}
	return res,err
}
//check password
func checkPass(pass string)error{
	if passlen:=len(pass);passlen<8||passlen>18{
		return errors.New("密码长度应该在8-18之间")
	}
	return nil
}

func encrypt(content []byte,seed []byte)[]byte{
	l:=len(seed)
	for i,b:=range content{
		content[i]=b^seed[i%l]
	}
	return content;
}

func encFile(oriFile ,password ,dstFile string)error{
	//1.check if pass valid
	if err:=checkPass(password);err!=nil{
		return err
	}	

	//3.give a random base64 seed

	seed:=randSeed()
	//4.seed enc the password
	encoder:=base64.NewEncoding(seed)
	contentEncSeed:=encoder.EncodeToString([]byte(password))
	header:=seed[:32]+contentEncSeed+seed[32:]

	//5.read File Contents
	
	buf,err:=ioutil.ReadFile(oriFile)
	if err!=nil{
		return err
	}
	
	out:=encrypt(buf,[]byte(contentEncSeed))
	outContent:=[]byte(header)
	outContent=append(outContent,byte('\n'))
	outContent=append(outContent,out...)
	if err=ioutil.WriteFile(dstFile,outContent,os.ModePerm);err!=nil{
		return err
	}
	return nil
}
func decFile(oriFile ,password ,dstFile string)error{
	data,err1:=ioutil.ReadFile(oriFile);
	if err1!=nil{
		return err1
	}
	headerIndex:=-1;
	for i:=0;i<85;i++{
		if(data[i]==byte('\n')){
			headerIndex=i;
			break;
		}
	}
	if(headerIndex==-1){
		return errors.New("解析错误")
	}
	header:=string(data[:headerIndex])
	headerl:=len(header)
	seed:=header[0:32]+header[headerl-32:]
	//4.seed enc the password
	encoder:=base64.NewEncoding(seed)
	dout,err1:=encoder.DecodeString(header[32:headerl-32])
	if err1!=nil{
		return err1
	}
	storePass:=string(dout)
	if storePass!=password{
		return errors.New("incorrect password")
	}
	contentEncSeed:=encoder.EncodeToString([]byte(password))
	contentByteArr:=data[headerIndex+1:]
	out:=encrypt(contentByteArr,[]byte(contentEncSeed))
	if err:=ioutil.WriteFile(dstFile,out,os.ModePerm);err!=nil{
		return err
	}
	return nil
}