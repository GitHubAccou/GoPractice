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
	var s Shape=Cube{4}
	fmt.Println(s.area())
	// var s1 Shape=Circle{3}
	s1:=&Circle{3}
	fmt.Println(s1.area())
}