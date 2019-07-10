package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xormplus/xorm"
	"log"
)

type User struct {
	Id int64
	Name string `xorm:"varchar(25) notnull unique"`
}

type Come1 struct {
	Id string `json:"id"`
	Price string `json:"price"`
	Time string `json:"time"`
	Client string `json:"client"`
}
var comes1 []Come1

type Dl interface {
	Insert(string) (int64, bool)
}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3","./test.db")
	engine.ShowSQL(true)
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	if err =  engine.Sync2(new(User)); err != nil {
		log.Fatal("数据表同步失败", err)
	}


}

func (u *User) Insert(name string) (int64, bool) {
	user := new(User)
	user.Name = name
	affected, err := engine.Insert(user)
	if err != nil {
		return affected, false
	}
	return affected, true
}

func main(){
	//f, err := excelize.OpenFile("./bk.xlsx")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 获取 Sheet1 上指定单元格
	//rows, err := f.GetRows("2019年")
	//for i := 3 ; i < (len(rows)) ; i ++ {
	//	com1 := new(Come1)
	//	com1.Id = rows[i][0]
	//	com1.Price = rows[i][3]
	//	com1.Time = rows[i][1]
	//	com1.Client = rows[i][4]
	//	// 添加到comes
	//	comes1 = append(comes1, *com1)
	//}
	//QueryClient(comes1[0].Client)
	usr := &User{}
	fmt.Println(usr)
}


//func QueryClient(user string) {
//	us := &User{Name:user}
//	hs ,err := engine.Get(us)
//	if err != nil{
//		fmt.Println(err)
//	}
//	if hs == true {
//		fmt.Println(hs, us)
//	}
//}