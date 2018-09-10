package main

import (
	"fmt"
	"os"
)

func main()  {
	userFile:="pangchenxu.txt"
	fout,err:=os.Create(userFile)
	defer fout.Close()
	if err!=nil{
		fmt.Println(userFile,err)
		return
	}
	for i:=0;i<10 ;i++  {
		fout.WriteString("just as test \r\n")
		fout.Write([]byte("just as test \r\n"))
	}
	fl,err:=os.Open(userFile)
	defer fl.Close()
	if err!=nil{
		fmt.Println(userFile,err)
		return

	}
	buf:=make([]byte,1024)
	for {
		n,_:=fl.Read(buf)
		if 0==n{
			break
		}
		os.Stdout.Write(buf[:n])
	}

}

