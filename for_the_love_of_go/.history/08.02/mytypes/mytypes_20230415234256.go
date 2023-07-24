package mytypes

/* Twice multiplies its receiver by 2 and returns the result
 */

// import "fmt"
type MyInt int

func (i MyInt) Twice() int {
	return i * 2
}
