package api

import "github.com/gin-gonic/gin"

// Option option
type Option struct {
	Code *int
	Msg  *string
	Data *interface{}
}

// WithCode With Code
func WithCode(code int) *Option {
	val := code
	return &Option{
		Code: &val,
	}
}

// WithMsg With Msg
func WithMsg(msg string) *Option {
	return &Option{
		Msg: &msg,
	}
}

// WithData With Data
func WithData(data interface{}) *Option {
	return &Option{
		Data: &data,
	}
}

func renderRespData(defaultCode int, options ...*Option) map[string]interface{} {
	respData := make(map[string]interface{})
	respData["code"] = defaultCode

	for _, option := range options {
		if option.Code != nil {
			respData["code"] = option.Code
		}
		if option.Msg != nil {
			respData["msg"] = option.Msg
		}
		if option.Data != nil {
			respData["data"] = option.Data
		}
	}

	return respData
}

// Ok Ok
func Ok(c *gin.Context, options ...*Option) {
	respData := renderRespData(0, options...)
	c.JSON(0, respData)
}

// Error Error
func Error(c *gin.Context, options ...*Option) {
	respData := renderRespData(500, options...)
	c.JSON(0, respData)
}
