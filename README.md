# Modeling and Simulation Repository

This is my personal repository for Modeling and Simulation course task.

## Installment
Install Golang properly
In your terminal, run:

```
~ go get github.com/pamungkaski/Modeling-and-Simulation
~ go mod vendor
```

## Random Walk
![](https://raw.githubusercontent.com/pamungkaski/Modeling-and-Simulation/master/random_walk.png)

In your terminal, run:

```
go run app/rng/main.go
```

The Parameters that can be edited are:
- Percentage of each move
- Plane dimension

to change it, edit `app/rng/main.go`

the plotting is saved in `random_walk.png` or `random_walk.svg`
## Parabolic
![](https://raw.githubusercontent.com/pamungkaski/Modeling-and-Simulation/master/parabolic.png)

In your terminal, run:

```
go run app/parabolic/main.go
```

Maximum Height :  3.826530612244897

Maximum Range :  8.83699391616774

Postion at t = 0.2:  x=1.0000000000000002 y=1.536050807568877

Postion at t = 0.3:  x=1.5000000000000002 y=2.1570762113533157

How long the bullet in the air:  1.767398783233548s


the plotting saved to `parabolic.svg` or `parabolic.png`

