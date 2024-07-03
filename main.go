package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const NumDepartmentsPerApplicant = 3

type Person struct {
	Firstname string
	Lastname  string
}

type Department string

type Applicant struct {
	Person
	Deps []Department
	GPA  float64
}

func (p *Person) Fullname() string {
	return p.Firstname + " " + p.Lastname
}

func main() {
	applicants := loadInfo()

	var maxApplicants int
	fmt.Scan(&maxApplicants)

	applicantsByDepartment := distributeApplicants(applicants, maxApplicants)
	departments := getDepartments(applicantsByDepartment)

	// display result
	for _, department := range departments {
		fmt.Println(department)
		for _, applicant := range *applicantsByDepartment[department] {
			fmt.Printf("%s %.2f\n", applicant.Fullname(), applicant.GPA)
		}
		fmt.Println()
	}
}

func loadInfo() (applicants []Applicant) {
	file, err := os.Open("applicants.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	applicants = make([]Applicant, 0, 5)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var firstname, lastname string
		var dep1, dep2, dep3 Department
		var gpa float64

		// fetch applicant info
		fmt.Sscan(line, &firstname, &lastname, &gpa, &dep1, &dep2, &dep3)

		applicant := Applicant{
			Person{firstname, lastname},
			[]Department{dep1, dep2, dep3},
			gpa,
		}

		applicants = append(applicants, applicant)
	}

	return
}

func distributeApplicants(applicants []Applicant, maxApplicants int) map[Department]*[]Applicant {
	departments := make(map[Department]*[]Applicant, NumDepartmentsPerApplicant)

	sortApplicantsByGPA(&applicants)

	for i := 0; i < NumDepartmentsPerApplicant; i++ {
		for k, j := 0, len(applicants); j > 0; j-- {
			applicant := applicants[k]
			dep := applicant.Deps[i]
			if nil == departments[dep] {
				departments[dep] = new([]Applicant)
			}
			if len(*departments[dep]) < maxApplicants {
				*departments[dep] = append(*departments[dep], applicant)
				applicants = append(applicants[:k], applicants[k+1:]...)
			} else {
				k++
			}
		}
	}

	// sort applicants by GPA for each department
	for dep := range departments {
		sortApplicantsByGPA(departments[dep])
	}

	return departments
}

func getDepartments(applicantsByDepartment map[Department]*[]Applicant) []Department {
	departments := make([]Department, 0, len(applicantsByDepartment))
	for dep := range applicantsByDepartment {
		departments = append(departments, dep)
	}

	sort.Slice(departments, func(i, j int) bool {
		return departments[i] < departments[j]
	})

	return departments
}

func sortApplicantsByGPA(applicants *[]Applicant) {
	sort.Slice(*applicants, func(i, j int) bool {
		a, b := (*applicants)[i], (*applicants)[j]
		if a.GPA != b.GPA {
			return a.GPA > b.GPA
		}
		return a.Fullname() < b.Fullname()
	})
}
