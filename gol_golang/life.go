package main

import (
	"fmt"
	"math/rand"
	"time"
)

//live function generates random array of 1's and 0's
func life(rows int, cols int, gens int) {
	var calc int
	calc = rows * cols
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1
	var now [1001]int
	t := [1000]bool{false}

	for i := 0; i < calc; i++ {
		now[i] = (rand.Intn(max-min+1) + min)
		if now[i] == 1 {
			t[i] = true
		} else {
			t[i] = false
		}
	}
	live(now, t, rows, gens)
}

func live(now [1001]int, t [1000]bool, rows int, gens int) {
	//best case check
	if gens < 1 {
		return
	}
	//delay between generations
	time.Sleep(190 * time.Millisecond)
	homeScreen()
	//assigning varibles
	var s []string
	var neighbours int
	var newarr [1001]int
	newbool := [1000]bool{false}
	//prints current generation
	fmt.Println("----------------Generation-------------", gens)
	for c := 1; c < len(t); c++ {
		if t[c] {
			s = append(s, "O")
		} else {
			s = append(s, " ")
		}
		//if c%rows==0{
		//  s = append(s, "\n")
		//}
	}
	fmt.Println(s)
	//next generation calcultaion
	for c := 1; c <= len(t); c++ {
		neighbours = 0
		if c-1 >= 1 {
			neighbours += now[c-1]
		}
		if c+1 <= 1000 {
			neighbours += now[c+1]
		}
		if c-rows-1 >= 1 {
			neighbours += now[c-rows-1]
		}
		if c-rows >= 1 && c-rows <= 1000 {
			neighbours += now[c-rows]
		}
		if c-rows+1 >= 1 && c-rows+1 <= 1000 {
			neighbours += now[c-rows+1]
		}
		if c+rows-1 >= 1 && c+rows-1 <= 1000 {
			neighbours += now[c+rows-1]
		}
		if c+rows <= 1000 {
			neighbours += now[c+rows]
		}
		if c+51 <= 1000 {
			neighbours += now[c+51]
		}
		//checking the game rules conditions
		newarr[c] = now[c]
		if now[c] == 0 {
			if neighbours == 3 {
				newarr[c] = 1
			}
		} else if now[c] == 1 {
			if neighbours == 2 || neighbours == 3 {
				newarr[c] = 1
			} else if now[c] == 1 {
				if neighbours >= 4 {
					newarr[c] = 0
				}
			}
		}

		for i := 0; i < 1000; i++ {
			if newarr[i] == 1 {
				newbool[i] = true
			} else {
				newbool[i] = false
			}
		}

	} //Recursive call
	live(newarr, newbool, rows, gens-1)
}

func homeScreen() {
	fmt.Println("\033[1;1H")
}
func clearScreen() {
	fmt.Println("\033[2J")
}
func main() {
	time.Sleep(90 * time.Millisecond)
	clearScreen()
	life(50, 20, 9) //life (rows, coloumn, generations < 10)
}
