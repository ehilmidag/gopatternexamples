package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var mainOffice = Employee{
	Name:   "",
	Office: Adress{0, "sanayi", "turkey"},
}
var warehouseOffice = Employee{
	Name:   "",
	Office: Adress{0, "bos arazi", "turkey"},
}

type Adress struct {
	Suite        int
	Street, City string
}

type Employee struct {
	Name   string
	Office Adress
}

func (e *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	c := gob.NewEncoder(&b)
	_ = c.Encode(e)
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewMainWarehouseEmployee(name string, suite int) *Employee {
	return newEmployee(&warehouseOffice, name, suite)
}

func main() {
	hilmi := NewMainOfficeEmployee("Hilmi", 10)
	sami := NewMainWarehouseEmployee("Sami", 20)
	fmt.Println(hilmi)
	fmt.Println(sami)
}
