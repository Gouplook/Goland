/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/11 17:49
@Description:

*********************************************/
package errno

import "fmt"

type GolandError struct {
	Code int
	Message string
}

func (e *GolandError)Error() string{
	return fmt.Sprintf("Goland error, code:%d message:%v", e.Code, e.Message)
}

var (
	NotHaveInstance = &GolandError{
		Code:    1,
		Message: "not have instance",
	}
	ConnFailed = &GolandError{
		Code:    2,
		Message: "connect failed",
	}
	InvalidNode = &GolandError{
		Code:    3,
		Message: "invalid node",
	}
	AllNodeFailed = &GolandError{
		Code:    4,
		Message: "all node failed",
	}
)

func IsConnectError(err error) bool{
	goErr, ok := err.(*GolandError)
	if !ok {
		return false
	}
	var result bool
	if goErr == ConnFailed {
		result = true
	}
	return result
}