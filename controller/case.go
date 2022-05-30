package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se_case_back_end/common"
	"se_case_back_end/model"
	"se_case_back_end/response"
	"strconv"
)

func CommitCase(c *gin.Context) {
	id := c.Param("id")
	idD, _ := strconv.Atoi(id)
	db := common.GetDB()
	rg := model.Register{}
	if err := db.Model(&model.Register{}).Where("id = " + id).Take(&rg).Error; err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "该挂号不存在")
		return
	}
	if err := db.Model(&model.Case{}).Where("RegisterID = " + id).Take(nil).Error; err == nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "该挂号已处理")
		return
	}

	var cas model.Case
	err := c.BindJSON(&cas)
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "输入违法")
		return
	}
	cas.RegisterID = uint(idD)

	var sps model.Sps
	err = c.BindJSON(&sps)
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "辅助检查输入违法")
		return
	}

	var trs model.Trs
	err = c.BindJSON(&trs)
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "处理意见输入违法")
		return
	}

	db.Create(cas)

	for i := 0; i < len(sps.Sps); i++ {
		sps.Sps[i].ClinicID = uint(idD)
		db.Create(sps.Sps[i])
	}

	for i := 0; i < len(trs.Trs); i++ {
		trs.Trs[i].ClinicID = uint(idD)
		db.Create(trs.Trs[i])
	}

	response.Success(c, gin.H{
		"record_id": id,
	}, "修改成功")
}
