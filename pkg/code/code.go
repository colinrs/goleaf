package code

import "net/http"

var (
	OKErr                    = &Err{HTTPCode: http.StatusOK, Code: 0, Msg: "success"}
	ErrParam                 = &Err{HTTPCode: http.StatusOK, Code: 10003, Msg: "参数有误"}
	UnknownErr               = &Err{HTTPCode: http.StatusOK, Code: 10004, Msg: "unknown error"}
	ErrBind                  = &Err{HTTPCode: http.StatusOK, Code: 10002, Msg: "Error occurred while binding the request body to the struct."}
	ErrSignParam             = &Err{HTTPCode: http.StatusOK, Code: 10004, Msg: "签名参数有误"}
	ErrValidation            = &Err{HTTPCode: http.StatusOK, Code: 20001, Msg: "Validation failed."}
	ErrDatabase              = &Err{HTTPCode: http.StatusOK, Code: 20002, Msg: "Database error."}
	ErrToken                 = &Err{HTTPCode: http.StatusOK, Code: 20003, Msg: "Error occurred while signing the JSON web token."}
	ErrInvalidTransaction    = &Err{HTTPCode: http.StatusOK, Code: 20004, Msg: "invalid transaction."}
	ErrEncrypt               = &Err{HTTPCode: http.StatusOK, Code: 20101, Msg: "密码加密错误"}
	ErrUserNotFound          = &Err{HTTPCode: http.StatusOK, Code: 20102, Msg: "用户不存在"}
	ErrTokenInvalid          = &Err{HTTPCode: http.StatusOK, Code: 20103, Msg: "Token错误"}
	ErrPasswordIncorrect     = &Err{HTTPCode: http.StatusOK, Code: 20104, Msg: "密码错误"}
	ErrUserExistBefor        = &Err{HTTPCode: http.StatusOK, Code: 20105, Msg: "用户已存在"}
	ErrUserCreate            = &Err{HTTPCode: http.StatusOK, Code: 20105, Msg: "用户创建错误"}
	ErrSendSMSTooMany        = &Err{HTTPCode: http.StatusOK, Code: 20109, Msg: "已超出当日限制，请明天再试"}
	ErrVerifyCode            = &Err{HTTPCode: http.StatusOK, Code: 20110, Msg: "验证码错误"}
	ErrEmailOrPassword       = &Err{HTTPCode: http.StatusOK, Code: 20111, Msg: "邮箱或密码错误"}
	ErrTwicePasswordNotMatch = &Err{HTTPCode: http.StatusOK, Code: 20112, Msg: "两次密码输入不一致"}
	ErrRegisterFailed        = &Err{HTTPCode: http.StatusOK, Code: 20113, Msg: "注册失败"}
	ErrCreatedUser           = &Err{HTTPCode: http.StatusOK, Code: 20114, Msg: "用户创建失败"}
	ErrDuplicateUser         = &Err{HTTPCode: http.StatusOK, Code: 20115, Msg: "用户已经注册"}
	ErrDuplicateProject      = &Err{HTTPCode: http.StatusOK, Code: 20116, Msg: "项目已经存在"}
	ErrProjectNotExist       = &Err{HTTPCode: http.StatusOK, Code: 20117, Msg: "项目不存在"}
)
