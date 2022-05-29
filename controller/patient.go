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
}

func GetRecord(c *gin.Context) {
	id := c.Param("id")
	db := common.GetDB()
	cas := model.Register{}
	if err := db.Model(&model.Register{}).Where("id = " + id).Take(&cas).Error; err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "查询失败")
		return
	}
	/* todo */
}
