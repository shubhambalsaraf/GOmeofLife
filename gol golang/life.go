package main

import (
	"fmt"
	"math/rand"
	"time"
)

//live function generates random array of 1's and 0's
func live(rows int, cols int, gens int) {
	var calc int
	calc = rows * cols
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1
	var now [1050]int
	for i := 0; i < calc; i++ {
		now[i] = (rand.Intn(max-min+1) + min)
	}
	//fmt.Println(now)
	life(now, rows, cols, gens)
}

func life(now [1050]int, rows int, cols int, gens int) {

	for gens >= 1 { //this loop will exit when generations become less than 1
		time.Sleep(1)
		homeScreen()
		fmt.Println("Generation", gens)

		var newarr [1001]int
		var neighbours int
		var s []string
		//This loop  prints the current generation
		for c := 0; c <= len(now)-1; c++ {
			if now[c] == 1 {
				s = append(s, "o")
				//fmt.Println("o")
			} else {
				s = append(s, " ")
				//fmt.Println("-")
			}

		}
		fmt.Println(s)
		//Next generation calculation
		//Check number of living neighbour
		for c := 1; c <=1000; c++ {
        neighbours = 0
        if c-1 >= 1{
            neighbours += now[c-1]
        }
        if c+1 <= 1000{
            neighbours += now[c+1]
        }
        if c-rows-1 >= 1{
            neighbours += now[c-rows-1]
        }
        if c-rows >= 1 && c-rows <=1000{
            neighbours += now[c-rows]
        }
        if c-rows+1 >= 1  && c-rows+1 <=1000{
            neighbours += now[c-rows+1]
        }
        if c+rows-1 >=1 && c+rows-1 <= 1000{
            neighbours += now[c+rows-1]
        }
        if c+rows <= 1000{
            neighbours += now[c+rows]
        }
        if c+51 <= 1000{
            neighbours += now[c+51]
        }
			newarr[c] = now[c]
			//Update status based on living neighbour
			if now[c] == 0 {
				if neighbours == 3 {
					newarr[c] = 1
				}
				if neighbours == 3 {
					newarr[c] = 1
				} else if neighbours == 2 || neighbours == 3 {
					newarr[c] = 1
				}
			}

		}
		//Recursive call
		gens = gens - 1
		live(newarr[1000], cols, gens)
	}
}

func homeScreen() {
	fmt.Println("\033[1;1H")
}

func clearScreen() {
	fmt.Println("\033[2J")
}

func main() {

	time.Sleep(1)
	clearScreen()
	live(50, 20, 6)
}
