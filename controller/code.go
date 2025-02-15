package controller

type ResCode int32

const (
	CodeSuccess    ResCode = 0
	CodeServerBusy ResCode = 4999 + iota
	CodeInvalidParam
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeServerBusy:   "服务繁忙",
	CodeInvalidParam: "请求参数错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
