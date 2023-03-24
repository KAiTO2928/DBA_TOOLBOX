package Index

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
}

func Index_redundant_Inspection(db *sql.DB) {
	sqlStr := "SELECT table_schema,table_name,redundant_index_name,redundant_index_columns FROM sys.schema_redundant_indexes GROUP BY table_schema,table_name,redundant_index_name,redundant_index_columns;"
	counts, err := db.Query(sqlStr)
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
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【redundant_index_name: %s】【redundant_index_columns:%s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.Field_Type)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【redundant_index_name: %s】【redundant_index_columns:%s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.Field_Type)
		}
	} else {
		fmt.Println("————————————————————【没有重复索引的表】—————————————————")
	}
	fmt.Printf("—————————————————↑重复索引表巡检完毕↑————————————————\n\n")
}
func Index_columns_Inspection(db *sql.DB) {
	var index_columns int = 5
	sqlStr := "SELECT s.table_schema, s.table_name,s.index_name,s.column_name FROM information_schema.STATISTICS s,(SELECT table_name,index_name,count(*)FROM information_schema.STATISTICS WHERE table_schema NOT IN ( 'information_schema', 'performance_schema', 'mysql', 'sys' ) GROUP BY table_name,index_name HAVING count(*)> ?) t WHERE s.table_name = t.table_name AND s.index_name = t.index_name;"
	counts, err := db.Query(sqlStr, index_columns)
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
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】【column_name:%s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.Field_Type)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】【column_name:%s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.Field_Type)
		}
	} else {
		fmt.Printf("————————————————————【没有索引列超%d过个的索引】———————————————\n", index_columns)
	}
	fmt.Printf("—————————————————↑索引列超过%d个的索引巡检完毕↑————————————————\n\n", index_columns)

}
func Index_unused_Inspection(db *sql.DB) {
	sqlStr := "select * from sys.schema_unused_indexes;"
	counts, err := db.Query(sqlStr)
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
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【index_name: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME)
		}
	} else {
		fmt.Printf("————————————————————【没有无用的索引】———————————————\n")
	}
	fmt.Printf("—————————————————↑无用的索引巡检完毕↑————————————————\n\n")

}
