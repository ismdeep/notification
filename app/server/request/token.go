package request

// NewToken 新建token结果
type NewToken struct {
	Name string `json:"name" binding:"required"`
}
