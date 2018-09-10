package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	Num:=10000000
	var list[]int
	for i:=Num;i>0 ;i--  {
        list=append(list,rand.Intn(10000))
	}
	//fmt.Println(list)
	length:=len(list)
	for root:=length/2-1;root>=0 ;root--  {
		sort(list,root,length)
	}
	for i := length - 1; i >= 1; i-- {
		list[0], list[i] = list[i], list[0]
		sort(list, 0, i)
	}  //调整位置并建并从第一个root開始建堆.假设不明确为什么,大家多把图画几遍就应该明朗了
	fmt.Println(list)
}
func sort(list []int,root,length int)  {
	for{
		child:=2*root+1
		if child>length{
			break
		}
		if child+1<length&&list[child]<list[child+1]{child++}
		if list[root]>list[child]{
			return
		}
		list[root],list[child]=list[child],list[root]
		root=child
	}
}
