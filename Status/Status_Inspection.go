package Status

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

func Important_status_Inspection() {
	var status_list = [...]string{"Uptime", "Opened_files", "Opened_table_definitions", "Opened_tables", "Max_used_connections", "Threads_created", "Threads_connected", "Aborted_connects", "Aborted_clients", "Table_locks_waited", "Innodb_buffer_pool_wait_free", "Innodb_log_waits", "Innodb_row_lock_waits", "Innodb_row_lock_time_avg", "Binlog_cache_disk_use", "Created_tmp_disk_tables"}
	for j := 0; j < len(status_list); j++ {
		var user User
		values := status_list[j]
		sqlStr := "show global status like" + "'%" + values + "'"
		err := Global.DB.QueryRow(sqlStr).Scan(&user.TABLE_SCHEMA, &user.TABLE_NAME)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("【name: %s】【value: %s】\n", Green(user.TABLE_SCHEMA), Red(user.TABLE_NAME))
	}
	Completed.Printf("——————————————————↑重要状态巡检完毕↑—————————————————————")
	fmt.Println(" \n ")
}
