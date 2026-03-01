package service

import (
	"log"
	"net/http"
	"zuoye/bff/dasic/config"
	"zuoye/bff/handler/request"
	__ "zuoye/srv/dasic/proto"

	"github.com/gin-gonic/gin"
)

func OrderAdd(c *gin.Context) {
	var form request.OrderAdd
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误",
			"code": 400,
		})
		return
	}
	r, err := config.UserClient.OrderAdd(c, &__.OrderAddReq{
		Name:  form.Name,
		Price: uint32(form.Price),
		Num:   int64(form.Num),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  r.Msg,
		"code": r.Code,
	})
	return
}
