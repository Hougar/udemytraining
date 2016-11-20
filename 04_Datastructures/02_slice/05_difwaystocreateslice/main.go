package main

import "fmt"

func main() {
	student := []string{}
	// short hand way to declare a slice  - othing got made so false address is nil
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	// var way -- nothing got made so false address is nil
	var etudiant []string
	var etudiants [][]string
	fmt.Println(etudiant)
	fmt.Println(etudiants)
	fmt.Println(etudiant == nil)
	// make way - creates the underlying array ! can reference the index
	// best way to make slices
	estudiante := make([]string, 35)
	estudiantes := make([][]string, 35)
	fmt.Println(estudiante)
	fmt.Println(estudiantes)
	fmt.Println(estudiante == nil)
}
