package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary float64
}

type Person struct {
	Employee
	isMarried bool
	address   string
}

func newPerson(name string, age int, salary float64, isMarried bool, address string) Person {
	return Person{
		Employee: Employee{
			name:   name,
			age:    age,
			salary: salary,
		},
		isMarried: isMarried,
		address:   address,
	}
}

func (o *Person) changeAddress(newAddress string) {
	o.address = newAddress
	fmt.Println("Address inside method:", o.address)
}

func main() {
	Person1 := Person{
		Employee: Employee{
			name:   "John",
			age:    30,
			salary: 50000.50,
		},
		isMarried: false,
		address:   "123 Main St",
	}

	person2 := newPerson("Jane", 28, 60000.75, true, "789 Oak St")
	fmt.Println("Person2 Details:", person2)
	person2.address = "101 Pine St updated"
	fmt.Println("Person2 Details:", person2)

	fmt.Println("Person Details:", Person1)
	fmt.Println("Name:", Person1.Employee.name)
	fmt.Println("Name:", Person1.name)
	fmt.Println("Age:", Person1.age)
	fmt.Println("Salary:", Person1.salary)
	fmt.Println("Married:", Person1.isMarried)
	fmt.Println("Address:", Person1.address)
	Person1.changeAddress("456 Elm St")
	fmt.Println("Address outside method:", Person1.address)

}
