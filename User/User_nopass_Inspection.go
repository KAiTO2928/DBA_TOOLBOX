package User

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

func User_nopass_Inspection(db *sql.DB) {
	sqlStr := "select user,host from mysql.user where authentication_string='';"
	counts, err := db.Query(sqlStr)
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
		fmt.Printf("%d.【user: %s】【host: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【user: %s】【host: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME)
		}
	} else {
		fmt.Println("————————————————————【没有无密码用户】——————————————————")
	}
	fmt.Printf("———————————————————↑无密码用户巡检完毕↑—————————————————\n\n")
}
