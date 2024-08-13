package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type rgb struct {
	red   uint32
	green uint32
	blue  uint32
}

func getCommonColor(r *os.File) (rgb, error) {
	m, _, err := image.Decode(r)
	if err != nil {
		return rgb{}, err
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

	r_total /= count
	g_total /= count
	b_total /= count

	return rgb{r_total, g_total, b_total}, nil
}

func main() {
	reader, err := os.Open("./image/CuteCat.png")
	if err != nil {
		log.Fatal("Failed to open image: $v", err)
	}

	defer reader.Close()

	rgb, err := getCommonColor(reader)
	if err != nil {
		log.Fatal("Failed to get common color: $v", err)
	}

	fmt.Printf("rgb(%v, %v, %v)\n", rgb.red, rgb.green, rgb.blue)
}
