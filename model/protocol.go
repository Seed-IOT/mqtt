package model

import (
	"mqtt/constant"
	"mqtt/log"
	"strings"

	"encoding/json"

	"github.com/sirupsen/logrus"
)

// ProtocolForJSON 协议表
type ProtocolForJSON struct {
	ID              uint
	ProductKey      string
	Data            ResContainer
	ProtocolVersion string
}

// Protocol 协议表
type Protocol struct {
	ID              uint `grom:"primary_key"`
	ProductKey      string
	ProtocolVersion string
	Data            string `gorm:"type:text"`
}

// ResContainer 协议json 第一层结构
type ResContainer struct {
	Name            string
	ProtocolVersion string
	ProductKey      string
	Description     string
	DataPoint       []ResObject
}

// ResObject 数据点对象
type ResObject struct {
	Name        string
	DisplayName string
	Description string
	Type        string // string bool uint8 uint16 uint32 enum
	Operations  string // R RW
	Len         uint   // 单位 bit 数值型的Len要结合Step来算
	Step        float32
	Max         int
	Min         int
}

// CreateProtocol 创建协议
func (srv *Service) CreateProtocol(productKey string, data ResContainer) (constant.BaseReturn, *constant.Error) {
	baseReturn := constant.BaseReturn{}
	errorData := new(constant.Error)
	errorData.Code = constant.SUCCESS

	// TODO校验

	protocol := Protocol{}

	//TODO 先查询是否存在 结合版本号
	srv.DB.First(&protocol, "product_key = ?", productKey)

	if strings.EqualFold(protocol.ProductKey, productKey) {
		// 提示存在
		log.GlobalLog.WithFields(logrus.Fields{
			"error": nil,
		}).Info("productKey existing")
		errorData.Code = constant.UNKNOWN_ERROR
		return baseReturn, errorData
	}

	//TODO 插入并排序 布尔 枚举 数值 string 并且可写在前

	protocol.ProductKey = productKey
	dataJSON, nil := json.Marshal(data)
	protocol.Data = string(dataJSON)

	if err := srv.DB.Create(&protocol).Error; err != nil {
		errorData.Code = constant.UNKNOWN_ERROR
		log.GlobalLog.WithFields(logrus.Fields{
			"error": err,
		}).Info("Create Protocol Error")
		return baseReturn, errorData

	}
	return baseReturn, errorData
}
