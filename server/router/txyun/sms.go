package txyun

import (
	v1 "github.com/defeng-hub/ByOfficeAutomatic/server/api/v1"
	"github.com/defeng-hub/ByOfficeAutomatic/server/middleware"
	"github.com/gin-gonic/gin"
)

type TxSmsRouter struct {
}

func (e *TxSmsRouter) InitTxYunRouter(Router *gin.RouterGroup) {

	// 腾讯云的sms模块注册
	txSmsRouter := Router.Group("txyun/sms/").Use(middleware.OperationRecord())
	//txSmsRouterWithoutRecord := Router.Group("txyun/sms/")
	TxSmsApi := v1.ApiGroupApp.TxyunApiGroup.SmsHandler
	{
		txSmsRouter.POST("UpdateTemplates", TxSmsApi.UpdateTemplatesApi)
		txSmsRouter.GET("GetSmsList", TxSmsApi.GetSmsList)
		txSmsRouter.POST("SendSms", TxSmsApi.SendSms)
		txSmsRouter.POST("GetAllSmsProject", TxSmsApi.GetAllSmsProject)
		txSmsRouter.POST("AddSmsProject", TxSmsApi.AddSmsProject)
		txSmsRouter.POST("DelSmsProject", TxSmsApi.DelSmsProject)
		txSmsRouter.POST("SmsProjectRows", TxSmsApi.SmsProjectRows)
		txSmsRouter.POST("DelSmsProjectRow", TxSmsApi.DelSmsProjectRow)
		txSmsRouter.POST("DelSmsProjectRows", TxSmsApi.DelSmsProjectRows)
		txSmsRouter.POST("AddSmsProjectRow", TxSmsApi.AddSmsProjectRow)
		txSmsRouter.POST("SendSmsByRows", TxSmsApi.SendSmsByRows)
		txSmsRouter.POST("ExportExcelSmsRows", TxSmsApi.ExportExcelSmsRows)

	}

	//	oss模块
	//txOssRouter := Router.Group("txyun/oss/").Use(middleware.OperationRecord())
	//TxOssApi := v1.ApiGroupApp.TxyunApiGroup.OssHandle

	//	腾讯的其他模块注册
}
