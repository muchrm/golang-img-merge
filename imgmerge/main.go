package imgmerge

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Pixel struct {
	Point image.Point
	Color color.Color
}

func (p *Pixel) IsTransparent() bool {
	_, _, _, alpha := p.Color.RGBA()
	return alpha <= 30000
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
func WriteImage(img image.Image, outImg string) error {
	out, err := os.Create(outImg)
	if err != nil {
		return err
	}
	err = png.Encode(out, img)
	return err
}
func MergeImage(imgName string, beseImg string, outImg string) error {
	img1, _, err := OpenAndDecode(beseImg)
	if err != nil {
		return err
	}
	img2, _, err := OpenAndDecode(imgName)
	if err != nil {
		return err
	}
	pixels1 := DecodePixelsFromImage(img1, 0, 0)
	pixels2 := DecodePixelsFromImage(img2, (img1.Bounds().Dx()/2)-(img2.Bounds().Dx()/3), (img1.Bounds().Dy()/2)-(img2.Bounds().Dy()/2))
	img := image.NewRGBA(image.Rect(0, 0, img1.Bounds().Dx(), img1.Bounds().Dy()))
	for _, px := range append(pixels1, pixels2...) {
		if !px.IsTransparent() {
			img.Set(px.Point.X, px.Point.Y, px.Color)
		}
	}
	draw.Draw(img, img.Bounds(), img, image.Point{0, 0}, draw.Src)
	err = WriteImage(img, outImg)
	return err
}
