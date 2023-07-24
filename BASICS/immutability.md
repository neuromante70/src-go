Questions:

1) from Learning Go: "Strings in Go are immutable; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it."
Strings are immutable because we cannot perform any append and remove operations on it. Every time we assign the value to the string it creates new instance of it and utilizes the memory. So doing a string manipulation can result in making copies.

The immutability of strings is not the same as immutability of variables.
Immutability of strings means that the characters in the string cannot be changed (we cannot perform any append and remove operations on it). This holds true for Go. Go makes use of it when slicing strings as shown in the example below.
Variables in Go are always mutable. 
When a string variable is changed, the internal fields of the variable (pointer and length) are changed. The address of variable never changes. Every time we assign the value to the string it creates new instance of it and utilizes the memory.

from Learning Go: Pointers Indicate Mutable Parameters
As we’ve already seen, Go constants provide names for literal expressions that can be calculated at compile time. There is no mechanism in the language to declare that other kinds of values are immutable. Modern software engineering embraces immutability. MIT’s course on Software Construction sums up the reasons why: “[I]mmutable types are safer from bugs, easier to understand, and more ready for change.
Mutability makes it harder to understand what your program is doing, and much harder to enforce contracts.”
The lack of immutable declarations in Go might seem problematic, but the ability to choose between value and pointer parameter types addresses the issue. As the Software Construction course materials go on to explain: “[U]sing mutable objects is just fine if you are using them entirely locally within a method, and with only one reference to the object.” Rather than declare that some variables and parameters are immutable, Go developers use pointers to indicate that a parameter is mutable.
Since Go is a call by value language, the values passed to functions are copies. For nonpointer types like primitives, structs, and arrays, this means that the called function cannot modify the original. Since the called function has a copy of the original data, the immutability of the original data is guaranteed.
