package main

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/client/orm"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	orm.RegisterDataBase("default", "mysql", "root:nikhilnn2@/yp?charset=utf8")

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
	if err != nil {
		fmt.Print(err)
		fmt.Println("wron")
	}

	// // update
	user.Name = "astaxie"
	_, err = o.Update(&user)
	if err != nil {
		fmt.Println("wron")
	}

	// // read one
	u := User{Id: user.Id}
	err = o.Read(&u)

	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "hellooooo")
	})
	r.Run("localhost:8080")

	// delete
	//num, err = o.Delete(&u)
}
