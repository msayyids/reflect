package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	type Avenger struct {
		Id   int    `required:"true" max:"10"`
		Name string `required:"true"`
		// Email     string `required:"true"`
		Age       int    `required:"true" min:"18" max:"50"`
		SuperHero string `required:"true" minLen:"4" maxLen:"25"`
	}
	//
	newAvenger := Avenger{
		1,
		"Sayyid",
		20,
		"d",
	}
	err := ValidateStructs(newAvenger)
	fmt.Println(err)
}

func ValidateStructs(s interface{}) error {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(s).Field(i).Interface()
			if value == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}

		if field.Tag.Get("min") != "" && field.Tag.Get("max") != "" {
			min, _ := strconv.Atoi(field.Tag.Get("min"))
			max, _ := strconv.Atoi(field.Tag.Get("max"))
			value := reflect.ValueOf(s).Field(i).Interface().(int)
			if value < min || value > max {
				return fmt.Errorf("minimal %s harus %d dan maksimal %s harus %d untuk melawan thanos", field.Name, min, field.Name, max)
			}
		}

		if field.Tag.Get("minLen") != "" && field.Tag.Get("maxLen") != "" {
			min, _ := strconv.Atoi(field.Tag.Get("min"))
			max, _ := strconv.Atoi(field.Tag.Get("max"))
			value := reflect.ValueOf(s).Field(i).Interface().(string)
			len := len(value)
			if len < min || len > max {
				return fmt.Errorf("nama %s harus minimal %d char dan max %d char", field.Name, min, max)
			}
		}

	}
	return nil
}
