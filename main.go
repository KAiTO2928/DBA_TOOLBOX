package main

import (
	"DB_OSInspection/Global"
	"DB_OSInspection/Menu"
	"database/sql"
	"flag"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func banner() {
	bannertext :=
		`
██████╗ ██████╗      ██████╗ ███████╗██╗███╗   ██╗███████╗██████╗ ███████╗ ██████╗████████╗██╗ ██████╗ ███╗   ██╗
██╔══██╗██╔══██╗    ██╔═══██╗██╔════╝██║████╗  ██║██╔════╝██╔══██╗██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║
██║  ██║██████╔╝    ██║   ██║███████╗██║██╔██╗ ██║███████╗██████╔╝█████╗  ██║        ██║   ██║██║   ██║██╔██╗ ██║
██║  ██║██╔══██╗    ██║   ██║╚════██║██║██║╚██╗██║╚════██║██╔═══╝ ██╔══╝  ██║        ██║   ██║██║   ██║██║╚██╗██║
██████╔╝██████╔╝    ╚██████╔╝███████║██║██║ ╚████║███████║██║     ███████╗╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║
╚═════╝ ╚═════╝      ╚═════╝ ╚══════╝╚═╝╚═╝  ╚═══╝╚══════╝╚═╝     ╚══════╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
██╗   ██╗ ██╗    ██████╗     ██████╗ ██╗   ██╗    ██╗  ██╗██╗   ██╗       ██╗       ███████╗     ██╗             
██║   ██║███║   ██╔═████╗    ██╔══██╗╚██╗ ██╔╝    ██║  ██║╚██╗ ██╔╝       ██║       ██╔════╝     ██║             
██║   ██║╚██║   ██║██╔██║    ██████╔╝ ╚████╔╝     ███████║ ╚████╔╝     ████████╗    ███████╗     ██║             
╚██╗ ██╔╝ ██║   ████╔╝██║    ██╔══██╗  ╚██╔╝      ██╔══██║  ╚██╔╝      ██╔═██╔═╝    ╚════██║██   ██║             
 ╚████╔╝  ██║██╗╚██████╔╝    ██████╔╝   ██║       ██║  ██║   ██║       ██████║      ███████║╚█████╔╝             
  ╚═══╝   ╚═╝╚═╝ ╚═════╝     ╚═════╝    ╚═╝       ╚═╝  ╚═╝   ╚═╝       ╚═════╝      ╚══════╝ ╚════╝              
                                                                                                                 
		`
	fmt.Println(bannertext)
}

func getErrorMessage(err error) {
	fmt.Println("Connection failed with the following error:")
	if strings.Contains(err.Error(), "using password: YES") {
		// 密码或用户名错误
		fmt.Println("Either the password or username is incorrect!")

	} else if strings.Contains(err.Error(), "Unknown database") {
		//数据库名错误
		fmt.Println("The database name was entered incorrectly!")

	} else if strings.Contains(err.Error(), "no such host") {
		//数据库地址错误
		fmt.Println("The database address is incorrect!")

	} else if strings.Contains(err.Error(), "dial tcp") {
		//端口错误
		fmt.Println("The port is incorrect!")

	} else {
		fmt.Println("other errors!")
		panic(err)
	}

}

func main() {
	//获取用户名
	u := flag.String("u", "root", "input username")
	//获取密码
	p := flag.String("p", "", "input password")
	//获取链接模式
	nw := flag.String("nw", "tcp", "input netWork")
	//获取地址和端口号
	P := flag.String("P", "localhost:3306", "input port")
	/*选择模式：all全部巡检、table只巡检表、index只巡检索引、variables只巡检重要参数、
	status只巡检重要状态、user只巡检用户、privileges只巡检权限、monitor新增功能*/
	m := flag.String("m", "all", "input model")
	//版本号
	v := flag.Bool("v", false, "input port")
	//帮助
	h := flag.Bool("h", false, "show help")

	flag.Parse()
	if *v {
		fmt.Println("DB OSInspection V2.0")
		return
	}

	if *h || *u == "" {
		fmt.Println("Usage: DB OSInspection [options]")
		flag.PrintDefaults()
		return
	}

	//调用banner
	banner()
	// 32 设置绿色，使用 33 设置黄色，使用 34 设置蓝色
	fmt.Printf("\033[31m%s\033[0m", "欢迎使用数据库巡检工具 DB OSInspection V2.0\n")
	fmt.Printf("————————————————————————————————————————————————————————\n")

	conn := fmt.Sprintf("%s:%s@%s(%s)/%s", *u, *p, *nw, *P, "mysql")
	db, err := sql.Open("mysql", conn)
	Global.DB = db
	if err != nil {
		getErrorMessage(err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		getErrorMessage(err)
		return
	}
	//设置最大连接数
	db.SetMaxOpenConns(10)
	//最大空闲连接数
	db.SetMaxIdleConns(10)
	//跳转菜单页面
	Menu.Menu(*m)

	fmt.Println("感谢您使用数据库巡检工具 DB OSInspection")

}
