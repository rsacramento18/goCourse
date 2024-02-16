package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {
	ricardo := data.Instructor{Id: 3, FirstName: "Ricardo"}
	ricardo.LastName = "Sacramento"
	// fmt.Print(ricardo.Print())

	// kyle := data.NewInstructor("Kyle", "Simpson")

	goCourse := data.Course{
		Id: 2, Name: "Go Fundamentals",
		Instructor: ricardo,
	}

	// fmt.Printf("%v", kyle)

	// fmt.Printf("%v\n", goCourse)

	swiftWS := data.NewWorkshop("Swift with iOS", ricardo)
	// fmt.Printf("%v", swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
