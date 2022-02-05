package tool

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"NOT NULL"`
	Age  int
	Addr string
}

func (*User) TableName() string {
	return "user"
}

func Mysqlstudy() {
	/* //mysql测试sqlx

	type stu struct {
		Id    int     `db:"sid"`
		Name  string  `db:"sname"`
		Score float64 `db:"score"`
	}

	dsn := "root:852456@tcp(localhost:3306)/test1"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	sql := "select * from t_stu"

	var stuss []stu

	err1 := db.Select(&stuss, sql)
	if err1 != nil {
		fmt.Printf("err1: %v\n", err1)
	}
	fmt.Printf("stuu: %v\n", stuss) */

	//使用gorm测试mysql

	dsn := "root:852456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err2 := gorm.Open("mysql", dsn)
	defer db.Close()
	if err2 != nil {
		fmt.Println(err2)
	}
	/* //建表
	db.Table("user").CreateTable(&User{})
	*/

	//自动建表
	db.AutoMigrate(&User{})

	// 创建数据
	//var user User = User{Name: "guang", Age: 20, Addr: "sefe"}
	s := []User{{Name: "aaa", Age: 12, Addr: "aaa"}, {Name: "bbb", Age: 12, Addr: "bbb"}}
	for i := 0; i < len(s); i++ {
		fmt.Printf("db.Create(&user).Error: %v\n", db.Create(&s[i]).Error)
	}

	var user User
	// 先查询数据 , 默认字段为 id
	db.First(&user, 1)
	fmt.Printf("user: %v\n", user)
	// 如果按其他字段查询 db.First(&user,"name=?","小黑")
	db.First(&user, "name=?", "bbb")
	fmt.Printf("user: %v\n", user)
	// 查询多条数据,最终是将数据存放在数组中,查询id>1的数据放在user地址上
	var users []User
	db.Where("id>=?", 1).Find(&users)
	fmt.Printf("users: %v\n", users)

	//删表
	db.DropTable("user")
}
