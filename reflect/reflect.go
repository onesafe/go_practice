package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	contry string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)

		for i:=0; i<v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Printf("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId: 456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name: "Naveen",
		id: 453,
		address: "Coimbatore",
		salary: 90000,
		contry: "India",
	}
	createQuery(e)

	i := 9
	createQuery(i)
}