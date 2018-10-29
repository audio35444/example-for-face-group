package main

import(
  "fmt"
  "./example/case1"
  "./example/case2"
  "./example/case3"
  ."./example/case4"
  myCase "./example/case5"
)

func main(){
    fmt.Println("Hello Face Group")
    fmt.Println(case1.HelloCase1())
    fmt.Println(case2.HelloCase2())
    fmt.Println(case3example.HelloCase3())
    fmt.Println(HelloCase4())
    fmt.Println(myCase.HelloCase5())
}
