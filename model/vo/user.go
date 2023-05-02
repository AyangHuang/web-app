package vo

type ExampleReq struct {
	Name string `json:"name" binding:"required"`
}

type ExampleResp struct {
	Name string `json:"name" binding:"required"`
}
