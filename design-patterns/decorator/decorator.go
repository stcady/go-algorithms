package main

import "fmt"

// IProcess interface
type IProcess interface {
	process()
}

// ProcessClass struct
type ProcessClass struct{}

// process method for the ProcessClass class
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// process method for the ProcessDecorator class
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("Process Decorator process and ")
		decorator.processInstance.process()
	}
}

func main() {
	process := &ProcessClass{}
	decorator := &ProcessDecorator{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
