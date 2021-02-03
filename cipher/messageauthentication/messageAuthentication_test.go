/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/3 14:21
@Description:

*********************************************/
package messageauthentication

import (
	"fmt"
	"testing"
)

func TestGenerateHamc(t *testing.T) {
	src := []byte("256在消息认证码中，需要发送者和接收者之间共享密钥")
	key := []byte("helloworld")

	hmac1 := GenerateHamc(src,key)
	b1 := VerifHamc(src,key,hmac1)

	fmt.Println(b1)
}