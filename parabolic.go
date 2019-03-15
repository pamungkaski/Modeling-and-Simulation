package mosi

import (
	"math"
)

type Parabolic struct {
	InitialSpeed float64
	InitialSpeedX float64
	InitialSpeedY float64
	InitialX float64
	InitialY float64
	InitialAngel float64
	Gravity float64
	WindAcceleration float64
}

func NewParabolic(speed, X, Y, angle, gravity, wind float64) Parabolic {
	return Parabolic{
		InitialSpeed: speed,
		InitialAngel: angle*math.Pi/180,
		Gravity: gravity,
		InitialX: X,
		InitialY: Y,
		InitialSpeedX: math.Abs(speed * math.Cos(angle*math.Pi/180)),
		InitialSpeedY: math.Abs(speed * math.Sin(angle*math.Pi/180)),
		WindAcceleration: wind,
	}
}

func (p *Parabolic)PositionXatT(T float64) float64 {
	return (p.InitialSpeedX * T) + ((p.WindAcceleration * math.Pow(T, 2)) / 2)
}

func (p *Parabolic)PositionYatT(T float64) float64 {
	return (p.InitialSpeedY * T) - ((p.Gravity * math.Pow(T, 2)) / 2)
}


// ini sampe bawah gk berguna sih
func (p *Parabolic)TatPositionX(X float64) float64 {
	return X / p.InitialSpeedX
}

func (p *Parabolic) MaxX () float64 {
	return 2*p.InitialSpeed*math.Sin(p.InitialAngel)*p.InitialSpeed*math.Cos(p.InitialAngel) / p.Gravity
}

func (p *Parabolic) MaxY () float64 {
	return math.Pow(p.InitialSpeed*math.Sin(p.InitialAngel), 2) / (2 * p.Gravity)
}

func (p *Parabolic) TimeInAir () float64 {
	return 2*p.InitialSpeed*math.Sin(p.InitialAngel)/p.Gravity
}