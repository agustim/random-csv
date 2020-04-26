package main

import (
	"fmt"
	"randomcsv/config"
)

func main() {

	ed, err := config.Read("./example.json")
	if err != nil {
		fmt.Println(err)
	}

	ed.GeneratePrint(10000)
}
