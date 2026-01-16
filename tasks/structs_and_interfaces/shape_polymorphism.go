// interfaces, polymorphism without inheritance

package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	legnth float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}


// Implementing Area() for all shapes
func (r Rectangle) Area() float64 {
	return r.legnth * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}


func GetTotalArea() float64 {

    // Initialising shapes
    r := Rectangle {
        width: 10,
        legnth: 5,
    }
    t := Triangle {
        base: 11,
        height: 2,
    }
    c := Circle {
        radius: 14.5,
    }

    // Creating a slice of Shapes and putting shapes into the slice, becase they satisfy the Shape interface
    s := []Shape{r, t, c}

    // Counting the sum
    var count float64
    for _, val := range s {
        count += val.Area()
    }
    return count
}