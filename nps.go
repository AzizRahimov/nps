package main

import "fmt"

func main() {
	score1 := 10
	score2 := 10
	score3 := 10

	promoters := 0
	detractors := 0






	promotersLowerBound := 9
	if score1 >= promotersLowerBound {
		promoters = promoters + 1
	}
	detractors = 6
	if score1 >= detractors {
		detractors = detractors + 1
	}
	if score2 >= promotersLowerBound {
		promoters = promoters + 1
	}


	if score2 >= detractors {
		detractors = detractors + 1
	}
	if score3 >= promotersLowerBound {
		promoters = promoters + 1
	}
	if score3 >= promotersLowerBound {
		promoters = promoters + 1
	}
	nps := (promoters - detractors ) * 100 / 3
fmt.Print(nps)



}

