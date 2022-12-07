package main

import (
	"fmt"
	_ "proj/routers"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("http://localhost:8080/")
	beego.Run()
}
