package Backup_module

import (
	"dba_toolbox/Global"
	"fmt"
	"os/exec"
	"time"

	"github.com/gookit/color"
)

var (
	mode int
	// color
	Green     = color.Green.Render
	Red       = color.Red.Render
	Completed = color.S256(255, 27)
)

func Backup() {
	fmt.Println("欢迎使用DBA_TOOLBOX的备份功能")
	fmt.Print("请输入备份模式:[1,全库全表备份]、[2,除去系统表备份]、[3,指定库备份]>")
	fmt.Scanln(&mode)

	switch mode {
	case 1:
		fileName := fmt.Sprintf("All_DB_%s.sql", time.Now().Format("20060102_150405"))
		// 备份命令
		cmd := exec.Command("mysqldump", "-u"+Global.User, "-p"+Global.Password, "-A", "--result-file="+fileName)
		// 执行备份命令
		err := cmd.Run()
		if err != nil {
			fmt.Println("备份数据库失败:", err)
			return
		}
		fmt.Println("备份全库全表成功:")
	case 2:
		rows, err := Global.DB.Query("SHOW DATABASES")
		if err != nil {
			fmt.Println("查询数据库列表失败:", err)
			return
		}
		defer rows.Close()

		// 遍历所有数据库
		for rows.Next() {
			var dbName string
			err := rows.Scan(&dbName)
			if err != nil {
				fmt.Println("获取数据库名失败:", err)
				continue
			}

			// 忽略系统数据库
			if dbName == "information_schema" || dbName == "mysql" || dbName == "performance_schema" || dbName == "sys" {
				continue
			}

			// fmt.Println(dbName)
			// 备份当前数据库
			go backup(dbName)
			time.Sleep(time.Second)
		}
	case 3:
		for {
			var Appoint_dbName string
			fmt.Print("请输入要备份的库名,输入结束请输入 0 \n")
			fmt.Scanln(&Appoint_dbName)
			if Appoint_dbName == "0" {
				break
			}
			go backup(Appoint_dbName)

		}
	default:
		Completed.Printf("请重新输入")
		fmt.Println(" \n ")
		Backup()
	}
}

// 备份指定数据库
func backup(dbName string) {
	// 备份文件名
	fileName := fmt.Sprintf("%s_%s.sql", dbName, time.Now().Format("20060102_150405"))

	// 备份命令
	cmd := exec.Command("mysqldump", "-u"+Global.User, "-p"+Global.Password, "--databases", dbName, "--result-file="+fileName)
	// fmt.Println(cmd)
	// 执行备份命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("备份数据库失败:", err)
		return
	}

	fmt.Println("备份数据库成功:", dbName)
}
