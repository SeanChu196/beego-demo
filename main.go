package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
)

func init() {

}
func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text(), " : ", models.IsValidEmail(scanner.Text()))
	//	fmt.Println(scanner.Text(), " : ", models.IsValidPhone(scanner.Text()))
	//	fmt.Println(scanner.Text(), " : ", models.IsValidPwd(scanner.Text()))
	//}
	beego.Run()

}
