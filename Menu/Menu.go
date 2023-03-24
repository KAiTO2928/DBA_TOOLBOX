package Menu

import (
	"DB_OSInspection/Index"
	"DB_OSInspection/Privileges"
	"DB_OSInspection/Status"
	"DB_OSInspection/Table"
	"DB_OSInspection/User"
	"DB_OSInspection/Variables"
	"database/sql"
	"fmt"
)

func Menu(db *sql.DB, m string) {
	//选择模式：all全部巡检、table只巡检表、index只巡检索引、variables只巡检重要参数、status只巡检重要状态、user只巡检用户、privileges只巡检权限
	switch m {
	case "all":
		//表巡检
		Table.Table_Data_Size_Inspection(db)
		Table.Table_Index_Size_Inspection(db)
		Table.Table_fragment_Size_Inspection(db)
		Table.Table_rows_Size_Inspection(db)
		Table.Table_chaset_Inspection(db)
		Table.Table_big_columns_Inspection(db)
		Table.Table_long_varchar_Inspection(db)
		Table.Table_no_index_Inspection(db)
		//索引巡检
		Index.Index_redundant_Inspection(db)
		Index.Index_columns_Inspection(db)
		Index.Index_unused_Inspection(db)
		//参数巡检
		Variables.Variables_Inspection(db)
		//状态巡检
		Status.Important_status_Inspection(db)
		//用户巡检
		User.User_nopass_Inspection(db)
		//权限巡检
		Privileges.Privileges_Inspection(db)
		Privileges.User_Privileges_Inspection(db)
	case "table":
		Table.Table_Data_Size_Inspection(db)
		Table.Table_Index_Size_Inspection(db)
		Table.Table_fragment_Size_Inspection(db)
		Table.Table_rows_Size_Inspection(db)
		Table.Table_chaset_Inspection(db)
		Table.Table_big_columns_Inspection(db)
		Table.Table_long_varchar_Inspection(db)
		Table.Table_no_index_Inspection(db)
	case "index":
		Index.Index_redundant_Inspection(db)
		Index.Index_columns_Inspection(db)
		Index.Index_unused_Inspection(db)
	case "variables":
		Variables.Variables_Inspection(db)
	case "status":
		Status.Important_status_Inspection(db)
	case "user":
		User.User_nopass_Inspection(db)
	case "privileges":
		Privileges.Privileges_Inspection(db)
		Privileges.User_Privileges_Inspection(db)
	default:
		fmt.Println("你未选择模式")
	}
}
