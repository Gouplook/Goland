/**
 * @Author: yinjinlin
 * @File:  goregx_test
 * @Description:
 * @Date: 2021/10/29 下午5:10
 */

package regx

import "testing"

func TestPhoneRegx(t *testing.T) {
	// phone := "02154377032"
	// phone := "021-5181939909"
	// 68 = 16 + 10 + 5 =
	phone :="0411-83613910"
	PhoneRegx(phone)
}
