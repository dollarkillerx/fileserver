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
	"io"
	"log"
	"os"
	"path/filepath"
)

var FileDir = ""

func UploadFile(ctx *gin.Context) {
	dir := ctx.Param("dir")
	file := ctx.Param("file")
	dirpath := filepath.Join(FileDir, dir)
	err := utils.DirPing(dirpath)
	if err != nil {
		log.Println(err)
		def.Return(ctx, def.Err500)
		return
	}
	defer ctx.Request.Body.Close()
	filename := filepath.Join(FileDir, dir, file)
	open, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		def.Return(ctx, def.Err500)
		return
	}
	_, err = io.Copy(open, ctx.Request.Body)
	if err != nil {
		log.Println(err)
		def.Return(ctx, def.Err500)
		return
	}
	def.Return(ctx, def.OK200)
}

func DownloadFile(ctx *gin.Context) {
	dir := ctx.Param("dir")
	file := ctx.Param("file")
	filename := filepath.Join(FileDir, dir, file)
	open, e := os.Open(filename)
	if e != nil {
		log.Println(e)
		def.Return(ctx, def.No404)
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+file)
	ctx.Header("Content-Transfer-Encoding", "binary")
	_, e = io.Copy(ctx.Writer, open)
	if e != nil {
		log.Println(e)
		def.Return(ctx, def.Err500)
	}
}
