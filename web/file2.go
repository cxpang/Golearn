package main

import "fmt"
import "os"
import "crypto/md5"
func main()  {
	userFile:="pangchenxu.txt"
	_,er:=os.Open(userFile)
	if er!=nil {
		file,err:= os.Create(userFile)
		if err!=nil{
			fmt.Println("创建失败")
			return
		}
		defer file.Close()
                password:=md5.New()
                password.Write([]byte("sadjanfas"))
                
		fmt.Sprintf("%s",password.Sum(nil))

	}




}
