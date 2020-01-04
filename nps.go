package main

import "fmt"

func main() {
	score1 := 10
	score2 := 10
	score3 := 10

	promoters := 0
	detractors := 0

	if score1 >=9{
		promoters = promoters + 1
	}
	if score1 >=6{
		detractors = detractors + 1
	}
	if score2 >=9{
		promoters = promoters + 1
	}


	if score2 >=6{
		detractors = detractors + 1
	}
	if score3 >=9{
		promoters = promoters + 1
	}
	if score3 >=9{
		promoters = promoters + 1
	}
	nps := (promoters - detractors ) * 100 / 3
fmt.Print(nps)



}

