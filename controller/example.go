package controller

import (
	"github.com/gin-gonic/gin"
	"web-app/model/vo"
)

// ExampleHandler 用户登录
// @Summary      获取用户信息Summary
// @Description  获取用户信息Description
// @Tags         获取用户信息Tags
// @Accept       json
// @Produce      json
// @Param        user body vo.ExampleReq  true  "请求 body 中的数据"
// @Success      200  {object}  ResponseData{Data=vo.ExampleResp}
// @Router       /api/v1/example [post]
func ExampleHandler(c *gin.Context) {
	var req vo.ExampleReq
	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		// response 报文
		ResponseError(c, CodeInvalidParam)
	}

	var resp = vo.ExampleRespData{
		Name: req.Name,
	}

	// response 报文
	ResponseSuccess(c, &resp)
}
