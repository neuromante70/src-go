package main

import (
	"errors"
	"fmt"
	"strconv"
)

var prt = fmt.Println

/*
We will build a school grade treacking service

Concepts and models to create in your program
1. Create a student that will have an id and name as the fields - id is int and name is string
2. Create constants for subjects such as english, math, science and history
3. Delete student by ID

Functionalities for grades
1. Add a grade for a student
2. Get the average score across all subjects for a student id
3. Get average score for a particular subject for a student id
4. Get the subject with the lowest score for a student id
5. Get the subject with the highest score for a student id
6. Get all grades for a student id


Hint: You will have 2 maps in this exercise

*/

// 1. Create a student that will have an id and name as the fields - id is int and name is string
type student struct {
	id   int
	name string
}

// Hint: You will have 2 maps in this exercise
var Students = map[int]student{}

func (s student) StudentExist(idStd int) (student, bool) {
	student, ok := Students[idStd]
	return student, ok
}

func (s student) AddStudent(std student) bool {
	_, exist := s.StudentExist(std.id)
	if exist {
		fmt.Println("The student", std.name, "is already present in db")
		return false
	}
	Students[std.id] = std
	fmt.Println("The student", std.name, "wasn't present in db and now is correctly inserted")
	return true
}

func (s student) FindStudent(idStd int) (student, error) {
	std, exists := s.StudentExist(idStd)
	if exists {
		return std, nil
	}
	return student{}, errors.New("the ID student " + strconv.Itoa(idStd) + " is not in the database")
}

func (s student) DeleteStudent(idStd int) {
	_, ok := s.StudentExist(idStd)
	if ok {
		delete(Students, idStd)
		fmt.Println("student", idStd, "correctly deleted")
		return
	}
	fmt.Println("student non present, cannot delete it")
}

// 2. Create constants for subjects such as english, math, science and history
type subject int

const (
	english     subject = iota // 0
	mathematics                // 1
	science                    // 2
	history                    // 3
)

func ReturnSbj(idSub int) string {
	switch idSub {
	case 0:
		return "english"
	case 1:
		return "math"
	case 2:
		return "science"
	case 3:
		return "history"
	default:
		return "not found"
	}
	// return [...]string{"english", "math", "science", "history"}
}

// 3. Create a grade struct that will hold a subject name string and an integer score denoting the 'marks'
type exam struct {
	id      int // exam number
	idStd   int // student id
	subject int // exam subject
	grade   int // exam grade
}

// Hint: You will have 2 maps in this exercise: 2nd 
var Exams = map[int]exam{} // [int] is the exam number

// A student can have grades for multiple subjects - so it is a one to many relation - grades will be kept in a list
type grade int

// var Grades grade

// Here I decided to use the six-point system adopted by Phillips Academy at Andover.
const (
	Failure              grade = iota // 0
	LowFailure                        // 1
	LowPassButCertifying              // 2
	Satisfactory                      // 3
	Good                              // 4
	SuperiorHonor                     // 5
	OutstandingHighHonor              // 6
)

func ReturnGrade(idGrade int) string {
	switch idGrade {
	case 0:
		return "Failure"
	case 1:
		return "Low Failure"
	case 2:
		return "Low Pass, but certifying"
	case 3:
		return "Satisfactory"
	case 4:
		return "Good"
	case 5:
		return "Superior (Honor)"
	case 6:
		return "Outstanding (High Honor)"
	default:
		return "not found"
	}
}

func (e exam) extPrint() { //extended print Exams
	prt("id exam:", e.id)
	prt("student name:", Students[e.idStd].name)
	prt("subject:", ReturnSbj(e.subject))
	prt("grade:", ReturnGrade(e.grade))
}

// 2. Get the average score across all subjects for a student id
func (s student) avgGrade() (avg int, err error) {
	var sum int
	k := 0
	for _, v := range Exams {
		if v.idStd == s.id {
			sum += v.grade
			k++
		}
	}
	if k == 0 {
		return 0, errors.New("exams found for this student")
	}
	avg = sum / k
	return avg, nil
}

// 3. Get average score for a particular subject for a student id
func (s student) avgGradeBySbj(sbj int) (avg int, err error) {
	var sum int
	k := 0
	for _, v := range Exams {
		if v.idStd == s.id && v.subject == sbj {
			sum += v.grade
			k++
		}
	}
	if k == 0 {
		return 0, errors.New("exams not found for this matter")
	}
	avg = sum / k
	return avg, nil
}

func prtAvg(avg int, err error) {
	if err != nil {
		prt(err)
		return
	}
	prt(avg)
}

func main() {
	s1 := student{1, "Carl"}
	s2 := student{2, "Janis"}
	s3 := student{3, "Luise"}
	s4 := student{4, "John"}
	s1.AddStudent(s1)
	s2.AddStudent(s2)
	s3.AddStudent(s3)
	s3.AddStudent(s4)
	_, err := s4.FindStudent(s4.id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s4.name, "present in database")
	}
	s4.DeleteStudent(s4.id)

	ex1 := exam{
		id:      1,            // exam number
		idStd:   1,            // Carl
		subject: int(science), // science
		grade:   int(Good),    // good (4)
	}
	ex4 := exam{4, 1, 1, 1} // exNr 4, Carl, math, low failure
	ex5 := exam{5, 1, 1, 3} // exNr 5, Carl, math, satisfactory
	ex6 := exam{6, 1, 0, 3} // exNr 6, Carl, english, satisfactory
	ex2 := exam{2, 2, 1, 3} // exNr 2, Janis, math, satisfactory
	ex7 := exam{7, 2, 2, 4} // exNr 7, Janis, science, good
	ex8 := exam{8, 2, 2, 2} // exNr 8, Janis, science, low pass but certifying
	ex3 := exam{3, 3, 3, 6} // exNr 3, Luise, history, outstanding
	ex9 := exam{9, 3, 0, 2} // exNr 3, Luise, english, low pass but certifying

	Exams[ex1.id] = ex1
	Exams[ex2.id] = ex2
	Exams[ex3.id] = ex3
	Exams[ex4.id] = ex4
	Exams[ex5.id] = ex5
	Exams[ex6.id] = ex6
	Exams[ex7.id] = ex7
	Exams[ex8.id] = ex8
	Exams[ex9.id] = ex9
	ex1.extPrint()
	ex9.extPrint()

	// grd, _ := s1.avgGrade()
	prt("the grades average of student", Students[1], " is ") // Carl grades average = 2.66
	prtAvg(s1.avgGrade())
	prt("calculating Carl grades average for math:") // Carl math average = 2
	prtAvg(s1.avgGradeBySbj(1))
	prt("the grades average of student", Students[2], " is ") // Janis grades average = 3
	prtAvg(s2.avgGrade())
	prt("calculating Janis grades average for science:") // Janis grades average for science = 3
	prtAvg(s2.avgGradeBySbj(2))
	prt("the grades average of student", Students[3], " is ") // Luise grades average = 4
	prtAvg(s3.avgGrade())
	prt("calculating Luise grades average for history:") // Janis grades average for history = 6
	prtAvg(s3.avgGradeBySbj(3))
	prt("calculating Luise grades average for math:") // Janis grades average for math = no passed exam
	prtAvg(s3.avgGradeBySbj(1))
}
