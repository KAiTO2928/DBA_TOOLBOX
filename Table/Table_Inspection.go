package Table

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

func Table_Data_Size_Inspection(db *sql.DB) {

	var totalSize int64 = 10737418240
	sqlStr := "SELECT TABLE_SCHEMA, TABLE_NAME, ROUND((DATA_LENGTH + INDEX_LENGTH)/1024/1024/1024, 4) AS SIZE_IN_GB FROM INFORMATION_SCHEMA.TABLES WHERE DATA_LENGTH + INDEX_LENGTH > ?  ORDER BY SIZE_IN_GB DESC;"
	counts, err := db.Query(sqlStr, totalSize)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.SIZE_IN_GB)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【SIZE_IN_GB: %f】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.SIZE_IN_GB)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.SIZE_IN_GB)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【SIZE_IN_GB: %f】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.SIZE_IN_GB)
		}
	} else {
		fmt.Println("————————————————————【没有超过10G的表】—————————————————")
	}
	fmt.Printf("—————————————————↑大于%dG数据表巡检完毕↑————————————————\n\n", totalSize/1024/1024/1024)
}

func Table_Index_Size_Inspection(db *sql.DB) {
	var indexSize int = 6
	sqlStr := "SELECT TABLE_SCHEMA, TABLE_NAME, COUNT(*) AS INDEX_COUNT FROM INFORMATION_SCHEMA.STATISTICS GROUP BY TABLE_SCHEMA, TABLE_NAME HAVING INDEX_COUNT > ?;"
	counts, err := db.Query(sqlStr, indexSize)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0

	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.INDEX_COUNT)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【INDEX_COUNT: %d】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.INDEX_COUNT)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.INDEX_COUNT)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【INDEX_COUNT: %d】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.INDEX_COUNT)
		}
	} else {
		fmt.Println("————————————————————【没有超过6个索引的表】——————————————————")
	}

	fmt.Printf("—————————————————↑超过%d个索引表巡检完毕↑————————————————\n\n", indexSize)

}
func Table_fragment_Size_Inspection(db *sql.DB) {
	var fragmentSize float32 = 0.5
	sqlStr := "SELECT TABLE_SCHEMA,TABLE_NAME,1- ( TABLE_ROWS * AVG_ROW_LENGTH )/(DATA_LENGTH + INDEX_LENGTH + DATA_FREE ) AS `fragment_pct`  FROM information_schema.TABLES WHERE TABLE_SCHEMA NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' ) AND (1- ( TABLE_ROWS * AVG_ROW_LENGTH )/(DATA_LENGTH + INDEX_LENGTH + DATA_FREE )) > ? AND ( DATA_LENGTH + INDEX_LENGTH + DATA_FREE ) > 1024 * 1024 * 1024;"
	counts, err := db.Query(sqlStr, fragmentSize)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.SIZE_IN_GB)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【FRAGMENT_PCT: %f】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.SIZE_IN_GB)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.SIZE_IN_GB)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【FRAGMENT_PCT: %f】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.SIZE_IN_GB)
		}
	} else {
		fmt.Println("—————————————————【没有超过50%碎片率的表】——————————————")

	}
	fmt.Printf("———————————————↑超过%f%%个碎片率表巡检完毕↑———————————————\n\n", fragmentSize*100)

}
func Table_rows_Size_Inspection(db *sql.DB) {
	var rowsSize int32 = 10000000
	sqlStr := "SELECT table_schema,table_name,table_rows FROM information_schema.TABLES WHERE table_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )AND table_rows > ? ORDER BY table_rows DESC;"
	counts, err := db.Query(sqlStr, rowsSize)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.SIZE_IN_GB)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【TABLE_ROWS: %f】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.SIZE_IN_GB)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.SIZE_IN_GB)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【TABLE_ROWS: %f】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.SIZE_IN_GB)
		}
	} else {
		fmt.Println("——————————————————【没有超过1千万行的表】————————————————")

	}
	fmt.Printf("—————————————————↑超过%d万行的表巡检完毕↑————————————————\n\n", rowsSize/10000)

}
func Table_chaset_Inspection(db *sql.DB) {
	var default_charset string = "utf8mb4"
	sqlStr := "SELECT DISTINCT TABLE_SCHEMA,TABLE_NAME,CCSA.CHARACTER_SET_NAME FROM INFORMATION_SCHEMA.TABLES T JOIN INFORMATION_SCHEMA.COLLATION_CHARACTER_SET_APPLICABILITY CCSA ON T.TABLE_COLLATION = CCSA.COLLATION_NAME WHERE CCSA.CHARACTER_SET_NAME <> ? ORDER BY TABLE_SCHEMA,TABLE_NAME;"
	counts, err := db.Query(sqlStr, default_charset)
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
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【CHARACTER_SET_NAME: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【CHARACTER_SET_NAME: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME)
		}
	} else {
		fmt.Println("————————————————————【没有非utf8mb4的表】————————————————")
	}
	fmt.Printf("—————————————————↑非字符集%s的表巡检完毕↑————————————————\n\n", default_charset)

}
func Table_big_columns_Inspection(db *sql.DB) {
	sqlStr := "SELECT table_schema,table_name,column_name,data_type FROM information_schema.COLUMNS WHERE data_type IN  ( 'blob', 'clob', 'text', 'medium text', 'long text' ) AND table_schema NOT IN  ('information_schema','performance_schema','mysql','sys');"
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
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【COLUMN_NAME: %s】【DATA_TYPE : %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.Field_Type)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.Field_Type)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【COLUMN_NAME: %s】【DATA_TYPE : %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.Field_Type)
		}
	} else {
		fmt.Println("————————————————————【没有大字段的表】———————————————————")

	}
	fmt.Printf("——————————————————↑大字段的表巡检完毕↑————————————————————\n\n")

}
func Table_long_varchar_Inspection(db *sql.DB) {
	var long_varchar int = 500
	sqlStr := "SELECT table_schema,table_name,column_name,CHARACTER_MAXIMUM_LENGTH FROM information_schema.COLUMNS WHERE DATA_TYPE = 'varchar' AND CHARACTER_MAXIMUM_LENGTH > ? AND table_schema NOT IN  ( 'information_schema', 'performance_schema', 'mysql', 'sys' );"
	counts, err := db.Query(sqlStr, long_varchar)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}

	defer counts.Close()
	count := 0
	if counts.Next() {
		var user User
		count += 1
		err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.INDEX_COUNT)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【COLUMN_NAME: %s】【CHARACTER_MAXIMUM_LENGTH : %d】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.INDEX_COUNT)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME, &user.CHARACTER_SET_NAME, &user.INDEX_COUNT)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】【COLUMN_NAME: %s】【CHARACTER_MAXIMUM_LENGTH : %d】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME, user.CHARACTER_SET_NAME, user.INDEX_COUNT)
		}
	} else {
		fmt.Println("————————————【没有varchar定义长度超过500的表】———————————")

	}
	fmt.Printf("————————————————↑varchar定义长度超过%d的表↑——————————————\n\n", long_varchar)

}

func Table_no_index_Inspection(db *sql.DB) {
	sqlStr := "SELECT t.table_schema,t.table_name FROM information_schema.TABLES AS t LEFT JOIN  ( SELECT DISTINCT table_schema, table_name FROM information_schema.`KEY_COLUMN_USAGE` ) AS kt ON  kt.table_schema = t.table_schema AND kt.table_name = t.table_name WHERE t.table_schema NOT IN  ( 'mysql', 'information_schema', 'performance_schema', 'sys' ) AND kt.table_name IS NULL;"
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
		fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME)
		for counts.Next() {
			var user User
			count += 1
			err := counts.Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("%d.【TABLE_SCHEMA: %s】【TABLE_NAME: %s】\n", count, user.TABLE_SCHEMA, user.TABLE_NAME)
		}
	} else {
		fmt.Println("———————————————————【没有无主键/索引的表】————————————————")

	}
	fmt.Printf("—————————————————↑无主键/索引的表巡检完毕↑————————————————\n\n")

}
