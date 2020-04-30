package handle

import (
	"era-shop.xyz/era-shop/pkg/erafile"
	"era-shop.xyz/era-shop/pkg/response"
	"github.com/gin-gonic/gin"
)

type Common struct {}

func (ch *Common) Upload(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	err := erafile.SizeLimit(fileHeader, "image")
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	filePath := erafile.FullPath(fileHeader)
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		response.Failed(c, "上传失败！")
		return
	}
	response.Success(c, "/"+filePath)
}