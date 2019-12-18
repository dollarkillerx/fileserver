/**
 * @Author: DollarKillerX
 * @Description: 消息统一返回
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:03 2019/12/18
 */
package def

import "github.com/gin-gonic/gin"

type Message struct {
	Code    int
	Message string
}

var OK200 = Message{Code: 200, Message: "OK"}
var Err500 = Message{Code: 500, Message: "500 Server Error"}
var No404 = Message{Code: 404, Message: "No file"}

func Return(ctx *gin.Context, message Message) {
	ctx.JSON(message.Code, message)
}
