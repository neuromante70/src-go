package mytypes

/* Twice multiplies its receiver by 2 and returns the result
 */

// import "fmt"
type MyInt int
type MyString string

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString)