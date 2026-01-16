package main
import (
	"math"
	"fmt"
)

/* Defining an interface an interface tells a function to focus on what an object can do,
rather than what the object is.
*/
type Shape interface {
	area() float64
}

// Defining a structure that has radius
type Cirlce struct{
	radius float64
}

// Defining a receiver function which explains what area() does and ties it to Circle struct 
func (c Cirlce) area() float64 {
	return math.Pi * c.radius * c.radius
}


/* Defining a "global" function to ger area() of whatever Share there will be
Example of polymorphism.
*/

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Cirlce{radius: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))

}