package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

type rgba struct {
	red   uint32
	green uint32
	blue  uint32
}

func get_most_used_color() {
}

func main() {
	reader, err := os.Open("./image/Designer.jpeg")
	if err != nil {
		panic(err)
	}
	m, _, err := image.Decode(reader)
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
		r_total += r >> 8
		g_total += g >> 8
		b_total += b >> 8
		count += 1
		fmt.Println(count)

		fmt.Printf("red: %v \n green: %v \n blue: %v \n", r_total/count, g_total/count, b_total/count)
	}
}
