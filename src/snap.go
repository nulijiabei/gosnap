package main

import (
	z "github.com/nutzam/zgo"
)

// ...
func main() {

	// 获取屏幕截图
	img, err := CaptureScreen()
	if err != nil {
		panic(err)
	}

	// 保存图片
	err = z.ImageEncodeJPEG("/dev/shm/screenshot.jpg", img)
	if err != nil {
		panic(err)
	}
	

}
