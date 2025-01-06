package code

var (
	// Success                   = NewError(1, "成功")
	Success       = NewSuss(1, "成功")
	Failed        = NewError(0, "失败")
	SuccessCreate = NewSuss(2, "新建成功")
	SuccessUpdate = NewSuss(3, "更新成功")
	SuccessDelete = NewSuss(4, "删除成功")

	ErrorNotFoundAPI          = NewError(incr(400), "找不到API")
	ErrorInvalidParams        = NewError(incr(400), "参数验证失败")
	ErrorTooManyRequests      = NewError(incr(400), "请求过多")
	ErrorInvalidAuthToken     = NewError(incr(400), "验证授权Token失败")
	ErrorNotUserAuthToken     = NewError(incr(400), "尚未登录,请先登录")
	ErrorInvalidUserAuthToken = NewError(incr(400), "登录状态失效,请重新登录")
	ErrorInvalidToken         = NewError(incr(400), "缺少用户凭证Token")
	ErrorServerInternal       = NewError(incr(500), "服务内部错误")
)
