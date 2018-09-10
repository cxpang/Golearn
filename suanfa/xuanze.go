package main

import (
	"fmt"
)

func main()  {
	a:=[]int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
    fmt.Println("排序之前的数组")
	fmt.Println(a)
	selectsort(a)
	fmt.Println(a)
}
func selectsort(items []int){
	n:=len(items)
	for i:=0;i<n ;i++  {
		minIndex:=i
		for j:=i;j<n ;j++  {
			if items[j]<items[minIndex]{
				minIndex=j
			}
		}
		items[i],items[minIndex]=items[minIndex],items[i]
	}
	fmt.Println("选择排序之后的数据")

}
