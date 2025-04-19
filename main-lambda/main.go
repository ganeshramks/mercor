package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

func init() {
	fmt.Println("init")
	_, err := orm.GetDB("mercor")
	if err != nil {
		err = orm.RegisterDataBase("mercor", "mysql", "mercor:mercor123@/mercor?charset=utf8")
	}
	fmt.Println("error:", err)
}

func main() {

	o := orm.NewOrm()
	user := User{Name: "slene"}
	// insert
	id, err := o.Insert(&user)
	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	// delete
	num, err = o.Delete(&u)

	fmt.Println("num: ", num, "err:", err, "id: ", id)
}