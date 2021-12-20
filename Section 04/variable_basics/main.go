package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)

	var name = "Frengky"
	fmt.Println("Name: ", name)

	var umur = 30
	_ = umur

	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 1991
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64 = 60000
		firstName string  = "frengky"
		gender    bool    = true
	)

	fmt.Println(salary, firstName, gender)

	var i, j int
	i, j = 5, 8
	i, j = j, i

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
