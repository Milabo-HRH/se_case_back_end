package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se_case_back_end/common"
	"se_case_back_end/model"
	"se_case_back_end/response"
	"strconv"
	"time"
)

func CommitCase(c *gin.Context) {
	id := c.Query("id")
	idD, _ := strconv.Atoi(id)
	db := common.GetDB()
	rg := model.Register{}
	if err := db.Model(&model.Register{}).Where("id = " + id).Take(&rg).Error; err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "该挂号不存在")
		return
	}
	rg.Status = "T"
	rg.UpdatedAt = time.Now()

	ca := model.Case{}
	if err := db.Model(&model.Case{}).Where("register_id = " + id).Take(&ca).Error; err == nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "该挂号已处理")
		return
	}
	var com model.Com
	err := c.BindJSON(&com)
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "输入违法")
		return
	}
	cas := com.Cas
	sps := com.Sps
	trs := com.Trs
	cas.CreatedAt = time.Now()
	cas.UpdatedAt = time.Now()
	db.Save(rg)
	db.Create(&cas)
	for i := 0; i < len(sps); i++ {
		sps[i].ClinicID = uint(idD)
		sps[i].CreatedAt = time.Now()
		sps[i].UpdatedAt = time.Now()
		db.Create(&sps[i])
	}

	for i := 0; i < len(trs); i++ {
		trs[i].ClinicID = uint(idD)
		trs[i].CreatedAt = time.Now()
		trs[i].UpdatedAt = time.Now()
		db.Create(&trs[i])
	}

	response.Success(c, gin.H{
		"record_id": id,
	}, "修改成功")
}
