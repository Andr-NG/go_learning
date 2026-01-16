package main

import "fmt"

// For-loop

func main() {
	// for i := range 11 {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i)
	// 	} else {
	// 		fmt.Println(i, "is odd")
	// 	}
	// }


	// Nested loops
	// for i := 1; i < 10; i++ {
	// 	for j :=1; j < 10; j++ {
	// 		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	// 	}
	// }

	//While-loop

	n := 1
	for n < 11 {
		fmt.Println(n)
		n++
	}

	// Loop control statements

	// for i := range 10 {
	// 	if i % 2 == 0 {
	// 		continue
	// 	} else if i == 9 {
	// 		break
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// } 

	//The goto statement

	// for i := 0; i <= 5; i++ {
	// 	if i == 3 {
	// 		goto break_point
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

// break_point:
// 	fmt.Println("i equals 3")


day := "Friday"
switch day {
case "Monday":
	fmt.Println("Shit! Here we go again!")
case "Tuesday", "Wednesday": 
	fmt.Println("Still alive? Goood! Carry on!")
case "Friday": 
	fmt.Println("LG!")
default:
	fmt.Println("Nothing happens here")
	}
}

