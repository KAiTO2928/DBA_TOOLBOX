package Variables

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

func Variables_Inspection(db *sql.DB) {
	var status_list = [...]string{"version", "innodb_buffer_pool_size", "innodb_flush_log_at_trx_commit",
		"innodb_log_file_size", "innodb_log_files_in_group", "innodb_file_per_table",
		"innodb_max_dirty_pages_pct", "sync_binlog", "max_connections",
		"table_open_cache", "table_definition_cache"}
	for j := 0; j < len(status_list); j++ {
		var user User
		values := status_list[j]
		sqlStr := "show global variables like" + "'" + values + "'"
		err := db.QueryRow(sqlStr).Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("【name: %s】【value: %s】\n", user.TABLE_SCHEMA, user.TABLE_NAME)
	}
	fmt.Printf("——————————————————↑重要参数巡检完毕↑—————————————————————\n\n")

}
