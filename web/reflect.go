package main
import "fmt"
import "reflect"
func main() {
	var x float64=3.4
	v:=reflect.ValueOf(x)
	fmt.Println("type:",v.Type())
	fmt.Println("kind si float64:",v.Kind()==reflect.Float64)
	fmt.Println("value:",v.Float())
}