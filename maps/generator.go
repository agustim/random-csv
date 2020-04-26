package maps

import (
	"fmt"
	"math/rand"

	"github.com/Pallinder/go-randomdata"
)

func (estructura *MapSlice) Header() string {
	var header string = ""

	for _, value := range *estructura {
		if header != "" {
			header = header + ","
		}
		header = header + fmt.Sprintf("%q", value.Key)
	}
	return header + "\n"
}

func (estructura *MapSlice) Line() string {
	var line string = ""

	for _, value := range *estructura {
		if line != "" {
			line = line + ","
		}
		line = line + FieldRandom(fmt.Sprintf("%v", value.Value))
	}
	return line + "\n"

}

func (estructura *MapSlice) Generate(iterations int) string {
	var ret string
	ret = estructura.Header()
	for i := 0; i < iterations; i++ {
		ret = ret + estructura.Line()
	}
	return ret
}
func (estructura *MapSlice) GeneratePrint(iterations int) {

	fmt.Print(estructura.Header())
	for i := 0; i < iterations; i++ {
		fmt.Print(estructura.Line())
	}
}

func FieldRandom(typeField string) string {

	var r string = ""
	var randSex int

	if rand.Intn(2) == 0 {
		randSex = randomdata.Male
	} else {
		randSex = randomdata.Female
	}
	switch typeField {
	case "name":
		r = randomdata.FirstName(randSex)
	case "male-name":
		r = randomdata.FirstName(randomdata.Male)
	case "female-name":
		r = randomdata.FirstName(randomdata.Female)
	case "lastname":
		r = randomdata.LastName()
	case "email":
		r = randomdata.Email()
	case "phone":
		r = randomdata.PhoneNumber()
	}
	return fmt.Sprintf("%q", r)
}
