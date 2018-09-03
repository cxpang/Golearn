package main 
import "fmt"
type Human struct{
	name string
	age int
	phone string 
}
type Student struct{
	Human
	school string
	loan float32
}
type Employee struct{
	Human
	company string
	money float32
}
func (h Human) sayhi(){
	fmt.Printf("hi,我是%s,你可以打我电话%s \n",h.name,h.phone)
}
func (h Human) sing(song string){
    fmt.Printf("我在唱歌。。。%s \n",song)
}
func (e Employee) sayhi() {
	fmt.Printf("hi,我是%s,我在%s地方工作,你可以打我电话%s \n",e.name,e.company,e.phone)
}
type Men interface{
	sayhi()
	sing(song string)
}
func main() {
	mike :=Student{Human{"mike",25,"13262618330"},"上海海事大学",5000}
	paul :=Employee{Human{"paul",30,"18133745211"},"上海浅橙",6000}
	var i Men
	i=mike
	i.sayhi()
	i.sing("我爱'我家")

	i=paul
	i.sayhi()
	i.sing("凉凉")

}