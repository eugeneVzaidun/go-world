package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// define struct for the helper
type Helper struct {
	Name string
	Age  int
}

// define a function that returns a helper
func NewHelper(name string, age int) *Helper {
	return &Helper{
		Name: name,
		Age:  age,
	}
}


// define a function that prints information about the helper
func (h *Helper) PrintInfo() {
	fmt.Printf("Name: %s\nAge: %d\n", h.Name, h.Age)
}

// define a function that prints age of the helper
func (h *Helper) PrintAge() {
	fmt.Printf("Age: %d\n", h.Age)
}

// define a function that prints name of the helper
func (h *Helper) PrintName() {
	fmt.Printf("Name: %s\n", h.Name)
}

// define funcion that return slice of helpers with given age
func GetHelpersByAge(helpers []*Helper, age int) []*Helper {
	var result []*Helper
	for _, helper := range helpers {
		if helper.Age == age {
			result = append(result, helper)
		}
	}
	return result
}


func ReadHelpers() []*Helper {
	// read file with ioutil.ReadFile and return slice of helpers
	file, err := ioutil.ReadFile("helpers.txt")
	if err != nil {
		fmt.Println(err)
	}
	var helpers []*Helper
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			helper := strings.Split(line, " ")
			ageInt, _ := strconv.Atoi(helper[1])
			helpers = append(helpers, NewHelper(helper[0], ageInt))
		}
	}
	return helpers
}