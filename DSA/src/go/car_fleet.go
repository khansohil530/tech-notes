package main

import "sort"

type Car struct {
	pos   int
	speed int
}

func constructCars(position []int, speed []int) []Car {
	cars := make([]Car, len(position))
	for i := range position {
		cars[i] = Car{position[i], speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos // reverse=True
	})
	return cars
}

func carFleet(target int, position []int, speed []int) int {
	var stk []float64
	cars := constructCars(position, speed)
	var l int
	for _, car := range cars {
		stk = append(stk, float64(target-car.pos)/float64(car.speed))
		l = len(stk)
		if len(stk) >= 2 && stk[l-1] <= stk[l-2] {
			stk = stk[:l-1]
		}
	}
	return len(stk)
}
