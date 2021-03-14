package main

import (
	"./database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
