package main

import (
	"fmt"
	"math"
)

func main() {
	question_1()
	question_2()
	question_3()
	question_4()
	question_5()
	question_6()
	question_7()
	question_8()
	question_9()
	question_10()
}

func question_1() {
	fmt.Println("Hello Go!")
}

func question_2() {
	var num1 = 5
	var num2 = 6
	fmt.Println("Sum", num1+num2)
}

func question_3() {
	num1, num2 := 5, 6
	fmt.Println("Sum", num1+num2)
}

func question_4() {
	var first string
	fmt.Println("Enter the first name")
	fmt.Scanln(&first)
	fmt.Print("First Name is ", first, "\n")
}

func question_5() {
	const PI = 3.14
	fmt.Println(PI)
}

func question_6() {
	var num = 16
	sqrtResult := math.Sqrt(float64(num))
	fmt.Println(sqrtResult)
}

func question_7() {
	var a, b int
	a, b = 5, 6
	a, b = b, a
	fmt.Println("Swapped number", a, b)
}

func question_8() {
	var (
		a = 5
		b = 7
	)
	fmt.Println(a, b)
}

func question_9() {
	var number1, number2 int
	fmt.Println("Enter the first Number")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second Number")
	fmt.Scanln(&number2)
	fmt.Print("product is ", (*&number1)*(*&number2), "\n")
	fmt.Print("product is ", number1*number2, "\n")
}

func question_10() {
	var number_int = 5
	floatNum := float64(number_int)
	fmt.Printf("%f\n", floatNum)
}
