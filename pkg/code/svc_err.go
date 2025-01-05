package code

var (
	ErrorUploadFileFail          = NewError(incr(600), "上传文件失败")
	ErrorUserRegister            = NewError(incr(600), "用户注册失败")
	ErrorPasswordNotValid        = NewError(incr(600), "密码不符合规则")
	ErrorUserLoginPasswordFailed = NewError(incr(600), "密码错误")
	ErrorUserLoginFailed         = NewError(incr(600), "用户登录失败")
	ErrorUserNotFound            = NewError(incr(600), "用户不存在")
	ErrorUserEmailAlreadyExists  = NewError(incr(600), "用户邮箱已存在")
	ErrorTokenExpired            = NewError(incr(600), "token已过期")
)
