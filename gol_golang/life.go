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
  t := [1000] bool{false}

	for i := 0; i < calc; i++ {
		now[i] = (rand.Intn(max-min+1) + min)
    if now[i]==1{
    t[i]= true
    } else {
    t[i]= false
    }
	} 
	life(now, t, rows, gens)
}

func life(now [1050] int,t [1000]bool, rows int,  gens int){
  for gens >=1{
      time.Sleep(50)
      homeScreen()
      var s []string
		  var neighbours int
      var newarr [1050]int
      newbool := [1000] bool{false}   
//prints current generation
  	  fmt.Println("----------------Generation-------------", gens)
      for c:=1;c<len(t);c++{
        if t[c]{
          s = append(s, "O")
        } else{
          s = append(s, " ")
        }
        //if c%rows==0{
        //  s = append(s, "\n")
        //}
      }
      fmt.Println(s)
//next generation calcultaion
      for c:=1;c<=len(t);c++{
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
 			  if now[c] == 0 {
				  if neighbours == 3 {
					  newarr[c] = 1
				  } 
        } else if now[c] ==1 {
          if neighbours == 2 || neighbours == 3 {
					  newarr[c] = 1
				  }
        }
    
      for i := 0; i < 1000; i++ {
        if newarr[i]==1{
        newbool[i]= true
        } else {
        newbool[i]= false
        }
	    }  
      //Recursive call
      gens = gens - 1
      if gens <=1{
        break
      }
      life(newarr, newbool , rows, gens)
    }
  }
}

func homeScreen() {
	fmt.Println("\033[1;1H")
}
func clearScreen() {
	fmt.Println("\033[2J")
}
func main() {
  time.Sleep(50)
  clearScreen()
	live(50, 20, 8)
}