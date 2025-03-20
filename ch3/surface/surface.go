package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 1000
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 // 1rad == 360 -> 1/6rads == 30deg
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", drawMesh)
	log.Fatal(http.ListenAndServe("localhost:3141", nil))
}

func drawMesh(w http.ResponseWriter, r *http.Request) {
	// draw forms
	fmt.Fprintf(w, "<!DOCTYPE html>\n<html>")

	fmt.Fprintf(w, `
	<form action="/" method="get">
		<label for="redSaturation">Red saturation</label>
		<input type="range"  max=100 min=0 name="redSaturation" id="redSaturation">
		
		<label for="blueSaturation">Blue saturation</label>
		<input type="range" max=100 min=0 name="blueSaturation" id="blueSaturation">

		<label for="sin">Sin multiplier</label>
		<input type="range" max=100 min=0 name="sin" id="sin">

		<label for="viewAngle">View Angle</label>
		<input type="range"  max=100 min=0 name="viewAngle" id="viewAngle"/>

	<input type="submit" value="Update sliders"/>
	</form>
	`)

	// w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	fmt.Printf("blueSat: %s\nredSat: %s\nview: %s\nsin: %s\n", r.FormValue("blueSaturation"), r.Form.Get("redSaturation"), r.Form.Get("viewAngle"), r.Form.Get("sin"))

	blueMultiplierStr := r.Form.Get("blueSaturation")
	if blueMultiplierStr == "" {
		blueMultiplierStr = "0"
	}
	blueMultiplier, err := strconv.ParseFloat(blueMultiplierStr, 64)
	if err != nil {
		panic(err)
	}
	blueMultiplier = blueMultiplier / 100
	redMultiplierStr := r.Form.Get("redSaturation")
	if redMultiplierStr == "" {
		redMultiplierStr = "0"
	}
	redMultiplier, err := strconv.ParseFloat(redMultiplierStr, 64)
	if err != nil {
		panic(err)
	}
	redMultiplier = redMultiplier / 100
	// viewAngle, err := strconv.ParseFloat(r.Form.Get("viewAngle"), 64)
	// if err != nil {
	// 	panic(err)
	// }
	// sinMultiplier, err := strconv.ParseFloat(r.Form.Get("sin"), 64)
	// if err != nil {
	// 	panic(err)
	// }

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			blueness := int((float64(i) / cells) * 0xFF * blueMultiplier)

			redness := int((float64(j) / cells) * 0xFF * redMultiplier)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%02X00%02X'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, redness, blueness)
		}
	}
	fmt.Fprintln(w, "</svg>")
	fmt.Fprintf(w, "</html>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x, y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z, ok := f(x, y)
	if !ok {
		return -1, -1, false
	}

	// Project (x, y, z) isometrically onto 2-D SVG canbas(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0.0)
	r = math.Sin(r)
	if math.IsNaN(r) {
		return -1, false
	}
	return r, true //
}
