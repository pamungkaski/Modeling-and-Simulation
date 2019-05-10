package main

import (
	"github.com/pamungkaski/Modeling-and-Simulation"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"log"
	"gonum.org/v1/plot/vg"
)

func main() {
	// the Percentage will follow
	// "Northeast", "East", "Southeast", "South", "Southwest", "West", "Northwest", "North"
	// Finish is the plane dimension to move
	randomwalk := mosi.NewRandomWalk([]float64{0.24, 0.41, 0.51, 0.53, 0.56, 0.66, 0.81, 1.00}, 20)
	fmt.Println(randomwalk.Moves)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	if err := plotutil.AddLinePoints(p, randomwalk.Route); err != nil {
		log.Fatal(err.Error())
	}

	// Save the plot to a SVG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "random_walk.svg"); err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "random_walk.png"); err != nil {
		panic(err)
	}
}