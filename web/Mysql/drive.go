package main

import (
	"github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)
func main(){
	db,err:=sql.Open("mysql","root:200517@/Go?charset=utf8")
	checkErr(err)
	stmt,err:=db.Prepare("insert USER set username=?, PASSWORD=?")
	checkErr(err)
	res,err:=stmt.Exec("庞晨旭","200517")
	checkErr(err)
	id,err:=res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

}
