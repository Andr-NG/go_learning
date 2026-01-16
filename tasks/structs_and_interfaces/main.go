package main

import "fmt"

func main() {

	// user := User{
	// 	Email:    "user@mail.com",
	// 	ID:       12345,
	// 	IsActive: true,
	// }

	// PrintUser(user)
	// fmt.Println("_________________")
	// ToggleActiveByValue(user)
	// PrintUser(user)
	// fmt.Println("_________________")
	// ToggleActiveByPointer(&user)
	// PrintUser(user)

	// Constructor functions
	// product, err := NewProduct("Case", -1)
	// if err != nil {
	//     fmt.Println(err)
	// } else {
	//     fmt.Println(*product)
	// }

	// Embedded structs/Composition

	// Customer := Customer{
	// 	Name: "Sergio",
	// 	City: "Paris", // Customer.City overrides Address.City
	// 	Address: Address{
	// 		City:    "Madrid",
	// 		Country: "Spain",
	// 	},
	// }
	// fmt.Println("Customer name is", Customer.Name)
	// fmt.Println("Customer city is", Customer.City)
	// fmt.Println("Overrided Address city is", Customer.Address.City)
	// fmt.Println("Country is", Customer.Country) // field promoted from Address to Customer

	// // Print full customer info
	// fmt.Println("\nFull Customer Info:")
	// fmt.Printf("%+v\n", Customer) //  %+v is a representation of the value with field names (for structs)

	// Methods
	// user := Account{
	//     Owner: "Andrey",
	//     Balance: 10000,
	// }
	// amount := 900
	// fmt.Println("Withdrawing", amount)
	// err := user.Withdraw(amount)
	// if err != nil {
	//     fmt.Println(err)
	// }

	// Interfaces
	// logger := ConsoleLogger{}
	// Sum(5, 6, logger)

	// fileLogger := FileLogger{}
	// Sum(10, 10, &fileLogger)

	// fmt.Println("Result logged by FileLogger:",  fileLogger.Logs)

	// Polymorphim without inheritance
	// result := GetTotalArea()
	// fmt.Println(result)


	
	repo := &InMemoryUserRepo{
        UserList: make(map[int]MyUser), // 1. Initialize the repository and the map inside it
    }
    
    service := UserService{
        repo: repo, // 2. Inject the repository into the service
    }

    user := MyUser{
        Email: "go@go.com",
        ID:    123,
    }

	fmt.Println("Registering a new user...")
	err := service.Register(user)
	if err != nil {
		fmt.Println("Registering a new user failed:", err)
	} else {
		fmt.Println("User registered", user)
	}

	fmt.Println("Fetching a user by ID...")
	user, err1 := service.GetUserById(13)
	if err1 != nil {
		fmt.Println("Fetching a user failed:", err1)
		
	} else {
		fmt.Println("User found", user)
	}

}
