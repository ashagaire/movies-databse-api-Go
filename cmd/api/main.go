package main

import (
	"fmt"
	internal "movies-api/internal/db"
)

func main(){

	fmt.Println("movies-api")
	internal.FromDb()

}