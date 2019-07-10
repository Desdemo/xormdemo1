package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Come struct {
	Id string `json:"id"`
	Price string `json:"price"`
	Time string `json:"time"`
	Client string `json:"client"`
}

var comes []Come

func main() {
	f, err := excelize.OpenFile("./bk.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取 Sheet1 上指定单元格
	rows, err := f.GetRows("2019年")
	for i := 3 ; i < (len(rows)) ; i ++ {
		com := new(Come)
		com.Id = rows[i][0]
		com.Price = rows[i][3]
		com.Time = rows[i][1]
		com.Client = rows[i][4]
		// 添加到comes
		comes = append(comes, *com)
	}

	fmt.Println(comes)
}
