package vo

type Result struct {
	Sucess bool        `json:"sucess"`
	Code   uint32      `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func NewResultSuccess(date interface{}) *Result {
	r := &Result{
		Sucess: true,
		Code:   0,
		Data:   date,
	}
	return r
}

func NewResultFail(msg string) *Result {
	r := &Result{
		Sucess: false,
		Code:   1,
		Msg:    msg,
	}
	return r
}
