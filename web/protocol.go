package web

import (
	"mqtt/constant"
	"mqtt/log"
	"mqtt/model"

	"github.com/gin-gonic/gin"
)

func (srv *server) CreateProtocol(c *gin.Context) {
	returnJSON := constant.BaseReturn{}
	httpError := new(constant.Error)
	httpError.Code = constant.SUCCESS
	// 获取参数
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	param := model.ProtocolForJSON{}
	c.Bind(&param)

	log.GlobalLog.Error("name", param.Data.Name)
	productKey := param.ProductKey
	data := param.Data

	// 插入数据库
	_, err := srv.service.CreateProtocol(productKey, data)
	if err.Code != 0 {
		// 返回错误
		httpError.Code = err.Code
	}

	httpErrorData, _ := httpError.UnmarshalJSON()
	returnJSON.Code = httpErrorData.Code
	returnJSON.Message = httpErrorData.Message
	c.JSON(200, returnJSON)
}

func (srv *server) UpdateProtocol(c *gin.Context) {

}

func (srv *server) DeleteProtocol(c *gin.Context) {

}

func (srv *server) QueryProtocol(c *gin.Context) {

}

func (srv *server) DownloadProtocol(c *gin.Context) {

}
