package main

import "github.com/muchrm/golang-img-merge/lib"

func main() {
	images := []string{"temp/a.png", "temp/a.png", "temp/a.png", "temp/a.png", "temp/a.png", "temp/a.png"}
	for _, image := range images {
		lib.MergeImage(image, "./temp/shirt.jpg", "./temp/output.jp")
	}
}
