package Privileges

import (
	"database/sql"
	"fmt"
)

type User struct {
	TABLE_SCHEMA       string
	TABLE_NAME         string
	CHARACTER_SET_NAME string
	Field_Type         string
	INDEX_COUNT        int
	SIZE_IN_GB         float32
}

// var user_user [100]string
// var user_usera_slice []string = user_user[:]
var user_usera_slice = make([]string, 0, 3)

func Privileges_Inspection(db *sql.DB) {
	sqlStr := "select user,host from mysql.user where user not in ('mysql.session','mysql.sys','mysql.infoschema');"
	counts, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	for counts.Next() {
		var user User
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
		user_usera_slice = append(user_usera_slice, user.TABLE_SCHEMA)

		if err != nil {
			panic(err.Error())
		}
		count += 1
		fmt.Printf("%d.【USER: %s】【HOST: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME)
	}

}

func User_Privileges_Inspection(db *sql.DB) {
	var user User
	for j := 0; j < len(user_usera_slice); j++ {
		values := user_usera_slice[j]
		sqlStr := "show grants for" + "'" + values + "'"
		err := db.QueryRow(sqlStr).Scan(&user.TABLE_SCHEMA)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("【name: %s\n", user.TABLE_SCHEMA)

	}
	fmt.Printf("——————————————————↑用户权限巡检完毕↑—————————————————————\n\n")
}
