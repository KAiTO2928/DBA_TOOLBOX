package Variables

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
	Red       = color.Red.Render
	Completed = color.S256(255, 27)
)

func Variables_Inspection() {
	var status_list = [...]string{"version", "innodb_buffer_pool_size", "innodb_flush_log_at_trx_commit",
		"innodb_log_file_size", "innodb_log_files_in_group", "innodb_file_per_table",
		"innodb_max_dirty_pages_pct", "sync_binlog", "max_connections",
		"table_open_cache", "table_definition_cache"}
	for j := 0; j < len(status_list); j++ {
		var user User
		values := status_list[j]
		sqlStr := "show global variables like" + "'" + values + "'"
		err := Global.DB.QueryRow(sqlStr).Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("【name: %s】【value: %s】\n", Green(user.TABLE_SCHEMA), Red(user.TABLE_NAME))
	}
	Completed.Printf("——————————————————↑重要参数巡检完毕↑—————————————————————")
	fmt.Println(" \n ")

}
