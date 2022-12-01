package main

import "fmt"

const secondsInHour = 3600

func main() {
	fmt.Println("Hello Go world!")
	distance := 60.8 // distance in km
	fmt.Printf("The distance in miles %f \n", distance*0.621)
}
