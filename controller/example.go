package controller

import (
	"aichat/model/vo"

	"github.com/gin-gonic/gin"
)

// ExampleHandler 用户登录
// @Summary      获取用户信息Summary
// @Description  获取用户信息Description
// @Tags         获取用户信息Tags
// @Accept       json
// @Produce      json
// @Param        user body vo.ExampleReq  true  "请求 body 中的数据"
// @Success      200  {object}  ResponseData{Data=vo.ExampleResp}
// @Router       /example [post]
func ExampleHandler(c *gin.Context) {
	var req vo.ExampleReq
	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		// response 报文
		ResponseError(c, CodeInvalidParam)
	}

	var resp = vo.ExampleResp{
		Name: req.Name,
	}

	// response 报文
	ResponseSuccess(c, &resp)
}
