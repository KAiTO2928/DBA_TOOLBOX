package User

import (
	"dba_toolbox/Global"
	"fmt"

	"github.com/gookit/color"
)

type User struct {
	TABLE_SCHEMA string
	TABLE_NAME   string
}

var (
	// color
	Green     = color.Green.Render
	Red       = color.Red.Render
	Completed = color.S256(255, 27)
)

func User_nopass_Inspection() {
	sqlStr := "select user,host from mysql.user where authentication_string='';"
	counts, err := Global.DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【user: %s】【host: %s】\n", count, Green(user.TABLE_SCHEMA), Red(user.TABLE_NAME))
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【user: %s】【host: %s】\n", count, Green(user.TABLE_SCHEMA), Red(user.TABLE_NAME))
		}
	} else {
		color.BgGreen.Println("————————————————————【没有无密码用户】——————————————————")
	}
	Completed.Printf("———————————————————↑无密码用户巡检完毕↑——————————————————")
	fmt.Println(" \n ")
}
