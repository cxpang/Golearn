package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"io"
	"os"
)


func a(data string) string {
	t := md5.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x \r\n",t.Sum(nil))
}
//对字符串进行SHA1哈希
func b(data string) string {
	t := sha1.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
func main(){
	var data string = "abc"
	userFile:="pangchenxu.txt"
	_,er:=os.Open(userFile)
	if er!=nil {
		file,err:= os.Create(userFile)
		if err!=nil{
			fmt.Println("创建失败")
			return
		}
		defer file.Close()
        fmt.Println(a(data))
		//file.WriteString(a(data))
		//file.WriteString(b(data))

	}
}