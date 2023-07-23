package Privileges

import (
	"DB_OSInspection/Global"
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
	Yellow    = color.Yellow.Render
	Red       = color.Red.Render
	Completed = color.S256(255, 27)
)

var user_username_slice = make([]string, 0, 30)
var user_userhost_slice = make([]string, 0, 30)

func Privileges_Inspection() {
	sqlStr := "select user,host from mysql.user where user not in ('mysql.session','mysql.sys','mysql.infoschema');"
	counts, err := Global.DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	for counts.Next() {
		var user User
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
		user_username_slice = append(user_username_slice, user.TABLE_SCHEMA)
		user_userhost_slice = append(user_userhost_slice, user.TABLE_NAME)
		if err != nil {
			panic(err.Error())
		}
		count += 1
		fmt.Printf("%d.【USER: %s】【HOST: %s】\n", count, Green(user.TABLE_SCHEMA), Red(user.TABLE_NAME))
	}

}

func User_Privileges_Inspection() {
	var user User
	for j := 0; j < len(user_username_slice); j++ {
		userName := user_username_slice[j]
		userHost := user_userhost_slice[j]
		sqlStr := "show grants for" + "'" + userName + "'@'" + userHost + "'"
		err := Global.DB.QueryRow(sqlStr).Scan(&user.TABLE_SCHEMA)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("【grants: %s\n】", user.TABLE_SCHEMA)

	}
	Completed.Printf("——————————————————↑用户权限巡检完毕↑—————————————————————")
	fmt.Println(" \n ")
}
