package main

import "fmt"

/*
Logger → abstraction / interface / “promise of behavior”

ConsoleLogger/FileLogger → concrete implementation (worker) / “actual tool”

Passing a logger to ProcessOrder → dependency injection
*/


// Defining the interface
type Logger interface {
	Log(message string)
}

// Defining workers that satisfy the Logger interface
type ConsoleLogger struct{}
type FileLogger struct{
    Logs []string
}

// Tying the Log method to these workers, so they be used as arguments to functions that
// expect Logger
func (c ConsoleLogger) Log(message string) {
	fmt.Println("Logged message:", message)
}
func (f *FileLogger) Log(message string) {
    f.Logs = append(f.Logs, message)
}


// Passing the Logger interface to describe what this function can do
// Passing the interface and NOT the worker, so we can pass any workers that satisfy Logger
func Sum(num1, num2 int, l Logger){
    l.Log(fmt.Sprintf("First number is %d", num1))
    l.Log(fmt.Sprintf("Second number is %d", num2))

    sum := num1 + num2
    l.Log(fmt.Sprintf("Sum is %d", sum))
}   

