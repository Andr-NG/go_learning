// embedded structs


package main


type Address struct {
    City string
    Country string
}


type Customer struct {
    Name string
    City string
    Address // embedded struct
}


