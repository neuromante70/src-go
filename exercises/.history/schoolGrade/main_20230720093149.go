/*
We will build a school grade tracking service

Concepts and models to create in your program
1. Create a student that will have an id and name as the fields - id is int and name is string
2. Create constants for subjects such as english, math, science and history
3. Delete student by ID

A student can have grades for multiple subjects - so it is a one to many relation - grades will be kept in a list

Functionalities for grades
1. Add a grade for a student
2. Get the average score across all subjects for a student id
3. Get average score for a particular subject for a student id
4. Get the subject with the lowest score for a student id
5. Get the subject with the highest score for a student id
6. Get all grades for a student id


Hint: You will have 2 maps in this exercise
*/

package main

import (
	"BASICS/exercises/schoolGrade/grades"
	// "errors"
	"fmt"
	// "strconv"
)

var prt = fmt.Println

func main() {
	s1 := grades.Student{1, "Carl"}
	s2 := grades.Student{2, "Janis"}
	s3 := grades.Student{3, "Luise"}
	s4 := grades.Student{4, "John"}
	s1.AddStudent(s1)
	s2.AddStudent(s2)
	s3.AddStudent(s3)
	s4.AddStudent(s4)
	_, err := s4.FindStudent(s4.d)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s4.name, "present in database")
	}
	s4.DeleteStudent(s4.id)

	ex1 := Exam{
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
	prt("lowest math grade for Carl:") // 1
	prtAvg(s1.lowGradeBySbj(1))
	prt("highest math grade for Carl:") // 3
	prtAvg(s1.highGradeBySbj(1))
	prt("highest math grade for Luise:") // no exams in math
	prtAvg(s3.highGradeBySbj(1))
	prt("All Carl grades:") // 4, 1, 3, 3
	prtAvg(s1.allGrade())
	// result, err := s1.allGrade()
	// if err != nil {
	// 	prt(err)
	// } else {
	// 	prt(result)
	// }
}
