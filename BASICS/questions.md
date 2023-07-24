Questions:

1) from Learning Go: "Strings in Go are immutable; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it."
Strings are immutable because we cannot perform any append and remove operations on it. Every time we assign the value to the string it creates new instance of it and utilizes the memory. So doing a string manipulation can result in making copies.

The immutability of strings is not the same as immutability of variables.
Immutability of strings means that the characters in the string cannot be changed (we cannot perform any append and remove operations on it). This holds true for Go. Go makes use of it when slicing strings as shown in the example below.
var s string = "Hello, World"
s[1] = 'c' // It is not allowed


Ok, maybe I understand what is meant by immutability in this context.
"Immutability does not imply that the object as stored in the computer's memory is unwriteable. Rather, immutability is a compile-time construct that indicates what a programmer can do through the normal interface of the object, not necessarily what they can absolutely do."

For the string Go use string interning, which is a method of storing only one copy of each distinct string value, which must be immutable.
The underlying structure of strings is basically similar to that of a slice: a string is internally a structure consisting of a pointer to a byte array and a field containing the length of the string (while a slice is a pointer to an array of numbers plus length and capacity fields). (https://research.swtch.com/godata)

The question I ask you is: how important is this concept? Is this something I need to understand deeply now or can I study it later?


2) multiple declaration and initialization is confusing for me. E.g.:
a := 5
a, b := 7, 11	// b is new and it is declared and initialized when a is only initialized witha new value
Is it better avoid it?


3) Wraparound vs Overflow vs Saturation. I found this concept very confusing and it is new to me. I mean, I understand how this works but I don't understand why this behavior is by design.
Why is there no panic error when running?
Even more confusing if I do a standard for loop with a strict inequality, the compiler will catch the bug, not if i check with minus or equal:
var b byte
for b = 250; b < 256; b++ {	//compiler error: constant 256 overflows byte
for b = 250; b <= 255; b++ { // infinite loop
But ok, I read how to solve this, it's up to the programmer to check.


4) why constant and literals are untyped? What does it mean and what are the implications?
