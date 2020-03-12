package main

import "fmt"

func print() {
	fmt.Println("abc")

}

func main() {
	var weihua = "ddad"
	var kg = weihua
	fmt.Println(weihua, kg)
	fmt.Printf("kg type=%T\n", kg)
	fmt.Printf("my name=%s\n", kg)
	go print()
}
