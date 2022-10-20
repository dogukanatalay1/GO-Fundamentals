package main

import (
	"fmt"
	"os"
)

// func main() {

// file, _ := os.Open("dosyam.txt")   --err yerine _ koyunca yokluğunu kontrol etmeme gerek kalmadı
// -- --
//    file, err := os.Open("dosyam.txt")
//    if err != nil {
//      fmt.Println(err)
//    }
//    fmt.Println(file.Name())
// -- --
// myError := errors.New("bu bir hata!")
// fmt.Println(myError.Error())
// -- --
// i := -2
// if i < 0 {
//   return 0, fmt.Errorf("error! %g", i)
// }
// -- --
// _ , err := os.Open("abc.rar")
// if err != nil {
//   fmt.Println(err.Error())
// }

// }

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}
