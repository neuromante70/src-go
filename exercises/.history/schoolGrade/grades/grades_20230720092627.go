package grades

import (
	"errors"
	"fmt"
	"strconv"
)

var prt = fmt.Println

// 1. Create a student that will have an id and name as the fields - id is int and name is string
type Student struct {
	Id   int
	Name string
}

// Hint: You will have 2 maps in this exercise: 1st map
var Students = map[int]Student{}

func (s Student) StudentExist(idStd int) (Student, bool) {
	Student, ok := Students[idStd]
	return Student, ok
}

func (s Student) AddStudent(std Student) bool {
	_, exist := s.StudentExist(std.id)
	if exist {
		fmt.Println("The student", std.name, "is already present in db")
		return false
	}
	Students[std.id] = std
	fmt.Println("The student", std.name, "wasn't present in db and now is correctly inserted")
	return true
}

func (s Student) FindStudent(idStd int) (Student, error) {
	std, exists := s.StudentExist(idStd)
	if exists {
		return std, nil
	}
	return Student{}, errors.New("the ID student " + strconv.Itoa(idStd) + " is not in the database")
}

// 3. Delete student by ID
func (s Student) DeleteStudent(idStd int) {
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
}

// 3. Create a grade struct that will hold a subject name string and an integer score denoting the 'marks'
type Exam struct {
	id      int // exam number
	idStd   int // student id
	subject int // exam subject
	grade   int // exam grade
}

// Hint: You will have 2 maps in this exercise: 2nd map
var Exams = map[int]Exam{} // [int] is the exam number

// A student can have grades for multiple subjects - so it is a one to many relation - grades will be kept in a list
type Grade int

// Here I decided to use the six-point system adopted by Phillips Academy at Andover.
const (
	Failure              Grade = iota // 0
	LowFailure                        // 1
	LowPassButCertifying              // 2
	Satisfactory                      // 3
	Good                              // 4
	SuperiorHonor                     // 5
	OutstandingHighHonor              // 6
)

func returnGrade(idGrade int) string {
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

// 2. Get the average score across all subjects for a student id
func (s Student) avgGrade() (avg int, err error) {
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
func (s Student) avgGradeBySbj(sbj int) (avg int, err error) {
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

// 4. Get the subject with the lowest score for a Student id
func (s Student) lowGradeBySbj(sbj int) (low int, err error) {
	low = 0
	for _, v := range Exams {
		if v.idStd == s.id && v.subject == sbj {
			if low == 0 {
				low = v.grade
			}
			if v.grade < low {
				low = v.grade
			}
		}
	}
	if low == 0 {
		return 0, errors.New("exams not found for this student")
	}
	return low, nil
}

// 5. Get the subject with the highest score for a Student id
func (s Student) highGradeBySbj(sbj int) (high int, err error) {
	high = 0
	for _, v := range Exams {
		if v.idStd == s.id && v.subject == sbj {
			if high == 0 {
				high = v.grade
			}
			if v.grade > high {
				high = v.grade
			}
		}
	}
	if high == 0 {
		return 0, errors.New("exams not found for this student")
	}
	return high, nil
}

// 6. Get all grades for a Student id
func (s Student) allGrade() (result []int, err error) {
	for _, v := range Exams {
		if v.idStd == s.id {
			result = append(result, v.grade)
		}
	}
	if result == nil {
		return nil, errors.New("exams not found for this student")
	}
	return result, nil
}

// extended print Exams
func (e exam) extPrint() {
	prt("id exam:", e.id)
	prt("student name:", Students[e.idStd].name)
	prt("subject:", returnSbj(e.subject))
	prt("grade:", returnGrade(e.grade))
}

// print average grades
// func prtAvg(avg int, err error) {
// 	if err != nil {
// 		prt(err)
// 		return
// 	}
// 	prt(avg)
// }

func prtAvg(avg interface{}, err error) { // I know, it's bad
	if err != nil {
		prt(err)
		return
	}
	prt(avg)
}