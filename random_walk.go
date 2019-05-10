package mosi

import (
	"time"
	"math/rand"
	"gonum.org/v1/plot/plotter"
)

type RandomWalk struct {
	Direction []string
	DirectionPercentage []float64
	MoveX []int
	MoveY []int
	CurrentStateX int
	CurrentStateY int
	Route plotter.XYs
	Moves []string
	Finish int
}

func NewRandomWalk (percentage []float64, finish int) (RandomWalk) {
	r := RandomWalk{
		Direction: []string{"Northeast", "East", "Southeast", "South", "Southwest", "West", "Northwest", "North"},
		DirectionPercentage:percentage,
		MoveX: []int{1, 1, 1, 0, -1, -1, -1, 0},
		MoveY: []int{1, 0, -1, -1, -1, 0, 1, 1},
		CurrentStateX: 0,
		CurrentStateY: 0,
		Finish: finish,
	}

	for r.CurrentStateX != r.Finish && r.CurrentStateY != r.Finish  {
		r.NextStep()
	}
	return r
}


func (r *RandomWalk) NextStep () {
	nextX := r.CurrentStateX
	nextY := r.CurrentStateY
	moves := false
	move := -1
	source := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source)

	for !moves {
		randomNumber := r1.Float64()
		for i, perc := range r.DirectionPercentage {
			if perc > randomNumber {
				move = i
				nextX += r.MoveX[i]
				nextY += r.MoveY[i]
				break
			}
		}
		if nextX >= 0 && nextX <= r.Finish && nextY >= 0 && nextY <= r.Finish {
			moves = true
		}
	}
	r.Moves = append(r.Moves, r.Direction[move])
	r.CurrentStateX = nextX
	r.CurrentStateY = nextY
	r.Route = append(r.Route, plotter.XY{X: float64(nextX), Y: float64(nextY)})
}