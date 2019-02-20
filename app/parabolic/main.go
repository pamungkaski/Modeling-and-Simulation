package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"../../../mosi"
)

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Parabolic"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	prbl := mosi.NewParabolic(10, 0, 0, 60, 9.8)
	//fmt.Println(prbl.InitialAngel, prbl.InitialSpeedX, prbl.InitialSpeedY)
	// A parabolic function
	quad := plotter.NewFunction(func(x float64) float64 {
		//fmt.Println(x)
		t := prbl.TatPositionX(x)
		return prbl.PositionYatT(t)
	})
	quad.Color = color.RGBA{B: 255, A: 255}


	p.Add(quad)
	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	fmt.Println(prbl.MaxX(), prbl.MaxY())
	p.X.Min = 0
	p.X.Max = prbl.MaxX() * 1.1
	p.Y.Min = 0
	p.Y.Max = prbl.MaxY() * 2.5

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "parabolic.svg"); err != nil {
		panic(err)
	}
}