package Backup_module

import (
	"fmt"

	_ "github.com/gookit/color"
)

func Backup() {
	fmt.Print("还未开发完成,主要作为备份功能")

	// import (
	//     "database/sql"
	//     "fmt"
	//     "os"
	//     "os/exec"
	//     "time"

	//     _ "github.com/go-sql-driver/mysql"
	// )

	// func main() {
	//     // 数据库连接信息
	//     db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	//     if err != nil {
	//         fmt.Println("数据库连接失败:", err)
	//         return
	//     }
	//     defer db.Close()

	//     // 获取所有数据库名
	//     rows, err := db.Query("SHOW DATABASES")
	//     if err != nil {
	//         fmt.Println("查询数据库列表失败:", err)
	//         return
	//     }
	//     defer rows.Close()

	//     // 遍历所有数据库
	//     for rows.Next() {
	//         var dbName string
	//         err := rows.Scan(&dbName)
	//         if err != nil {
	//             fmt.Println("获取数据库名失败:", err)
	//             continue
	//         }

	//         // 忽略系统数据库
	//         if dbName == "information_schema" || dbName == "mysql" || dbName == "performance_schema" || dbName == "sys" {
	//             continue
	//         }

	//         // 备份当前数据库
	//         backup(dbName)
	//     }
	// }

	// // 备份指定数据库
	// func backup(dbName string) {
	//     // 备份文件名
	//     fileName := fmt.Sprintf("%s_%s.sql", dbName, time.Now().Format("20060102_150405"))

	//     // 备份命令
	//     cmd := exec.Command("mysqldump", "-uroot", "-ppassword", "--databases", dbName, "--result-file="+fileName)

	//     // 执行备份命令
	//     err := cmd.Run()
	//     if err != nil {
	//         fmt.Println("备份数据库失败:", err)
	//         return
	//     }

	//     fmt.Println("备份数据库成功:", dbName)
	// }
}
