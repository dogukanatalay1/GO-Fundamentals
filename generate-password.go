package main

import (
	"fmt"
	"math/rand"
	"time"
)

// creating random password
var source = rand.NewSource(time.Now().UnixNano())

const _charset = "abcdefghjklmnoprstuvwxyzABCDEFGHJKLMNOPRSTUVWYXYZ0123456789"

func GeneratePassword(length int) string {
   x := make([]byte, length)
   for i := range x {
     x[i] = _charset[source.Int63() % int64(len(_charset))]
   }
   return string(x)
}

func main() {
  password := GeneratePassword(15)
  fmt.Println(password)
}




