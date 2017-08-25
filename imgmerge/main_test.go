package imgmerge

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestIsTransparent(t *testing.T) {
	type args struct {
		r uint8
		g uint8
		b uint8
		a uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Test IsTranparent $1", args{0, 0, 0, 0}, true},
		{"Test IsTranparent $2", args{0, 0, 0, 255}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pixel := &Pixel{
				Point: image.Point{0, 0},
				Color: color.RGBA{tt.args.r, tt.args.g, tt.args.b, tt.args.a},
			}
			got := pixel.IsTransparent()
			if got != tt.want {
				t.Errorf("IsTranparent %v, want %v", got, tt.want)
			}
		})
	}
}
func TestOpenAndDecode(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test OpenAndDecode $1", "../temp/shirt.jpg", false},
		{"Test OpenAndDecode $2", "../shirt.jpg", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := OpenAndDecode(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenAndDecod error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
func TestMergeImage(t *testing.T) {

	imgs := []string{
		"temp/1.png",
		"temp/2.png",
		"temp/3.png",
		"temp/4.png",
		"temp/5.png",
		"temp/6.png",
		"temp/7.png",
		"temp/8.png",
		"temp/9.png",
		"temp/10.png",
		"temp/11.png",
	}

	for i, img := range imgs {
		t.Run("", func(t *testing.T) {
			MergeImage(img, "temp/shirt.png", fmt.Sprintf("%s%d%s", "temp/out", i, ".png"))
		})
	}
}
