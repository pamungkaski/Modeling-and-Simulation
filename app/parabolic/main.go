package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/pamungkaski/Modeling-and-Simulation"
)

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Parabolic"

	prbl := mosi.NewParabolic(10, 0, 0, 60, 9.8)
	// A parabolic function
	quad := plotter.NewFunction(func(x float64) float64 {
		//fmt.Println(x)
		t := prbl.TatPositionX(x)
		return prbl.PositionYatT(t)
	})
	quad.Color = color.RGBA{B: 255, A: 255}

	// An fxt
	fxt := plotter.NewFunction(func(t float64) float64 { return prbl.PositionXatT(t) })
	fxt.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	fxt.Width = vg.Points(2)
	fxt.Color = color.RGBA{G: 255, A: 255}

	// The fyt
	fyt := plotter.NewFunction(func(t float64) float64 { return prbl.PositionYatT(t) })
	fyt.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
	fyt.Width = vg.Points(4)
	fyt.Color = color.RGBA{R: 255, A: 255}

	p.Add(quad, fxt, fyt)
	p.Legend.Add("Fx(y)", quad)
	p.Legend.Add("Fx(t)", fxt)
	p.Legend.Add("Fy(t)", fyt)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	fmt.Println(prbl.MaxX(), prbl.MaxY())
	p.X.Min = 0
	p.X.Max = prbl.MaxX() * 1.1
	p.Y.Min = 0
	p.Y.Max = prbl.MaxY() * 1.5

	// Save the plot to a SVG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "parabolic.svg"); err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "parabolic.png"); err != nil {
		panic(err)
	}
	fmt.Println("Maximum Height : ", prbl.MaxY())
	fmt.Println("Maximum Range : ", prbl.MaxX())
	fmt.Println("Postion at t = 0.2: ", prbl.PositionXatT(0.2), prbl.PositionYatT(0.2))
	fmt.Println("Postion at t = 0.3: ", prbl.PositionXatT(0.3), prbl.PositionYatT(0.3))
	fmt.Println("How long the bullet in the air: ", prbl.TimeInAir())
}