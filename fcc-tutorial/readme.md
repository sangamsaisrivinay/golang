- go is statically typed
    - datatype shd be given during declaration
    - or it autodetects datatype at intialization and cant be changed later
    eg: var n= "4"
    var is string datatype
- strongly typed
    - operations perfoemd on variables depends in datatype
- go is compiled
    - gives executable binary file after coimpiling
- fast compile time
- built in concurrency - no need of addl packages
- simple syntax, auto garbage collection

---

# structure of a go pre
- packages - collection of modules
- modules - collection of go files

---

# initializing go project
- initializing new project means we initialize a new module

    `go mod init <module_name>`


    eg: here we are creating eg, first-project
    - we will have go.mod file which contains the module name and the installed dependencies along with their versions

---

# functions in go
- declared using `func` keyword
eg:
<pre>
package main
import "fmt"
func main(){
    fmt.Print("hello")
}
</pre>


---

# package


- we use import to use a package in go file
eg: `import "fmt"`
- go doesnt allow import and keeping the package unused


- every go file belong to a package
- it is identified by typing the package name at the top of pre in the go file

eg: `package main`


# package main
- for executable files, it require the main package
- it looks for main() function as the entry point in the main package
- main function doesnt take any parameters
- no need to create such functions for other packages

- other packages can serve as reusable libraries in go

---

# points to remember: rel b/w folders and packages in go
- all the files in a folder should belong to same package
- package name and folder name need not be same, but can be same for clarity

---

# compile and run go file
- command for build
`go build <file>.go`
- this generates a .exe file which we can run directly like a normal exe file
`./filename` gives the intended output of the file
- command to compile and run
`go run <file.go>`

---

# datatypes

- int (int32)
- int8
- int16
- uint
- float32
- float64
- string

## rules for operations on datatypes
- add/sub/mult/div not applicable for diff types
- one of the variables shd be casted

<pre>
var a int = 4;
var b float32 = 65;
var res float32 = a + b; ❌
res = float32(a) + b;
</pre>

- division of 2 integers, will give integer trounded to nearest number
- use `%` for getting remainder

<pre>
func main(){
	var a int= 4;
	var b int = 16;
	var c int = (a) / b
	var d int = a % b
	fmt.Println(c)
	fmt.Println(d)
}
</pre>


- use + to concatenate strings
- we have `len()` function to get the length of string, but it gives no of bytes, but not the actual string length
- when we have special characters, iy takes 2 bytes

<pre>
fmt.Println(len("α"))
gives length as 2
</pre>

- so we use a 3rd party library to get the string length

## declaring strings
- "" for single line string
- `` for multi line strings

## special case for string
- use the package `unipre/utf8` which provide the function `RuneCountInString` to get the exact string length

<pre>
import ("fmt"
"unipre/utf8"
);

func main(){
	fmt.Println("c"+"d")
	fmt.Println(utf8.RuneCountInString("α"))
}

gives the length as 1
</pre>


- booleans
have eitehr true or false value
---

# walrus operator :=

- used to remove the var declaration for variables
- short hand variable declaration operator

usage:
<pre>
x := 5 ✅
var x:= 5 ❌
var1, var2 := 4, 5


func main(){
	 x, y := 5, 6
	fmt.Println(x)
	fmt.Println(y)
}

</pre>

---

# constants, variables

- unused variables are not allowed in go file

eg:
<pre>
import "fmt"
func abc(){
    var num int
    fmt.Print(num)
}
</pre>

- constants: no vslue change, must be initialized during declaration

const myVar int = 5

---


# functions with parameters

## passing parameters in function

<pre>
package main

import ("fmt"
);

func print_str(str string){
	fmt.Print(str)
}

func main(){
	print_str("hello")
}
</pre>

## return value from function

<pre>
package main

import ("fmt"
);

func add_nums(a int, b int) int{
	return a+b
}

func main(){
	fmt.Print(add_nums(5,4))
}

</pre>

## return multiple values

<pre>
package main

import ("fmt"
);

func div_mod(a int, b int) (int, int){
	return a/b, a%b
}

func main(){
	fmt.Print(div_mod(5,4))
}
</pre>

## handle errors returned by function
- we use the `"errors"` module to handle any errors/ exceptions occured during the function execution and we will return that error from function

<pre>
package main

import ("fmt"
"errors"
);

func div_mod(a int, b int) (int, int, error){
	var err error
	if(b==0){
		err= errors.New("cant divide by zero")
		return 0,0,err
	}
	return a/b, a%b, err
}

func main(){
	var div, mod, err= div_mod(5,0)
	if(err!=nil){
		fmt.Print(err.Error())
	}
	fmt.Print(div,mod)
}
</pre>

---

# control structures and logical operators

- &&
- ||
- !

## if-else-if

<pre>
loggedIn := false
if !loggedIn {
    fmt.Println("Please log in first")
}
</pre>

<pre>
day := "Sunday"
if day == "Saturday" || day == "Sunday" {
    fmt.Println("Weekend")
}
</pre>

<pre>
age := 25
if age > 18 && age < 60 {
    fmt.Println("Eligible to work")
}
</pre>

## switch

### switch without expression

<pre>
age := 20

switch {
case age < 13:
    fmt.Println("Child")
case age < 18:
    fmt.Println("Teenager")
default:
    fmt.Println("Adult")
}
</pre>

### basic switch

<pre>
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("End of the work week")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Midweek day")
    }
}

</pre>

*in go switch, break is not requires, when a case is matched it automatically breaks out of the loop*

### fallthrough

- we use `fallthrough` if we wnt to explicitly execute next case

<pre>
num := 1
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")  // This runs because of fallthrough
default:
    fmt.Println("Default")
}

</pre>

---


# Arrays, Slices, Maps

## Arrays

- fixed length
`eg: var intArr [3]int`
- same type
all elements are same type in array
initialized to default type values, eg. 0 for int
- indexable
can access elements using indeces
`eg: fmt.Print(intArr[0])`
slicing: uses `:` operator to print range of elements
`eg: fmt.Print(intArr[0:2]) - prints element at 0, 1 indeces`
- contiguous memeory locations
<pre>
func main(){
	var intArr [3]int= [3]int{0,1,2}
	fmt.Println(intArr[1])
	fmt.Println(intArr[1:2])
	//get location of elements using &
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}
</pre>
- static initialization
    - var intArr [3]int= [3]int{0,1,2} or
    - intArr := [3]int{0,1,2}

`intArr := [...]int{0,1,2}`
- this also initializes array without providing the explicit size
- since we gave only 3 elements statically, the array size will still be 3

## Slices
- slices wrap around arrays to give additional functionality

<pre>
intSlice := []int{0,1,2}
</pre>

- appending the element to a slice
    - before appending the element, if there is no vacancy in the array, a new array is formed and the existing elements are copied into the new array and the element os appended to the array
- checking the length of slice will have 3
- capacity of slice is 3

<pre>
intSlice := []int{0,1,2}
append(intSlice, 4)
</pre>

- length of slice will be 4
- capacity of slice is 6, but cant access uninitialized elements
- the capacity will be doubled to the current capacity when appending a new element to slice

<pre>
package main

import ("fmt"
);


func main(){
	intArr := []int {0,1,2};
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	fmt.Println(&intArr[0])
	intArr = append(intArr, 5)
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	fmt.Println("memory address changed...")
	fmt.Println(&intArr[0])
    fmt.Println(intArr[5])
}
</pre>

<pre>
[0 1 2]
3
3
0x140000ac018
[0 1 2 5]
4
6
memory address changed...
0x140000c0000
panic: runtime error: index out of range [5] with length 4
</pre>

### functions in slice
length - get length of slice `len(sliceName)`
capacity - get total elements the slice can accomodate `cap(sliceName)`
append - appends a variable or a slice `append(sliceName, varName)`


### spread operator
spread operator in slice `...`

- appending a slice

<pre>
package main

import ("fmt"
);


func main(){
	intArr := []int {0,1,2};
	intArr1 := []int {3,3,3, 4 ,5};
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	intArr = append(intArr, intArr1...)
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	intArr= append(intArr, 10)
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
}
</pre>

<pre>
[0 1 2]
3
3
[0 1 2 3 3 3 4 5]
8
8
[0 1 2 3 3 3 4 5 10]
9
16
</pre>

### make() function for slice

- to set the length, initial capacity of the slice

var intSlice []int= make([]int, 3, 8)

<pre>
s := make([]int, 3, 5)  // length = 3, capacity = 5
fmt.Println(len(s), cap(s))  // Output: 3 5

s := make([]int, 3)  // length = 3, capacity = 3
fmt.Println(s)       // Output: [0 0 0]

</pre>