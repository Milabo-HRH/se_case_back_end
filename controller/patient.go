package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se_case_back_end/common"
	"se_case_back_end/model"
	"se_case_back_end/response"
	"strconv"
)

func GetRecordByPID(c *gin.Context) {
	id := c.Param("patient_id")
	db := common.GetDB()
	var records []model.Case
	if err := db.Model(&model.Case{}).Where("UserID = " + id).Order("id desc").Find(&records).Error; err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "查询失败")
		return
	}
	var res []gin.H
	for i := 0; i < len(records); i++ {
		res = append(res, gin.H{
			"date": records[i].CreatedAt.Format("2006-01-02 15:04:05"),
			"url":  "/api/view/" + strconv.Itoa(int(records[i].ID)),
		})
	}
	response.Success(c, gin.H{"data": res}, "用户所有挂号")
}

func GetRecord(c *gin.Context) {
	id := c.Param("id")
	db := common.GetDB()
	rg := model.Register{}
	if err := db.Model(&model.Register{}).Where("id = " + id).Take(&rg).Error; err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	/* todo */
	if rg.Department == "" {
		response.Response(c, http.StatusBadRequest, 400, nil, "该对象不存在")
	}
	//var cas gin.H
	//var reg gin.H
	var sup []gin.H
	var tre []gin.H
	reg := gin.H{
		"id":         rg.ID,
		"name":       rg.Name,
		"age":        rg.Age,
		"gender":     rg.Gender,
		"department": rg.Department,
		"status":     rg.Status,
		"regTime":    rg.CreatedAt.Format("2006-01-02 15:04:05"),
		"userID":     rg.UserID,
	}
	ca := model.Case{}
	if err := db.Model(&model.Case{}).Where("RegisterID = " + id).Take(&ca).Error; err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	cas := gin.H{
		"cc":   ca.CC,
		"hopi": ca.HOPI,
		"pmh":  ca.PMH,
		"pe":   ca.PE,
		"pd":   ca.PD,
		"pc":   ca.RC,
		"edu":  ca.EDU,
	}
	var sp []model.Supplement
	if err := db.Model(&model.Supplement{}).Where("ClinicID = " + id).Order("id desc").Find(&sp).Error; err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "查询失败")
		return
	}
	for i := 0; i < len(sp); i++ {
		sup = append(sup, gin.H{
			"checkName": sp[i].CheckName,
			"result":    sp[i].Result,
		})
	}
	var tm []model.Treatment
	if err := db.Model(&model.Treatment{}).Where("ClinicID = " + id).Order("id desc").Find(&tm).Error; err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "查询失败")
		return
	}
	for i := 0; i < len(tm); i++ {
		tre = append(tre, gin.H{
			"medName": tm[i].MedName,
			"val":     tm[i].Val,
			"unit":    tm[i].Unit,
			"usage":   tm[i].Usage,
		})
	}
	response.Success(c, gin.H{"reg": reg, "cas": cas, "sup": sup, "tre": tre}, "单次病历")

}
