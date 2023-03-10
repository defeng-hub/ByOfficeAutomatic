package sms

import (
	"fmt"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	smsmodel "github.com/defeng-hub/ByOfficeAutomatic/server/model/txyun/sms"
)

func (e *TencentSmsService) AllSmsProject() (list interface{}, total int64, err error) {
	var Projects []smsmodel.SmsProject
	db := global.GVA_DB.Model(smsmodel.SmsProject{})

	err = db.Count(&total).Error
	if err != nil {
		return Projects, 0, err
	} else {
		err = db.Preload("SmsTemplate").Find(&Projects).Error
		if err != nil {
			return Projects, 0, err
		}
	}
	return Projects, total, err
}

func (e *TencentSmsService) AddSmsProject(req *smsmodel.AddSmsProjectReq) (*smsmodel.SmsProject, error) {
	db := global.GVA_DB.Model(smsmodel.SmsTemplate{})
	template := smsmodel.SmsTemplate{}
	tx := db.Find(&template, "template_id = ?", req.TemplateId)
	if tx.RowsAffected != 1 {
		return nil, fmt.Errorf("短信模板不存在")
	}

	//	找到了这个模板, 那么就制造smsProject
	db = global.GVA_DB.Model(smsmodel.SmsProject{})
	project := smsmodel.SmsProject{
		ProjectName: req.ProjectName,
		Comment:     req.Comment,
		TemplateId:  template.TemplateId,
		RunNum:      0,
		RowNum:      0,
		ParamNum:    template.GetParamsNum(),
	}
	err := db.Create(&project).Error

	if err != nil {
		return nil, err
	} else {
		return &project, nil
	}
}

func (e *TencentSmsService) DelSmsProject(req *request.GetById) (*smsmodel.SmsProject, error) {
	db := global.GVA_DB.Model(smsmodel.SmsProject{})
	project := smsmodel.SmsProject{}
	tx := db.Where("id = ?", req.ID).Find(&project)
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("模板ID不存在")
	}

	tx = db.Where("id = ?", req.ID).Delete(&smsmodel.SmsProject{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &project, nil
}

// SmsProjectRows 下边是project row 的service
// SmsProjectRows 通过项目ID拿到项目的全部行数据
func (e *TencentSmsService) SmsProjectRows(info *smsmodel.SmsProjectRowsPageReq) (
	list interface{}, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(smsmodel.SmsProjectRow{})

	var List []smsmodel.SmsProjectRow
	err = db.Where("sms_project_id = ?", info.ID).Count(&total).Error
	if err != nil {
		return List, total, err
	} else {
		//err = db.Where("sms_project_id = ?", info.ID).Limit(limit).Offset(offset).Find(&List).Error
		err = db.Where("sms_project_id = ?", info.ID).Where("phone LIKE ?", "%"+info.Keyword+"%").Limit(limit).Offset(offset).Find(&List).Error
	}
	return List, total, err
}

func (e *TencentSmsService) DelSmsProjectRow(req *request.GetById) (
	*smsmodel.SmsProjectRow, error) {
	row := smsmodel.SmsProjectRow{}
	row.ID = uint(req.ID)

	db := global.GVA_DB.Model(smsmodel.SmsProjectRow{})
	tx := db.Unscoped().Delete(&row)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, fmt.Errorf("删除失败")
	}
	return &row, nil
}

func (e *TencentSmsService) DelSmsProjectRows(req *request.IdsReq) error {

	db := global.GVA_DB.Model([]*smsmodel.SmsProjectRow{})
	tx := db.Unscoped().Delete([]*smsmodel.SmsProjectRow{}, req.Ids)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return fmt.Errorf("删除失败")
	} else {
		return nil
	}
}

func (e *TencentSmsService) AddSmsProjectRow(req []*smsmodel.SmsProjectRow) (
	*smsmodel.SmsProjectRow, error) {

	db := global.GVA_DB.Model(smsmodel.SmsProjectRow{})
	tx := db.Create(&req)
	if tx.Error != nil && tx.RowsAffected != 0 {
		return nil, tx.Error
	}
	return nil, nil
}

func (e *TencentSmsService) UpdateSmsProjectRow(req *smsmodel.SmsProjectRow) (
	*smsmodel.SmsProjectRow, error) {
	return nil, nil
}

func (e *TencentSmsService) SendSmsByRows(req *request.IdsReq) error {
	var rows []*smsmodel.SmsProjectRow
	//	传入ids,project_id
	db := global.GVA_DB.Model(rows)

	tx := db.Preload("SmsProject").Find(&rows, req.Ids)
	if tx.Error != nil {
		return tx.Error
	}
	for _, row := range rows {
		var params []string
		if row.Param1 != "" {
			params = append(params, row.Param1)
		}
		if row.Param2 != "" {
			params = append(params, row.Param2)
		}
		if row.Param3 != "" {
			params = append(params, row.Param3)
		}
		if row.Param4 != "" {
			params = append(params, row.Param4)
		}
		if row.Param5 != "" {
			params = append(params, row.Param5)
		}
		if row.Param6 != "" {
			params = append(params, row.Param6)
		}
		if row.Param7 != "" {
			params = append(params, row.Param7)
		}
		if row.Param8 != "" {
			params = append(params, row.Param8)
		}
		if row.Param9 != "" {
			params = append(params, row.Param9)
		}
		err := e.SendSms(int(row.SmsProject.TemplateId), []string{row.Phone}, params)
		if err != nil {
			global.GVA_LOG.Error("批量发送短信err:" + err.Error())
		}
	}
	return nil
}
