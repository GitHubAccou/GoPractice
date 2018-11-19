package main 
import(
	"fmt"
	"log"
	"time"
	"math"
)

func main(){
	// fmt.Println(add(1,2))

	// fmt.Println(add([]int{1,2,3,4,5,6,7,8,9,}...))

	// x:=[]int{1,2,3,4,5}
	// y:=[]int{6,7,8,9,10,11,12,13}
	// z:=[6]int{20,21,22,23,24,25}
	// fmt.Println(append(x,y...))
	// fmt.Println(append(x,z[:]...))
	// fmt.Println(0.6+0.7)

	// fmt.Println()
	// fmt.Println()
	// bigSlowOperation()
	cal("Min",1,5)

}

func cal(name string,x,y float64){
	fmt.Println(math.name)
	fmt.Println(math.name,x,y)
}

func add(vs ...int)(z int){
	fmt.Printf("params len:%d\n",len(vs))
	for _,v:=range vs{
		z+=v	
	}
	return z
}
func reduce(a,b int)int{
	return a-b
}
func mul(a,b int)int{
	return a*b
}
func div(a,b int)int{
	return a/b
}
func bigSlowOperation() {
    defer trace("bigSlowOperation")() // don't forget the extra parentheses
    // ...lots of work…
    time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { 
        log.Printf("exit %s (%s)", msg,time.Since(start)) 
    }
}
