package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se_case_back_end/common"
	"se_case_back_end/model"
	"se_case_back_end/response"
	"strconv"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	name := c.PostForm("name")
	age, err := strconv.ParseUint(c.PostForm("age"), 10, 0)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "年龄应为整数")
		return
	}
	gender := c.PostForm("gender")
	department := c.PostForm("department")
	userID, err := strconv.ParseUint(c.PostForm("userID"), 10, 0)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户ID应为整数")
		return
	}
	doctorID, err := strconv.ParseUint(c.PostForm("doctorID"), 10, 0)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "医生ID应为整数")
		return
	}
	doctorName := c.PostForm("doctorName")

	newReg := model.Register{
		Name:       name,
		Age:        uint(age),
		Gender:     gender,
		Department: department,
		Status:     "F",
		UserID:     uint(userID),
		DoctorID:   uint(doctorID),
		DoctorName: doctorName,
	}

	db.Create(&newReg)
	response.Success(c, gin.H{"id": newReg.ID}, "新建病历成功")
}
