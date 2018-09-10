package main

import "fmt"
import "os"

func main()  {
	userFile:="pangchenxu.txt"
	_,err:=os.Stat(userFile)

	if err!=nil{

		fmt.Println("文件不存在.创建文件")
		os.Create(userFile)
	}

	//fileob.WriteString("chenandsnas")

	os.Remove(userFile)

}