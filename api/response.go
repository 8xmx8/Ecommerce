package api

type Response struct {
	Msg  string `json:"message"`
	Code uint8  `json:"code" example:"200"`
}

// OK 正常返回
func (res *Response) OK(msg string) *Response {
	res.Code = 200
	res.Msg = msg
	return res
}

// Error 错误返回
func (res *Response) Error(msg string) *Response {
	res.Code = 0
	res.Msg = msg
	return res
}
