package main
import(
	"fmt"
)

type Shape interface{
	area() float32
}
type Cube struct{
	Len float32
}
type Circle struct{
	R float32
}
func(cube Cube)area()float32{
	return cube.Len*cube.Len
}
func(circle *Circle)area()float32{
	return 3.14*circle.R*circle.R
}


func main(){
	//实现方法是指方法的receiver是类型本身，而不是类型的指针！！
	var s Shape=Cube{4}
	fmt.Println(s.area())
	// var s1 Shape=Circle{3}    // 如果不注释掉，该行会报错，因为Circle并没有实现Shape接口
	s1:=&Circle{3}
	fmt.Println(s1.area()) //虽然Circle 没有实现Shape接口，但是他拥有area方法
}
