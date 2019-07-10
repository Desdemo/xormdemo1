package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xormplus/xorm"
	"log"
	"time"
)
var engine1 *xorm.Engine

type Stu struct {
	Id int64
	Name string `xorm:"varchar(25) notnull unique"`
	Price float64
	Wei float64
	Status int
	BlTime time.Time
}

func main()  {
	var err error
	engine1, err = xorm.NewEngine("sqlite3","./test1.db")
	engine1.ShowSQL(true)
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	if err =  engine1.Sync2(new(Stu)); err != nil {
		log.Fatal("数据表同步失败", err)
	}
	var a []string
	sql := "select name from stu order by id desc"
	results, err := engine1.QueryString(sql)
	for _, res := range results {
		a = append(a, res["name"])
	}
	fmt.Println(a)
	//var n string
	//n = "dongdong"
	//results := make([]Stu, 0)
	//err = engine1.Where("name = ?",n).Find(&results)
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println(results)

	//var a, b, c float64
	//a = 45.5
	//c = b
	//sql := "select *  from stu order by id desc "
	//results, err := engine1.QueryString(sql)
	//for k, res := range results {
	//	fmt.Println(k)
	//	b ,_ = strconv.ParseFloat(res["price"],64)
	//	if a >= b {
	//		a = a -b
	//		c = 0
	//		fmt.Println("实收", b, "未收", c)
	//	} else {
	//		c = b - a
	//		b = a
	//		fmt.Println("实收", b, "未收", c)
	//		break
	//	}
	//}
}
