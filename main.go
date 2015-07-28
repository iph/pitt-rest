package main

import "fmt"

type LaundryRoom struct {
}

type ComputerLab struct {
	MachineGroups []MachineGroup
}

type MachineGroup struct {
	OS     string
	Open   int
	Closed int
}

func main() {
	courses := ParseCourses("2161", "AFRCNA")
	for _, course := range courses {
		fmt.Println(course)
	}
}
