/**
 * @Author: yinjinlin
 * @File:  gopicture
 * @Description:
 * @Date: 2021/8/4 下午1:21
 */

package gopicture

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"os"
)


// 加载图片
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer   file.Close()
	img, _,err = image.Decode(file)
	return
}

func saveImage(path string, img image.Image) (err error) {
	file, err := os.Open(path)

	if err != nil {
		return
	}

	defer file.Close()
	return
}
func Picture() {
	src, err := LoadImage("123.png")
	if err != nil {
		log.Fatal(err)
	}


	// 图片定义
	src = imaging.CropAnchor(src,300, 300, imaging.Center)
	// src = imaging.Fill(src, 100, 100, imaging.Center, imaging.Lanczos)

	// 调整图片宽度
	src = imaging.Resize(src, 200,0,imaging.Lanczos)
	// 模糊
	// img1 := imaging.Blur(src,5)

	// 灰度
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2,20)  // 对比度，20 是百分比
	img2 = imaging.Sharpen(img2,2)

	dst := imaging.New(200,200,color.NRGBA{0,0,0,0})
	dst = imaging.Paste(dst,img2,image.Pt(0,0))


	err = imaging.Save(dst, "out_example.jpg")
	if err != nil {
		log.Fatal("failed to save image: ")
	}



}
