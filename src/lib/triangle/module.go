package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-lib/src/input"
	s "github.com/whitman-colm/go-lib/src/other"
	"math"
)

///////////////////////////////////
//pythagorean theorem calculator://
////give it a & b, it finds c,/////
////give it b & c, it finds a!/////
///////////////////////////////////

var Version string = "\033[1;35m1.0.0\033[0m"

//Now with 53% more exportability!

func Startup() {
	//the function that's the "main menu" of sorts
	isDone := false
	doWhat := ""
	for isDone == false {
		//for loop so you don't have to come back here each time.
		//and a switch statment to force a responce.
		switch doWhat {
		case "a":
			findC()
		case "b":
			findAorB()
		case "c":
			isDone = true
		default:
			fmt.Println(c.CL + c.B01 + "pythagorean theorem calculator " + Version)
			s.Spacer(2)
			fmt.Println(c.B1 + "What do you need solved?" + c.X)
			fmt.Println(c.Y + "Enter one of the letters in red:")
			s.Spacer(1)
			fmt.Println(c.R + "{A}" + c.B0 + "  I have A & B, I need C!")
			fmt.Println(c.R + "{B}" + c.B0 + "  I have B & C, I need A!")
			fmt.Println(c.R + "{C}" + c.B0 + "  Nothing! Stop talking to me perv!")
			doWhat = i.ReturnString(c.M + ">>> " + c.B)
		}
	}
}

func findC() {
	//This the the part where we declare A&B
	fmt.Println(c.CL + c.B1 + "What is side A's size?")
	a := i.ReturnDouble(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is side B's size?")
	b := i.ReturnDouble(c.M + ">>> " + c.B)
	//Now we math.
	a2 := a * b
	b2 := b * b
	c2 := a2 + b2
	c1 := math.Sqrt(c2)
	//Defining a squared, b squared and c squared
	//along with solving for c (I used c1 as to prevent the possibility of problems with c, colors)

	//Now we show work
	fmt.Printf("%s%s%f (a) squared is %f, %f (b) squared is %f.\n", c.CL, c.B1, a, a2, b, b2)
	s.Spacer(1)
	fmt.Printf("%s%f+%f=%f.\n", c.B1, a2, b2, c2)
	fmt.Printf("%sThe square root of %f is %f.\n", c.B1, c2, c1)
	s.Spacer(2)
	fmt.Printf("%sc = %e\n", c.B00, c1)
}

func findAorB() {
	//This the the part where we declare A&B
	fmt.Println(c.CL + c.B1 + "What is side C's size?")
	c1 := i.ReturnDouble(c.M + ">>> " + c.B)
	fmt.Println(c.CL + c.B1 + "What is side B's size?")
	b := i.ReturnDouble(c.M + ">>> " + c.B)
	//Now we math.
	c2 := a * a
	b2 := b * b
	a2 := a2 - b2
	a := math.Sqrt(c2)
	//Defining a squared, b squared and c squared
	//along with solving for c (I used c1 as to prevent the possibility of problems with c, colors)

	//Now we show work
	fmt.Printf("%s%s%f (c) squared is %f, %f (b) squared is %f.\n", c.CL, c.B1, c1, c2, b, b2)
	s.Spacer(1)
	fmt.Printf("%s%f-%f=%f.\n", c.B1, c2, b2, a2)
	fmt.Printf("%sThe square root of %f is %f.\n", c.B1, a2, a)
	s.Spacer(2)
	fmt.Printf("%sa = %e\n", c.B00, a)
}
