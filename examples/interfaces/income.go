package main

import "fmt"

type Income interface {
	source() string
	calculate() int
}

type FixedBilling struct {
	projectName string
	bidAmount   int
}

func (f FixedBilling) source() string {
	return f.projectName
}

func (f FixedBilling) calculate() int {
	return f.bidAmount
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (t TimeAndMaterial) source() string {
	return t.projectName
}

func (t TimeAndMaterial) calculate() int {
	return t.noOfHours * t.hourlyRate
}

type Advertisement struct {
	addName      string
	noOfClick    int
	ratePerClick int
}

// Implment this method `source() string` for Advertisement
func (a Advertisement) source() string {
	return a.addName
}

// Implment this method `calculate() int` for Advertisement
func (a Advertisement) calculate() int {
	return a.noOfClick * a.ratePerClick
}

// Assignment create a income stream for rental income
// Define a new strut with propertyName, monthlyRent, calculate rent for one year
type RentalIncome struct {
	propertyName string
	monthlyRent  int
}

func (r RentalIncome) source() string {
	return r.propertyName
}
func (r RentalIncome) calculate() int {
	return r.monthlyRent * 12
}

func main() {
	p1 := FixedBilling{"prj 1", 2000}
	p2 := TimeAndMaterial{"prj 2", 10, 40}
	p3 := TimeAndMaterial{"prj 3", 2, 100}
	p4 := TimeAndMaterial{"prj 4", 4, 100}
	p5 := Advertisement{"add 1", 10000, 1}
	p6 := RentalIncome{"Newyork", 20000}

	// calculate net income
	// income := p1.bidAmount + p2.calculateIncome() + p3.calculateIncome()
	// fmt.Println("Total income", income)
	calculateNetIncome([]Income{p1, p2, p3, p4, p5, p6})
}

func calculateNetIncome(steams []Income) {
	var netIncome int
	for _, s := range steams {
		fmt.Printf("Income from %v is %v\n", s.source(), s.calculate())
		netIncome += s.calculate()
	}

	fmt.Println("Total net income:", netIncome)
}
