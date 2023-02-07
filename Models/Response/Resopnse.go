package Response

// _ResponsePostList 接口响应数据
type _ResponsePostList struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
	Data    string  `json:"data"`    // 数据
}
