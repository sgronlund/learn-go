package main

import "fmt"

type Customer struct {
	Name string
	Loan int
}

func (c *Customer) increaseLoan() {
	c.Loan += 100
}

func main() {
	customer := &Customer{"Bengt", 500}
	customer.increaseLoan()
	fmt.Println(customer.Loan)
}
