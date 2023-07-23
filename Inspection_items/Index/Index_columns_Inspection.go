package Index

import (
	"dba_toolbox/Global"
	"fmt"

	"github.com/gookit/color"
)

type User struct {
	TABLE_SCHEMA       string
	TABLE_NAME         string
	CHARACTER_SET_NAME string
	Field_Type         string
}

var (
	// color
	Green     = color.Green.Render
	Yellow    = color.Yellow.Render
	Red       = color.Red.Render
	Completed = color.S256(255, 27)
)

func Index_redundant_Inspection() {
	sqlStr := "SELECT table_schema,table_name,redundant_index_name,redundant_index_columns FROM sys.schema_redundant_indexes GROUP BY table_schema,table_name,redundant_index_name,redundant_index_columns;"
	counts, err := Global.DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【redundant_index_name: %s】【redundant_index_columns:%s】\n", count, Yellow(user.TABLE_SCHEMA), Green(user.TABLE_NAME), Red(user.CHARACTER_SET_NAME), Red(user.Field_Type))
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【redundant_index_name: %s】【redundant_index_columns:%s】\n", count, Yellow(user.TABLE_SCHEMA), Green(user.TABLE_NAME), Red(user.CHARACTER_SET_NAME), Red(user.Field_Type))
		}
	} else {
		color.BgGreen.Println("————————————————————【没有重复索引的表】————————————————")

	}
	Completed.Printf("———————————————————↑重复索引表巡检完毕↑——————————————————")
	fmt.Println(" \n ")

}
func Index_columns_Inspection() {
	var index_columns int = 5
	sqlStr := "SELECT s.table_schema, s.table_name,s.index_name,s.column_name FROM information_schema.STATISTICS s,(SELECT table_name,index_name,count(*)FROM information_schema.STATISTICS WHERE table_schema NOT IN ( 'information_schema', 'performance_schema', 'mysql', 'sys' ) GROUP BY table_name,index_name HAVING count(*)> ?) t WHERE s.table_name = t.table_name AND s.index_name = t.index_name;"
	counts, err := Global.DB.Query(sqlStr, index_columns)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】【column_name:%s】\n", count, Yellow(user.TABLE_SCHEMA), Green(user.TABLE_NAME), Red(user.CHARACTER_SET_NAME), Red(user.Field_Type))
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】【column_name:%s】\n", count, Yellow(user.TABLE_SCHEMA), Green(user.TABLE_NAME), Red(user.CHARACTER_SET_NAME), Red(user.Field_Type))
		}
	} else {
		color.BgGreen.Printf("—————————————————【没有索引列超%d过个的索引】————————————", index_columns)
		fmt.Println(" \n ")

	}
	Completed.Printf("——————————————↑索引列超过%d个的索引巡检完毕↑————————————", index_columns)
	fmt.Println(" \n ")

}
func Index_unused_Inspection() {
	sqlStr := "select * from sys.schema_unused_indexes;"
	counts, err := Global.DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】\n", count, Yellow(user.TABLE_SCHEMA), Green(user.TABLE_NAME), Red(user.CHARACTER_SET_NAME))
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】\n", count, Yellow(user.TABLE_SCHEMA), Green(user.TABLE_NAME), Red(user.CHARACTER_SET_NAME))
		}
	} else {
		color.BgGreen.Printf("——————————————————————【没有无用的索引】—————————————————\n")
	}
	Completed.Printf("———————————————————↑无用的索引巡检完毕↑——————————————————")
	fmt.Println(" \n ")

}
