package Menu

import (
	"DB_OSInspection/Index"
	"DB_OSInspection/Inspection"
	"DB_OSInspection/Monitor"
	"DB_OSInspection/Privileges"
	"DB_OSInspection/Status"
	"DB_OSInspection/Table"
	"DB_OSInspection/User"
	"DB_OSInspection/Variables"
	"fmt"
)

func Menu(m string) {
	//选择模式：all全部巡检、table只巡检表、index只巡检索引、variables只巡检重要参数、status只巡检重要状态、user只巡检用户、privileges只巡检权限
	switch m {
	case "all":
		//表巡检
		Table.Table_Data_Size_Inspection()
		Table.Table_Index_Size_Inspection()
		Table.Table_fragment_Size_Inspection()
		Table.Table_rows_Size_Inspection()
		Table.Table_chaset_Inspection()
		Table.Table_big_columns_Inspection()
		Table.Table_long_varchar_Inspection()
		Table.Table_no_index_Inspection()
		//索引巡检
		Index.Index_redundant_Inspection()
		Index.Index_columns_Inspection()
		Index.Index_unused_Inspection()
		//参数巡检
		Variables.Variables_Inspection()
		//状态巡检
		Status.Important_status_Inspection()
		//用户巡检
		User.User_nopass_Inspection()
		//权限巡检
		Privileges.Privileges_Inspection()
		Privileges.User_Privileges_Inspection()
		Inspection.Inspection()
	case "table":
		Table.Table_Data_Size_Inspection()
		Table.Table_Index_Size_Inspection()
		Table.Table_fragment_Size_Inspection()
		Table.Table_rows_Size_Inspection()
		Table.Table_chaset_Inspection()
		Table.Table_big_columns_Inspection()
		Table.Table_long_varchar_Inspection()
		Table.Table_no_index_Inspection()
	case "index":
		Index.Index_redundant_Inspection()
		Index.Index_columns_Inspection()
		Index.Index_unused_Inspection()
	case "variables":
		Variables.Variables_Inspection()
	case "status":
		Status.Important_status_Inspection()
	case "user":
		User.User_nopass_Inspection()
	case "privileges":
		Privileges.Privileges_Inspection()
		Privileges.User_Privileges_Inspection()
	case "monitor":
		Monitor.DB_monitor()

	default:
		fmt.Println("你未选择模式")
	}
}
