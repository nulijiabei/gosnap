package main

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"image"
)

// 截图
func CaptureScreen() (*image.RGBA, error) {

	// 屏幕连接
	var c *xgb.Conn
	
	// XGB
		var xgbConn *xgb.Conn   

	// 判断是否连接到了屏幕
	if xgbConn== nil {
		// 建立连接
		conn, err := xgb.NewConn()
		// 判断是否连接成功
		if err != nil {
			return nil, err
		}
		// 成功保存全局
		xgbConn = conn
		// 赋值
		c = conn
	} else {
		// 已经连接
		c = xgbConn
	}

	screen := xproto.Setup(c).DefaultScreen(c)

	rect := image.Rect(0, 0, int(screen.WidthInPixels), int(screen.HeightInPixels))

	// ------------------------------------------ //

	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(c, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
	
}
