package vo

type ExampleReq struct {
	Name string `json:"name" binding:"required"`
}

type ExampleRespData struct {
	Name string `json:"name" binding:"required"`
}
