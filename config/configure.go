package config

import (
	"fmt"
	"io/ioutil"
	"randomcsv/maps"
)

func Read(filename string) (maps.MapSlice, error) {

	var ei maps.MapSlice

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	err = ei.UnmarshalJSON(data)

	return ei, err

}
