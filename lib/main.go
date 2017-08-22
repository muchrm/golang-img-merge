package lib

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

func (p *Pixel) IsTransparent() bool {
	_, _, _, alpha := p.Color.RGBA()
	return alpha == 0
}
func OpenAndDecode(filepath string) (image.Image, string, error) {
	imgFile, err := os.Open(filepath)
	img, format, err := image.Decode(imgFile)
	imgFile.Close()
	return img, format, err
}
func DecodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}
	}
	return pixels
}
func WriteImage(img image.Image) {
	out, err := os.Create("./temp/output.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
}
func MergeImage(imgName string) {
	img1, _, err := OpenAndDecode("./temp/shirt.jpg")
	if err != nil {
		panic(err)
	}
	img2, _, err := OpenAndDecode(imgName)
	if err != nil {
		panic(err)
	}
	pixels1 := DecodePixelsFromImage(img1, 0, 0)
	pixels2 := DecodePixelsFromImage(img2, 400-(img2.Bounds().Dx()/2), 300-(img2.Bounds().Dy()/2))
	img := image.NewRGBA(image.Rect(0, 0, 800, 600))
	for _, px := range append(pixels1, pixels2...) {
		if !px.IsTransparent() {
			img.Set(px.Point.X, px.Point.Y, px.Color)
		}
	}
	WriteImage(img)
}
