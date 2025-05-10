package helpers

func FilterColorfulColor(pixels []int) (isColorful bool) {
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

	if saturation >= threshold {
		isColorful = true
	}

	return
}
