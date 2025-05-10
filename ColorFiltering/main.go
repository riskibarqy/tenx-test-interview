package main

import (
	"fmt"
)

func filterColor(pixels []int) (isColorful bool) {
	threshold := 0.2
	min, max := pixels[0], pixels[0]

	for _, v := range pixels {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	saturation := float64((max - min)) / float64(max)

	if saturation > threshold {
		isColorful = true
	}

	return
}

func main() {
	image := [][]int{
		{34, 203, 55}, {67, 76, 73}, {99, 105, 93}, {178, 173, 169}, {144, 89, 54},
		{22, 20, 18}, {10, 40, 50}, {171, 180, 211}, {150, 150, 90}, {50, 150, 150},
	}

	newImages := [][]int{}

	for _, v := range image {
		if filterColor(v) {
			// fmt.Println("colorful")
			newImages = append(newImages, v)
		} else {
			// fmt.Println("not colorful")
			newImages = append(newImages, []int{0, 0, 0})
		}
	}

	fmt.Println(newImages)

}
