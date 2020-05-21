package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	maxZ          = 0.9850673555377986
	minZ          = -0.21722891503668823
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var colors = [...]string{"0000ff", "1b00e4", "3600c9", "5200ad", "6e0091", "8a0075",
	"8f0070", "ab0054", "c70038", "e3001c", "ff0000"}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			curMaxZ := az
			bx, by, bz := corner(i, j)
			curMaxZ = math.Max(curMaxZ, bz)
			cx, cy, cz := corner(i, j+1)
			curMaxZ = math.Max(curMaxZ, cz)
			dx, dy, dz := corner(i+1, j+1)
			curMaxZ = math.Max(curMaxZ, dz)

			if math.IsNaN(curMaxZ) {
				continue
			}

			rate := (curMaxZ - minZ) / (maxZ - minZ)
			color := int(rate*10)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colors[color])
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
