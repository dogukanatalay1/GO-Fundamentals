package main

import "fmt"

func main() {
  var mainArray = [3][3]int{{2,3,4},{5,6,7},{8,9}}

  sum := 0
  fmt.Println(sum)

  for _, childArray := range mainArray {
    for _, i := range childArray {
      if i > 3 {
        sum += i
      }
    }
  }

  fmt.Println(sum)
}

