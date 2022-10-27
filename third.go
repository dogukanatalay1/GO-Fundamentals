package main

import "fmt"

type person struct {
  name string
  age int
}

func initPerson() *person {
  m:= person{name:"noname", age:5}
  fmt.Printf("initPerson --> %p\n", &m)
  return &m
}

// stack doesn't need a garbage collector, its self cleaning
// but heap needs garbage collector to be cleaned

// giving too much work to garbage collector can 
// affect performance

func main() {
  // fmt.Println(initPerson()) 
  fmt.Printf("main --> %p\n", initPerson())
}


