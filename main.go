/**
 * @Author: DollarKillerX
 * @Description: main.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:49 2019/12/18
 */
package main

import (
	"fileserver/router"
	"fileserver/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatalln("请输入运行地址 比如 ./main 0.0.0.0:8081")
	}

	// 如果用户指定了路径
	if len(os.Args) == 3 {
		server.FileDir = os.Args[2]
	}

	app := gin.New() // 最小核心启动

	app.MaxMultipartMemory = 1024 << 20 // 1024M

	app.Use(gin.Recovery()) // 注册recovery中间件防止panic

	router.Registered(app)
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("程序运行在: ", os.Args[1])
	err := app.Run(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
}
