package main

import (
	"fmt"

	"github.com/muchrm/golang-img-merge/imgmerge"
)

func main() {
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
		imgmerge.MergeImage(img, "temp/shirt.png", fmt.Sprintf("%s%d%s", "temp/out", i, ".png"))
	}

}
