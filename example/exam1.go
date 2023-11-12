package example

import (
	"fmt"
	"reflect"
	"strconv"
)

// main
func exam1() {
	// // -- input print n output
	// fmt.Println("enter your name: ")
	// textInput := ""
	// fmt.Scan(&textInput)
	// fmt.Println("" + textInput)

	age := 2

	fmt.Println(reflect.TypeOf(age))

	// -- convert -- 
	height := "100"

	var h int

	// -- str to int number --
	h, err := strconv.Atoi(height)
	if err != nil {
		fmt.Println(err)
	}

	water := "12.40"
	// -- str to float number --
	fwater, err := strconv.ParseFloat(water, 64)

	fmt.Println(h)
	fmt.Println(fwater)

	txtShow := ""

	// -- number to str --
	txtShow = fmt.Sprintf("water = %f , height = %d", fwater, h)

	fmt.Println(txtShow)

	// -- loop --
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	score := 80
	// -- switch case --

	switch score {
	case 80:
		fmt.Println("A")
	case 70:
		fmt.Println("B")
	default:
		fmt.Println("unknow")
	}

	// -- if else --
	if score >= 80 {
		fmt.Println("A")
	} else if score > 70 && score < 80 {
		fmt.Println("B")
	} else {
		fmt.Println("unknow")
	}
}
