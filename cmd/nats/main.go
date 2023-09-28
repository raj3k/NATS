package main

import "fmt"

const (
	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

func main() {
	fmt.Printf("The value of Red    is %v\n", Red)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Yellow is %v\n", Yellow)
	fmt.Printf("The value of Green  is %v\n", Green)
	fmt.Printf("The value of Blue   is %v\n", Blue)
	fmt.Printf("The value of Indigo is %v\n", Indigo)
	fmt.Printf("The value of Violet is %v\n", Violet)
}
