package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/pamungkaski/Modeling-and-Simulation"
	"gonum.org/v1/plot/plotutil"
	"log"
)

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Parabolic"

	prbl := mosi.NewParabolic(140, 0, 0, 30, 9.8, -10)
	pts := plotter.XYs{}
	minus := 0
	maxX := 0.0
	maxY := 0.0
	for i := 0; minus < 2 ; i++{
		pets := plotter.XY{}
		pets.X =  prbl.PositionXatT(float64(i)/float64(10))
		pets.Y =  prbl.PositionYatT(float64(i)/float64(10))
		fmt.Println(pets, float64(i)/float64(10))
		if maxX < pets.X {
			maxX = pets.X
		}
		if maxY < pets.Y {
			maxY = pets.Y
		}
		if pets.Y < 0 {
			minus++
			pts = append(pts, pets)
		} else {
			pts = append(pts, pets)
		}
	}
	 if err := plotutil.AddLinePoints(p, pts); err != nil {
		log.Fatal(err.Error())
	 }

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	fmt.Println(prbl.MaxX(), prbl.MaxY())
	if pts[pts.Len() - 1].X < 0 {
		p.X.Min = pts[pts.Len() - 1].X
	} else {
		p.X.Min = 0
	}
	p.X.Max = maxX * 1.5
	p.Y.Min = 0
	p.Y.Max = maxY * 2

	// Save the plot to a SVG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "parabolic.svg"); err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "parabolic.png"); err != nil {
		panic(err)
	}

	// ini kebawah juga gk guna
	fmt.Println("Maximum Height : ", prbl.MaxY())
	fmt.Println("Maximum Range : ", prbl.MaxX())
	fmt.Println("Postion at t = 0.2: ", "x = ", prbl.PositionXatT(0.2), "y = ", prbl.PositionYatT(0.2))
	fmt.Println("Postion at t = 0.3: ", "x = ", prbl.PositionXatT(0.3), "y = ", prbl.PositionYatT(0.3))
	fmt.Println("How long the bullet in the air: ", prbl.TimeInAir(), "s")
}