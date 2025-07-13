package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Course struct {
	Name   string
	Weight float32
}

type Student struct {
	Name   string
	Grades map[Course]float32
}

func (c Course) String() string {
	return fmt.Sprintf("%s (Weight: %.2f)", c.Name, c.Weight)
}

func (s Student) String() string {
	var sb strings.Builder
	sb.WriteString("====================================\n")
	sb.WriteString(fmt.Sprintf("Student Name : %s\n", s.Name))
	sb.WriteString("------------------------------------\n")
	sb.WriteString("Grades:\n")
	if len(s.Grades) == 0 {
		sb.WriteString("  (No grades available)\n")
	} else {
		for course, grade := range s.Grades {
			sb.WriteString(fmt.Sprintf("  • %-25s | Grade: %6.2f\n", course.String(), grade))
		}
	}
	sb.WriteString("====================================\n")
	return sb.String()
}

func ReadStudent() (name string, cnum uint8, success bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter number of courses you took: ")
	numCoursesStr, _ := reader.ReadString('\n')
	numCoursesStr = strings.TrimSpace(numCoursesStr)
	numCoursesInt, err := strconv.Atoi(numCoursesStr)

	if err != nil {
		fmt.Println("Invalid number of courses.")
		return "", 0, false
	}

	cnum = uint8(numCoursesInt)
	success = true

	return
}

func ReadCourses(cnum uint8) map[Course]float32 {
	var cg map[Course]float32

	for i := 1; i <= int(cnum); i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter name for course #%d: ", i)
		courseName, _ := reader.ReadString('\n')
		courseName = strings.TrimSpace(courseName)

		fmt.Printf("Enter weight for course #%d: ", i)
		weightStr, _ := reader.ReadString('\n')
		weightStr = strings.TrimSpace(weightStr)
		weight, err := strconv.ParseFloat(weightStr, 32)
		if err != nil {
			fmt.Println("Invalid weight. Please enter a valid number.")
			i--
			continue
		}

		fmt.Printf("Enter grade for course #%d: ", i)
		gradeStr, _ := reader.ReadString('\n')
		gradeStr = strings.TrimSpace(gradeStr)
		grade, err := strconv.ParseFloat(gradeStr, 32)
		if err != nil {
			fmt.Println("Invalid grade. Please enter a valid number.")
			i--
			continue
		}

		if cg == nil {
			cg = make(map[Course]float32)
		}

		course := Course{
			Name:   courseName,
			Weight: float32(weight),
		}

		cg[course] = float32(grade)
	}

	return cg
}

func CalculateGPA(courses map[Course]float32) float32 {
	var totalWeightedGrades float32
	var totalWeights float32

	for course, grade := range courses {
		totalWeightedGrades += grade * course.Weight
		totalWeights += course.Weight
	}

	if totalWeights == 0 {
		return 0
	}

	return totalWeightedGrades / totalWeights
}

func DisplayStudent(student *Student) {
	fmt.Println(student)
	fmt.Printf("GPA: %.2f\n", CalculateGPA(student.Grades))
}

func main() {
	name, cnum, ok := ReadStudent()
	if !ok {
		return
	}
	courses := ReadCourses(cnum)


	student := Student{
		Name:   name,
		Grades: courses,
	}

	DisplayStudent(&student)
}
