package base


import "fmt"

type Person struct {
	name string
	age int
	address string
}

func main()  {
	person:=make(map[int] Person)

	person[0]=Person{name:"庞晨旭",age:25,address:"上海海事大学"}

	fmt.Println(person)
}