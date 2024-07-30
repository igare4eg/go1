package main

import "fmt"

func main() {
	fmt.Println("Hello GO!")
	var age int8 = 29
	fmt.Println(age)
	var number float64 = 275.672
	fmt.Println(number)
	name := "Igor"
	fmt.Println(name)
	fmt.Println(len(name))
	var name2 string
	var age2 uint8
	fmt.Println("What's your name?")
	fmt.Scan(&name2)
	fmt.Println("Hello " + name2 + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age2)
	fmt.Println("Hello", age2)
	fmt.Println("You are " + fmt.Sprint(age2) + " years old!")
}
