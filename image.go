package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type rgb struct {
	red   uint32
	green uint32
	blue  uint32
}

func getMostUsedColor(r *os.File) rgb {
	m, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}

	bounds := m.Bounds()

	r_total := uint32(0)
	g_total := uint32(0)
	b_total := uint32(0)

	count := uint32(0)

	for x := bounds.Min.X; x < bounds.Max.X; x += 100 {
		r, g, b, _ := m.At(x, 0).RGBA()
		count += 1
		r_total += r >> 8
		g_total += g >> 8
		b_total += b >> 8
	}

	r_total = r_total / count
	g_total = g_total / count
	b_total = b_total / count

	return rgb{r_total, g_total, b_total}
}

func main() {
	reader, err := os.Open("./image/CuteCat.png")
	if err != nil {
		panic(err)
	}
	rgb := getMostUsedColor(reader)

	fmt.Printf("red: %v \ngreen: %v \nblue: %v \n", rgb.red, rgb.green, rgb.blue)
}
