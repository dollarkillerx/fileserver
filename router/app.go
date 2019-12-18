/**
 * @Author: DollarKillerX
 * @Description: app.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:53 2019/12/18
 */
package router

import (
	"fileserver/server"
	"github.com/gin-gonic/gin"
)

func Registered(app *gin.Engine) {
	app.GET("/:dir/:file", server.DownloadFile) // 文件下载
	app.POST("/:dir/:file", server.UploadFile)  //文件上传
}
