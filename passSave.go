package main
import (
	"fmt"
	"math/rand"
	// "encoding/base64"
	"strings"
	"io/ioutil"
	"errors"
)

var cs []rune=[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
func main(){
	con:=[]byte("BBXNUTZML")
	seed:=[]byte("Asnfls+/9FG")
	fmt.Println("en:")
	after:=encrypt(con,seed)
	fmt.Println(string(after))

	fmt.Println("de:")
	
	ori:=encrypt(after,seed)
	fmt.Println(string(ori))
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
		return errors.New("password length must between 8 and 18")
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
	//2.check if file exists
	if

	//3.give a random base64 seed

	seed:=randSeed()
	//4.seed enc the password
	encoder:=base64.NewEncoding(seed)
	contentEncSeed:=encoder.EncodeToString([]byte(password))
	header:=seed[:32]+contentEncSeed+seed[32:]

	//5.read File Contents
	
	ioutil.ReadAll()
	
	

}
func decFile(oriFile ,password ,dstFile string)error{

}

func checkFile()

