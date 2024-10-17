package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name   string `required:"true" max:"10"`
	Addres string `required:"true" max:"10"`
	Age    string `required:"true" max:"10"`
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			fmt.Printf("%s is required\n", f.Name)
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}

		}
	}
	return result
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println(valueField.Tag.Get("max"))
		fmt.Println(valueField.Tag.Get("required"))
	}
}

func main() {
	readField(Sample{
		Name: "John",
	})

	readField(Person{
		Name:   "John",
		Age:    "20",
		Addres: "Test",
	})

	person := Person{
		Name:   "John",
		Age:    "20",
		Addres: "",
	}
	fmt.Println(isValid(person))
}
