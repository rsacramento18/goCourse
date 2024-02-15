package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {
	ricardo := data.Instructor{Id: 3, FirstName: "Ricardo"}
	ricardo.LastName = "Sacramento"
	fmt.Print(ricardo.Print())
}
