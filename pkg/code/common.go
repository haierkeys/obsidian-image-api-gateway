package code

var (
	// Success                   = NewError(1, "成功")
	Success       = NewSuss(1, "成功")
	Failed        = NewError(0, "失败")
	SuccessCreate = NewSuss(2, "新建成功")
	SuccessUpdate = NewSuss(3, "更新成功")
	SuccessDelete = NewSuss(4, "删除成功")

	ErrorServerInternal = NewError(incr(500), "服务内部错误")

	ErrorNotFoundAPI             = NewError(incr(400), "找不到API")
	ErrorInvalidParams           = NewError(incr(400), "参数验证失败")
	ErrorTooManyRequests         = NewError(incr(400), "请求过多")
	ErrorInvalidAuthToken        = NewError(incr(400), "访问令牌效验失败")
	ErrorNotUserAuthToken        = NewError(incr(400), "尚未登录,请先登录")
	ErrorInvalidUserAuthToken    = NewError(incr(400), "登录状态失效,请重新登录")
	ErrorInvalidToken            = NewError(incr(400), "您的访问缺少用户令牌")
	ErrorTokenExpired            = NewError(incr(400), "用户令牌已过期")
	ErrorUserRegister            = NewError(incr(400), "用户注册失败")
	ErrorPasswordNotValid        = NewError(incr(400), "密码不符合规则")
	ErrorUserLoginPasswordFailed = NewError(incr(400), "密码错误")
	ErrorUserRegisterIsDisable   = NewError(incr(400), "用户注册已关闭")
	ErrorUserLoginFailed         = NewError(incr(400), "用户登录失败")
	ErrorUserNotFound            = NewError(incr(400), "用户不存在")
	ErrorUserEmailAlreadyExists  = NewError(incr(400), "用户邮箱已存在")
	ErrorUserUsernameNotValid    = NewError(incr(400), "用户名不符合规则,用户名长度为3-15位,只能包含字母、数字或下划线")
	ErrorUserPasswordNotMatch    = NewError(incr(400), "两次输入的密码不一致")
	ErrorDBQuery                 = NewError(incr(400), "数据库查询失败")

	ErrorUploadFileFailed        = NewError(incr(600), "上传文件失败")
	ErrorInvalidCloudStorageType = NewError(incr(600), "云存储类型无效")
	ErrorInvalidStorageType      = NewError(incr(600), "存储类型无效")

	ErrorInvalidCloudStorageAccountId = NewError(incr(600), "云存储账户ID无效")
	ErrorInvalidCloudStorageRegion    = NewError(incr(600), "云存储区域无效")
	ErrorInvalidCloudStorageEndpoint  = NewError(incr(600), "云存储端点无效")
)
