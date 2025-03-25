package code

var (
	Failed                = NewError(0, lang{zh: "失败", en: "Failed"})
	Success               = NewSuss(1, lang{zh: "成功", en: "Success"})
	SuccessCreate         = NewSuss(2, lang{zh: "创建成功", en: "Create Success"})
	SuccessUpdate         = NewSuss(3, lang{zh: "更新成功", en: "Update Success"})
	SuccessDelete         = NewSuss(4, lang{zh: "删除成功", en: "Delete Success"})
	SuccessPasswordUpdate = NewSuss(5, lang{zh: "密码修改成功", en: "Password Update Success"})

	ErrorServerInternal                     = NewError(incr(500), lang{zh: "服务器内部错误", en: "Server Internal Error"})
	ErrorNotFoundAPI                        = NewError(incr(400), lang{zh: "找不到API", en: "Not Found API"})
	ErrorInvalidParams                      = NewError(incr(400), lang{zh: "参数验证失败", en: "Invalid Params"})
	ErrorTooManyRequests                    = NewError(incr(400), lang{zh: "请求过多", en: "Too Many Requests"})
	ErrorInvalidAuthToken                   = NewError(incr(400), lang{zh: "访问令牌效验失败", en: "Invalid Auth Token"})
	ErrorNotUserAuthToken                   = NewError(incr(400), lang{zh: "尚未登录,请先登录", en: "Not logged in. Please log in first."})
	ErrorInvalidUserAuthToken               = NewError(incr(400), lang{zh: "登录状态失效,请重新登录", en: "Session expired, please log in again."})
	ErrorInvalidToken                       = NewError(incr(400), lang{zh: "您的访问缺少用户令牌", en: "Your access is missing a user token"})
	ErrorTokenExpired                       = NewError(incr(400), lang{zh: "用户令牌已过期", en: "User token has expired"})
	ErrorUserRegister                       = NewError(incr(400), lang{zh: "用户注册失败", en: "User registration failed"})
	ErrorPasswordNotValid                   = NewError(incr(400), lang{zh: "密码不符合规则", en: "Password does not meet the rules"})
	ErrorUserLoginPasswordFailed            = NewError(incr(400), lang{zh: "密码错误", en: "Password error"})
	ErrorUserOldPasswordFailed              = NewError(incr(400), lang{zh: "当前密码验证错误", en: "Current password verification error"})
	ErrorMultiUserPublicAPIClosed           = NewError(incr(400), lang{zh: "多用户开放接口已经关闭,请联系管理员配置 config.user.is-user-enable 选项", en: "Multi-user open interface has been closed, please contact the administrator to configure the config.user.is-user-enable option"})
	ErrorUserRegisterIsDisable              = NewError(incr(400), lang{zh: "用户注册已关闭,请联系管理员配置 config.user.register-is-enable 选项", en: "User registration is closed, please contact the administrator to configure the config.user.register-is-enable option"})
	ErrorUserLoginFailed                    = NewError(incr(400), lang{zh: "用户登录失败", en: "User login failed"})
	ErrorUserNotFound                       = NewError(incr(400), lang{zh: "用户不存在", en: "Username does not exist"})
	ErrorUserAlreadyExists                  = NewError(incr(400), lang{zh: "用户已经存在", en: "Username already exists"})
	ErrorUserEmailAlreadyExists             = NewError(incr(400), lang{zh: "用户邮箱已存在", en: "User email already exists"})
	ErrorUserUsernameNotValid               = NewError(incr(400), lang{zh: "用户名不符合规则,用户名长度为3-15位,只能包含字母、数字或下划线", en: "The username does not meet the rules, the username length is 3-15 digits, and can only contain letters, numbers or underscores"})
	ErrorUserPasswordNotMatch               = NewError(incr(400), lang{zh: "密码与密码确认不一致", en: "Password and password confirmation do not match"})
	ErrorUserPasswordNotValid               = NewError(incr(400), lang{zh: "密码不符合规则,密码长度为6-20位,只能包含字母、数字或下划线", en: "Password does not meet the rules, password length is 6-20 digits, and can only contain letters, numbers or underscores"})
	ErrorDBQuery                            = NewError(incr(400), lang{zh: "数据库查询失败", en: "Database query failed"})
	ErrorUploadFileFailed                   = NewError(incr(400), lang{zh: "上传文件失败", en: "Upload file failed"})
	ErrorInvalidCloudStorageType            = NewError(incr(400), lang{zh: "云存储类型无效", en: "Invalid cloud storage type"})
	ErrorInvalidStorageType                 = NewError(incr(400), lang{zh: "存储类型无效", en: "Invalid storage type"})
	ErrorInvalidCloudStorageBucketName      = NewError(incr(400), lang{zh: "云存储桶名无效", en: "Invalid cloud storage bucket name"})
	ErrorInvalidCloudStorageAccessKeyID     = NewError(incr(400), lang{zh: "云存储访问密钥ID无效", en: "Invalid cloud storage access key ID"})
	ErrorInvalidCloudStorageAccessKeySecret = NewError(incr(400), lang{zh: "云存储访问密钥无效", en: "Invalid cloud storage access key"})
	ErrorInvalidCloudStorageAccountId       = NewError(incr(400), lang{zh: "云存储账户ID无效", en: "Invalid cloud storage account ID"})
	ErrorInvalidCloudStorageRegion          = NewError(incr(400), lang{zh: "云存储区域无效", en: "Invalid cloud storage region"})
	ErrorInvalidCloudStorageEndpoint        = NewError(incr(400), lang{zh: "云存储端点无效", en: "Invalid cloud storage endpoint"})
	ErrorUserCloudflueR2Disabled            = NewError(incr(400), lang{zh: "多用户开放网关存储类型 Cloudflue R2 未开启", en: "Multi-user open gateway storage type Cloudflue R2 is not enabled"})
	ErrorUserALIOSSDisabled                 = NewError(incr(400), lang{zh: "多用户开放网关存储类型 Aliyun OSS 未开启", en: "Multi-user open gateway storage type Aliyun OSS is not enabled"})
	ErrorUserAWSS3Disabled                  = NewError(incr(400), lang{zh: "多用户开放网关存储类型 AWS S3 未开启", en: "Multi-user open gateway storage type AWS S3 is not enabled"})
	ErrorUserMinIODisabled                  = NewError(incr(400), lang{zh: "多用户开放网关存储类型 MinIO 未开启", en: "Multi-user open gateway storage type MinIO is not enabled"})
	ErrorUserLocalFSDisabled                = NewError(incr(400), lang{zh: "多用户开放网关存储类型 服务器本地存储 未开启", en: "Multi-user open gateway storage type server local storage is not enabled"})
	ErrorWebDAVInvalidEndpoint              = NewError(incr(400), lang{zh: "WebDAV服务器URL无效", en: "WebDAV server URL is invalid"})
	ErrorWebDAVInvalidUser                  = NewError(incr(400), lang{zh: "WebDAV服务器用户名不能为空", en: "WebDAV server username is invalid"})
	ErrorWebDAVInvalidPassword              = NewError(incr(400), lang{zh: "WebDAV服务器密码不能为空", en: "WebDAV server URL is invalid"})
	ErrorInvalidAccessURLPrefix             = NewError(incr(400), lang{zh: "访问地址前缀不能为空", en: "Access URL prefix cannot be empty"})
)
