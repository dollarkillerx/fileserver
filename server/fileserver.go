/**
 * @Author: DollarKillerX
 * @Description: fileserver.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:01 2019/12/18
 */
package server

import (
	"fileserver/def"
	"fileserver/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func UploadFile(ctx *gin.Context) {
	dir := ctx.Param("dir")
	file := ctx.Param("file")
	err := utils.DirPing(dir)
	if err != nil {
		def.Return(ctx, def.Err500)
		return
	}
	defer ctx.Request.Body.Close()
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		def.Return(ctx, def.Err500)
		return
	}
	filename := dir + "/" + file
	err = ioutil.WriteFile(filename, bytes, 00666)
	if err != nil {
		def.Return(ctx, def.Err500)
		return
	}
	def.Return(ctx, def.OK200)
}

func DownloadFile(ctx *gin.Context) {
	dir := ctx.Param("dir")
	file := ctx.Param("file")
	filename := dir + "/" + file
	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		def.Return(ctx, def.No404)
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+file)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Writer.Write(bytes)
}
