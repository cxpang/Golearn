package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("进入冒泡排序算法");
	t:=time.Now().Unix()
	fmt.Println(t)
	a:=[]int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	maopaosort(a)
}
func maopaosort(item []int){
	for i:=0;i<len(item)-1 ;i++  {
		for j:=i;j<len(item)-1 ;j++  {
			if item[i]>item[j]{
				item[i],item[j]=item[j],item[i]
			}
		}

	}
	fmt.Println(item)
	t1:=time.Now().Unix()
	fmt.Println(t1)
}

