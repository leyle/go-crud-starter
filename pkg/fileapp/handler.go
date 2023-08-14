package fileapp

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/leyle/go-crud-starter/internal/ginsetup"
	"github.com/leyle/go-crud-starter/pkg/errcode"
	"github.com/leyle/go-crud-starter/utils"
	"net/http"
	"strconv"
	"strings"
)

func UploadFileHandler(c *gin.Context) {
	ctx := configandcontext.GetAPIContextFromGinContext(c)
	t0 := utils.PrintFuncStartLog(ctx)
	defer utils.PrintFuncEndLog(ctx, t0)

	file, header, err := c.Request.FormFile("file")
	ginsetup.StopExec400(err, errcode.ErrInvalidClientArgument, "")

	upFile, err := saveFileToGridFS(ctx, file, header)
	ginsetup.StopExec400(err, errcode.ErrWriteGridFSFailed, "")

	metaInfo := upFile.Metadata

	ginsetup.ReturnOKJson(c, metaInfo)
	return
}

func DownloadFileByIdHandler(c *gin.Context) {
	ctx := configandcontext.GetAPIContextFromGinContext(c)
	t0 := utils.PrintFuncStartLog(ctx)
	defer utils.PrintFuncEndLog(ctx, t0)

	fileId := c.Query("id")
	if fileId == "" {
		ctx.Logger.Error().Msg("no id in url query string")
		ginsetup.Return400Json(c, errcode.ErrInvalidClientArgument, "no id in url query string")
		return
	}

	upFile, err := getFileFromGridFSById(ctx, fileId)
	ginsetup.StopExec400(err, errcode.ErrGetGridFSFailed, "")

	// if debug=yes, then only return metadata as json string
	debug := c.Query("debug")
	if strings.ToLower(debug) == "yes" {
		ginsetup.ReturnOKJson(c, upFile.Metadata)
		return
	}

	// normal file returns as stream, image will be rendered by browser
	fileSize := strconv.FormatInt(upFile.Metadata.Size, 10)
	c.Header("Content-Length", fileSize)

	dataBuf := upFile.Data[:512]
	cType := ginsetup.GuessFileContentType(dataBuf)
	defaultCT := "application/octet-stream"
	if cType == defaultCT {
		c.Header("Content-Disposition", "attachment; filename="+upFile.Metadata.OriginalFilename)
	}
	c.Data(http.StatusOK, cType, upFile.Data)
	return
}
