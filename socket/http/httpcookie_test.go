/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  httpcookie_test
 * @Version: 1.0.0
 * @Date: 2020/12/20 21:39
 */
package http

import (
	"net/http"
	"testing"
)

func TestCookie(t *testing.T) {
	http.HandleFunc("/setCookie" ,setCookie)
	http.HandleFunc("/getCookies", getCookie)

	http.ListenAndServe(":10086",nil)
}
