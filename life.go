package main

import (
    "fmt"
    "math/rand"
    "time"
)

func live(rows int, cols int, gens int){
  var calc int
  calc = rows * cols
  rand.Seed(time.Now().UnixNano())
  min := 0
  max := 1
  var now[100]int
    for i:=0; i < calc; i++{
    now[i] = (rand.Intn(max - min + 1) + min)
    }
  fmt.Println(now)
  life(now , rows, cols, gens )
}

func life(now [100]int, rows int, cols int, gens int){
  if gens < 1{
    //break
    //return 0 //return not working
    //return 0, errors.New("math: square root of negative number")
    return 0, fmt.Errorf("Generation is %g", gens)
  }
  
  time.Sleep(1)
  homeScreen()

  fmt.Println("Generation", gens)

var newarr[100] int
var neighbour int
  for c:=1; c<=len(now);c++{
    if now[c] ==1{
      fmt.Println("o")
    } else {
      fmt.Println(" ")
    }
    if c%rows == 0{
      fmt.Println("\n")
    }
  }
  for c:=1; c<=1000;c++{
    neighbour = 0
    if c-1 >=1{
      neighbour += now[c-1]
    }
    if c+1 <= 1000{
      neighbour += now[c+1]
    }
    if c-rows-1 >=1{
      neighbour += now[c-rows-1]
    }
    if c-rows >=1 && c-rows <= 1000{
      neighbour += now[c-rows]
    }
    if c-rows+1 >= 1 && c-rows+1 <= 1000{
      neighbour += now[c-rows+1]
    }
    if c+rows-1 >=1 && c+rows-1 <=1000{
      neighbour += now[c+rows-1]
    }
    if c+rows <= 1000{
      neighbour += now[c+rows]
    }
    if c+51 <= 1000{
      neighbour += now[c+51]
    }
    newarr[c]= now[c]
    if now[c] == 0{
      if neighbour == 3 {
        newarr[c] =1
      } 
      if neighbour ==3{
        newarr[c] = 1
      } else if neighbour==2 || neighbour== 3 {
        newarr[c] = 1
      }
    }

  }

  gens = gens -1
  live(newarr , cols, gens)
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
  live(5,2,2)
}