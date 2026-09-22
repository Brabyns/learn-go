package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type SalariesEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariesEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariesEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64 //Monthly base
	CommissionRate float64 //e.g., 0.05 for 5%
	SaleAmount     float64
}


func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SaleAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales: $%2f)", ce.Name, ce.BaseSalary, ce.CommissionRate*100, ce.SaleAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P){
	fmt.Printf(" -Processing: %s\n", employee)
}

func ProcessPayroll(employees []Payable){
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0

	for _, emp := range employees{
		PrintEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("Monthly Pay: $%2f\n", pay)
		totalPayroll += pay
	}

	fmt.Printf("\nToatl Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("--------------------------")
}


func main() {

	fmt.Println("Welcome to the Payroll Processor!")

	salEmp := SalariesEmployee{Name: "Alice Wonderland", AnnualSalary: 72000.00}
	hrEmp := HourlyEmployee{Name: "Bob The Builder", HourlyRate: 25.00, HoursWorked: 160.0}
	comEmp := CommissionEmployee{Name: "Chalie Chaplin", BaseSalary: 2000.00, CommissionRate: 0.10, SaleAmount: 15000.00}


	payrollList := []Payable{
		salEmp,
		hrEmp,
		comEmp,
		HourlyEmployee{Name: "Diana Prince", HourlyRate: 30.00, HoursWorked: 150.0},
	}

	ProcessPayroll(payrollList)


}
