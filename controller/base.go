package controller

import (
	"dc-kanban/config"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zctod/tool/common/utils"
	"net/http"
	"strconv"
)

type Base struct{
	Ctx *gin.Context
}

func (t *Base) Api(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func (t *Base) Succ(c *gin.Context, msg string) {
	t.Api(c, http.StatusOK, gin.H{
		"msg": msg,
	})
}

func (t *Base) Data(c *gin.Context, data interface{}) {
	t.Api(c, http.StatusOK, data)
}

func (t *Base) Array(c *gin.Context, data interface{}) {

	data, err := utils.ArrayStructToMap(data)
	if err != nil {
		t.Err(c, err.Error())
		return
	}
	t.Data(c, data)
}

func (t *Base) Err(c *gin.Context, msg string) {

	fmt.Println("[Error]", msg)
	t.Api(c, http.StatusBadRequest, gin.H{
		"msg": msg,
	})
}

func (t *Base) Paginate(c *gin.Context, db *gorm.DB) error {

	var page, pageSize string
	switch c.Request.Method {
	case "GET":
		page = c.DefaultQuery("page", "1")
		pageSize = c.Query("pageSize")
		break
	case "POST":
		break
	default:
		return errors.New("解析失败")
	}

	var pageN, pageSizeN int
	var err error
	pageN, err = strconv.Atoi(page)
	if err != nil {
		return errors.New("请填写page为数值类型")
	}
	if pageN == 0 {
		pageN = 1
	}
	if pageSize == "" {
		pageSizeN = config.PAGINATE_PAGESIZE
	} else {
		pageSizeN, err = strconv.Atoi(pageSize)
		if err != nil {
			return errors.New("请填写pageSize为数值类型")
		}
		if pageSizeN == 0 {
			pageSizeN = config.PAGINATE_PAGESIZE
		}
	}

	db.Limit(pageSizeN).Offset((pageN - 1) * pageSizeN)
	return nil
}
