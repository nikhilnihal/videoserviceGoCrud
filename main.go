package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}


func init() {
	// register model
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:pw@/yp?charset=utf8")

	orm.RegisterModel(new(User))	
}


func main() {

	fmt.Println("hello world")
	fmt.Println("hey hi")
	fmt.Println("Adding lambda")

	o := orm.NewOrm()

	user := User{Name: "slene"}

	// // insert
	_, err := o.Insert(&user)
	if err !=nil {
		fmt.Print(err)
		fmt.Println("wron")
	}

	// // update
	user.Name = "astaxie"
	_, err = o.Update(&user)
	if err !=nil {
		fmt.Println("wron")
	}


	// // read one
	u := User{Id: user.Id}
	err = o.Read(&u)

	// delete
	//num, err = o.Delete(&u)	
}