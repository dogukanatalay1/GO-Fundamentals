package main

import (
	"fmt"
	"os"
)

func main() {
   file, err := os.Open("dosyam.txt")
   if err != nil {
     fmt.Println(err)
   }
}
