package main

import( 
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{
	p [][]uint8
}
func (i Image) At(x, y int) color.Color {
	v:= i.p[x][y]
	return color.RGBA{v, v, 255, 255}
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i.p), len(i.p[0]))
}
func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8,dy)

	for i:=0; i < dy; i++{
		y[i] = make([]uint8,dx)
		for j:=0; j<dx ; j++{
			y[i][j] = uint8((i+j)/2)
			//y[i][j] = uint8(i^j)
		}
	}
	return y
}

func main() {
	m := Image{Pic(90,100)}
	pic.ShowImage(m)
}

