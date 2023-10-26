package utils

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func GenerateAvatar(width, height, blockWidth, blockHeight int, outputPath string) error {

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 随机生成像素块
	for x := 0; x <= width/2; x++ {
		for y := 0; y < height; y++ {
			// 每个块的左上角像素
			if x%blockWidth == 0 && y%blockHeight == 0 {
				r := uint8(rand.Uint32())
				g := uint8(rand.Uint32())
				b := uint8(rand.Uint32())
				a := uint8(rand.Uint32())
				img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
				// 使头像左右对称，设置对应的右半部分像素块颜色相同
				x2 := width - x
				if x2 >= 0 {
					img.Set(x2, y, color.RGBA{R: r, G: g, B: b, A: a})
				}
			} else {
				// 如果不是块的左上角像素，则使用相同的颜色，模拟块的颜色
				c := img.RGBAAt(x-x%blockWidth, y-y%blockWidth)
				img.Set(x, y, c)

				// 使头像左右对称，设置对应的右半部分像素块颜色相同
				x2 := width - x
				if x2 >= 0 {
					img.Set(x2, y, c)
				}
			}
		}
	}

	// 保存图片
	f, _ := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	err := png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}

func QuickGenAvatar(outputPath string) error {
	return GenerateAvatar(420, 420, 140, 140, outputPath)
}
