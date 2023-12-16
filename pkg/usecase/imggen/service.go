package imggen

import (
	"context"
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
)

type ImggenService struct {
	layerConfig []LayerConfig
}

func NewImggenService(layerConfig []LayerConfig) *ImggenService {
	return &ImggenService{
		layerConfig: layerConfig,
	}
}

type LayerConfig struct {
	radius int
	number int
	size   int
}

func NewLayerConfig(radius int, number int, size int) LayerConfig {
	return LayerConfig{
		radius: radius,
		number: number,
		size:   size,
	}
}

func (i *ImggenService) GenerateImage(ctx context.Context, icons <-chan image.Image) (image.Image, error) {
	res := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	left := len(icons) // 残りの画像の枚数

	// 各レイヤーごとの処理
LOOP:
	for _, c := range i.layerConfig {
		for i := 0; i < min(c.number, left); i++ {
			var icon image.Image
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case icon = <-icons:
				if icon == nil {
					break LOOP
				}
			}

			// 各レイヤーで定められた大きさにリサイズ
			icon_resized := image.NewRGBA(image.Rect(0, 0, c.size, c.size))
			draw.CatmullRom.Scale(icon_resized, icon_resized.Bounds(), icon, icon.Bounds(), draw.Over, nil)

			// 各レイヤーで定められた位置に，円形に切り抜いてから配置
			theta := 2 * math.Pi * float64(i) / float64(min(c.number, left))
			pos := image.Point{
				X: (1024-c.size)/2 + int(float64(c.radius)*math.Cos(theta)),
				Y: (1024-c.size)/2 - int(float64(c.radius)*math.Sin(theta)),
			}
			draw.DrawMask(
				res,
				image.Rect(pos.X, pos.Y, pos.X+c.size, pos.Y+c.size),
				icon_resized,
				image.Point{},
				&circle{
					p: image.Point{
						X: c.size / 2,
						Y: c.size / 2,
					},
					r: c.size / 2,
				},
				image.Point{},
				draw.Over,
			)
		}
		left -= c.number
	}
	return res, nil
}

// 円形に切り抜くのに必要なやつ
type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
