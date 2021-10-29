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
	// phone := "021-51819399019999"
	phone :="04122-81760160"
	PhoneRegx(phone)
}
