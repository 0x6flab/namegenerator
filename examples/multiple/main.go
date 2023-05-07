package main

import (
	"fmt"

	"github.com/0x6flab/namegenerator"
)

var (
	FamilyName   = []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor"}
	MaleNames    = []string{"James", "John", "Robert", "Michael", "William", "David", "Richard", "Charles", "Joseph", "Thomas"}
	FemaleNames  = []string{"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen"}
	GeneralNames = []string{"James", "John", "Robert", "Michael", "William", "David", "Richard", "Charles", "Joseph", "Thomas", "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen"}
)

func main() {
	generator := namegenerator.NewNameGenerator("male")
	name := generator.GenerateNames(10)
	fmt.Println(name)
}
