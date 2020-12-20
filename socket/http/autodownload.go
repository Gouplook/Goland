/*******************************
 * @Author: Yinjinlin
 * @Description:
 * @File:  autodownload
 * @Version: 1.0.0
 * @Date: 2020/11/29 17:44
 **********************************/
package http

import "io"

type Reader struct {
	io.Reader
	Total int64
	Current int64
}

func (r *Reader)Read(p []byte)(n int, err error){
	//n,err := r.Reader.Read(p)
	//r.Current +=int64(n)
	return
}
// 文件自动下载
func AutoDownload(){

}