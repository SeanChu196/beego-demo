package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
)

//func init() {
//	// PostgreSQL 配置
//	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
//	orm.RegisterDataBase("default", "postgres", "user=postgres password=asdf dbname=testUsr host=127.0.0.1 port=5432 sslmode=disable")
//	// 自动建表
//	orm.RunSyncdb("default", false, true)
//
//}
func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text(), " : ", models.IsValidEmail(scanner.Text()))
	//	fmt.Println(scanner.Text(), " : ", models.IsValidPhone(scanner.Text()))
	//	fmt.Println(scanner.Text(), " : ", models.IsValidPwd(scanner.Text()))
	//}
	//	usr := "13776216006"
	//	pwd := "seabZaii88"
	//	id, err := models.InsertUser(usr, pwd)
	//	if err != nil {
	//		fmt.Println("Faile to create usr ", usr)
	//	} else {
	//		fmt.Println("Id:", id, "\tusr:", usr)
	//	}
	//	fmt.Println("=====================================================")
	//	user := models.User{Usr: "15151768576"}
	//	er := models.ReadUser(&user, "Usr")
	//	if er != nil {
	//		fmt.Println("Faile to query usr ", user.Usr)
	//	} else {
	//		fmt.Println(user)
	//	}
	beego.Run()
}
