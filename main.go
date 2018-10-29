package main

import(
  "fmt"
  "./example/case1"
  "./example/case2"
)

func main(){
    fmt.Println("Hello Face Group")
    fmt.Println(case1.HelloCase1())
    fmt.Println(case2.HelloCase2())
}
