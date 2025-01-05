package code

var (
	// Success                   = NewError(1, "成功")
	Success = NewSuss(1, "成功")

	ErrorNotFoundAPI      = NewError(incr(400), "找不到API")
	ErrorInvalidParams    = NewError(incr(400), "参数验证失败")
	ErrorTooManyRequests  = NewError(incr(400), "请求过多")
	ErrorInvalidAuthToken = NewError(incr(400), "验证授权Token失败")
	ErrorInvalidToken     = NewError(incr(400), "缺少用户凭证Token")

	ErrorServerInternal = NewError(incr(500), "服务内部错误")
)
