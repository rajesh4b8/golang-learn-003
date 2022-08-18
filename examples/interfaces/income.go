package main

import (
	"fmt"
	"strings"
)

type Income interface {
	source() string // method to get the source of income
	calculate() int // method to calculate and return the total income
}

// fixedBilling implementing Income
type fixedBilling struct {
	projectName string
	bidAmount   int
}

func (f fixedBilling) source() string {
	return f.projectName
}

func (f fixedBilling) calculate() int {
	return f.bidAmount
}

type timeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (t timeAndMaterial) source() string {
	return t.projectName
}

func (t timeAndMaterial) calculate() int {
	return t.noOfHours * t.hourlyRate
}

type advertisement struct {
	addName      string
	noOfClick    int
	ratePerClick int
	company      string
}

// Implment this method `source() string` for Advertisement
func (a advertisement) source() string {
	return a.addName
}

// Implment this method `calculate() int` for Advertisement
func (a advertisement) calculate() int {
	return a.noOfClick * a.ratePerClick
}

// Assignment create a income stream for rental income
// Define a new strut with propertyName, monthlyRent, calculate rent for one year

func main() {
	var p1 Income = fixedBilling{"prj 1", 2000}
	var p2 Income = timeAndMaterial{"prj 2", 10, 40}
	p3 := timeAndMaterial{"prj 3", 2, 100}
	p4 := timeAndMaterial{"prj 4", 4, 100}
	p5 := advertisement{"add 1", 10000, 1, "Google"}

	// calculate net income
	// income := p1.bidAmount + p2.calculateIncome() + p3.calculateIncome()
	// fmt.Println("Total income", income)
	calculateNetIncome([]Income{p1, p2, p3, p4, p5})

	var contents []byte = make([]byte, 10)
	r := strings.NewReader("12345")
	r.Read(contents)
	fmt.Println("contents:", string(contents))
}

func calculateNetIncome(steams []Income) {
	var netIncome int
	for _, s := range steams {
		fmt.Printf("Income from %v is %v\n", s.source(), s.calculate())
		netIncome += s.calculate()
	}

	fmt.Println("Total net income:", netIncome)
}
