package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("colors.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	palette := map[color.NRGBA]int{}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y).(color.NRGBA)

			if v, ok := palette[c]; ok {
				palette[c] = v + 1
			} else {
				palette[c] = 1
			}
		}
	}

	/*
		for c, v := range palette {
			if v > 7 {
				rgb := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
				//fmt.Printf("<div style=\"background:%s;height:50px\">%s</div>\n", k, k)
				fmt.Printf("\"%s\",\n", rgb)
			}
		}
	*/

	toRemove := map[color.NRGBA]int{}

	for c, v := range palette {
		if v < 7 {
			toRemove[c] = 1
			continue
		}

		if _, ok := toRemove[c]; ok {
			continue
		}

		for oc, _ := range palette {
			if diff(c, oc) < 32 && diff(c, oc) != 0 {
				//rgb := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
				//rgb2 := fmt.Sprintf("#%02x%02x%02x", oc.R, oc.G, oc.B)
				//fmt.Printf("too close, %s - %s\n", rgb, rgb2)
				toRemove[oc] = 1
			}
		}
	}

	println()

	for c, _ := range toRemove {
		delete(palette, c)
	}

	for c, _ := range palette {
		rgb := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
		//fmt.Printf("%s\n", rgb)
		fmt.Printf("<div style=\"background:%s;height:50px\">%s</div>\n", rgb, rgb)
	}

	fmt.Println("<pre>")
	for c, _ := range palette {
		rgb := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
		fmt.Printf("\"%s\",\n", rgb)
	}
	fmt.Println("</pre>")
}

func diff(c1, c2 color.NRGBA) int {
	dr := abs(c1.R - c2.R)
	dg := abs(c1.G - c2.G)
	db := abs(c1.B - c2.B)

	return dr + dg + db
}

func abs(i uint8) int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}
