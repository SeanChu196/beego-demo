package main

import (
	"fmt"

	"./models"
)

func main() {
	pwd := models.GetPwd("48018122236")
	fmt.Println(pwd == "")
}
