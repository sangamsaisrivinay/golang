### missed
- [ ] defer
- [ ] methods, function vs method
- [ ] break, continue, goto
- [ ] pointers



# theory

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
- boolean

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
- use the package `unicode/utf8` which provide the function `RuneCountInString` to get the exact string length

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

<pre>
if exist{
		fmt.Println("exist")
}else{
    fmt.Println("not exist")
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
delete - delete a value

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
- default capacity is length of slice

var intSlice []int= make([]int, 3, 8)

<pre>
s := make([]int, 3, 5)  // length = 3, capacity = 5
fmt.Println(len(s), cap(s))  // Output: 3 5

s := make([]int, 3)  // length = 3, capacity = 3
fmt.Println(s)       // Output: [0 0 0]

</pre>

## Maps

- can create map using make() function

### declaration
<pre>
var myMap map[string]int = make(map[string]int)
</pre>

### static initialization
<pre>
var myMap map[string]int = map[string]int {"a":5, "b":4}
</pre>

### accessing values
- `myMap[key]`

### we can use a additional variable to check if the value for given key exist or not
`var val, exist= myMap1["b"]`

<pre>
package main

import ("fmt"
);


func main(){
	var myMap map[string]int = make(map[string]int)
	fmt.Println(myMap)
	var myMap1 map[string]int = map[string]int{"a":5, "b":6}
	fmt.Println(myMap1["a"])
	fmt.Println(myMap1["c"])
	var val, exist= myMap1["b"]
	if exist{
		fmt.Println("exist")
	}else{
		fmt.Println("not exist")
	}
	val, exist= myMap["c"]
	if exist{
		fmt.Println("exist")
	}else{
		fmt.Println("not exist")
	}
	fmt.Println(val)
}
</pre>

<pre>
map[]
5
0
exist
not exist
0
</pre>
---

# for loop

- variants in for loop

<pre>
i:=0
for {
    if i>10{
        break
    }
    fmt.Println(i)
    i++
}
</pre>

<pre>
for i:=0; i<10; i++ {
		fmt.Println(i)
	}
</pre>

<pre>
i:=0
for i<10{
    fmt.Println(i)
    i++
}
</pre>

## range - for loop: iterating arrays, slices, maps using loops

`range` actually iterates over and prints the indeces of the datastructure in case of array, slice
to get the actual values we use 2 variables to access it


### arrays
- iterate and print indeces
<pre>
myArr := [5]int {0,1,2}
	for val:= range myArr{
		fmt.Println(val)
	}
</pre>

<pre>
the actually printed values are the indeces
0
1
2
3
4
</pre>

- to print the values, we can use `_` operator to ignore it

<pre>
myArr := [5]int {0,1,3,4}
for _, val:= range myArr{
    fmt.Println(val)
}
</pre>

<pre>
0
1
3
4
0
</pre>

- to print both index, value we can do

<pre>
myArr := [5]int {0,1,3,4}
for index, val:= range myArr{
    fmt.Println(index, val)
}
</pre>
0 0
1 1
2 3
3 4
4 0
<pre>

</pre>

### slices
<pre>
mySlice := []int {0,1,2}
for val := range mySlice{
    fmt.Println(val)
}
</pre>

<pre>
the actually printed values are the indeces
0
1
2
</pre>

- to print the actual values we use
<pre>
mySlice := []int {0,1,2}
for _, val := range mySlice{
    fmt.Println(val)
}
</pre>

### maps
<pre>
var myMap1 map[string]int = map[string]int{"a":5, "b":6}
for key, val := range myMap1{
    fmt.Printf(`key: %s, value: %d`, key, val)
    fmt.Println()
}
</pre>

<pre>
key: a, value: 5
key: b, value: 6
</pre>

<b>
go has no concept of while loop
</b>

---

# _ operator in go

- In Go, the underscore `_` is called the blank identifier. It's a special symbol used when you want to ignore a value returned by a function or loop, or when you intentionally don't use a variable.

| Usage                  | Purpose                      |
| ---------------------- | ---------------------------- |
| `_ = something`        | Discard unused variable      |
| `for _, val := range`  | Ignore index in loops        |
| `for i, _ := range`    | Ignore value in loops        |
| `_, err := function()` | Discard result or error      |
| `import _ "package"`   | Import only for side effects |


---

# Strings, Runes and Bytes

- normal characters are occupying 16 bits for (-127 to 128), so it occupies 1byte as per utf16 in go
- if there are any special characters like emoji they need 32 bits, so we will get 2 seperate bytes acc to utf16 which is standard in go
- hence for any special characters we get string length as 2
- in runes its calculated exactly in no of runes
- hence for 32 bit also it counts length as 1 only

the range operator will print the utf value of any given character in the string

eg:
<pre>
package main

import ("fmt"
"unicode/utf8"
);


func main(){

	str := "helloα"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	for k, v := range str{
		fmt.Println(k,v)
	}
}
</pre>

<pre>
7
6
0 104
1 101
2 108
3 108
4 111
5 945
</pre>


- in go, strings are immutable
- if you append any character to the string, it will create a new string and append it
this is memory inefficient
- we can use string builder packages to use mutable strings, where new strig is not created

eg:
<pre>
package main

import (
	"fmt"
	"strings"
);

func main(){
strSlice := []string{"h", "e", "l", "l", "o"}
var strBuilder strings.Builder
for _, i:= range strSlice{
	strBuilder.WriteString(i)
}
finalStr:= strBuilder.String()
fmt.Println(finalStr)
}

o/P: hello
</pre>

---

# structs and interfaces

## structs
struct is a own type in go
it can contain multiple datatypes in it

### keywords
- struct
- type

### declaring
`type structName struct{}`

eg:
<pre>
package main

import(
	"fmt"
)

type gasEngine struct{
	milesPerGallon int8
	gallons int8
}

func main(){
    engine1 := gasEngine{25, 30}
    fmt.Printf(`mpg: %d, gallons: %d`, engine1.milesPerGallon, engine1.gallons)
}
</pre>

### nested structs
- using a struct as type inside another struct

<pre>
package main

import(
	"fmt"
)

type gasEngine struct{
	milesPerGallon int8
	gallons int8
	ownerInfo owner
}

type owner struct{
	name string
}

func main(){

	owner1 := owner{"raj"}
	engine1 := gasEngine{25, 30, owner1}
	fmt.Printf(`mpg: %d, gallons: %d, owner: %s`, engine1.milesPerGallon, engine1.gallons, engine1.ownerInfo.name)

}

o/p: mpg: 25, gallons: 30, owner: raj
</pre>


### anonymous structs
- struct without a name
- not reusable

eg1:
<pre>
type gasEngine struct{
	milesPerGallon int8
	gallons int8
	owner
}

type owner struct{
	name string
}

func main(){
	engine1 := gasEngine{25, 30, owner{"raj"}}
	fmt.Printf(`mpg: %d, gallons: %d, owner: %s`, engine1.milesPerGallon, engine1.gallons, engine1.name)
}
</pre>

eg2:
<pre>
myEngine2 := struct{
                mpg uint8
            }{5}
fmt.Println(myEngine2.mpg)
</pre>


- now lets define a struct for electric also

<pre>
type ev struct{
    miles unit8
    kwh uint8
}

this will also have miles left method with same logic as gas engine
</pre>


- so, we have multiple structs with same field types and methods for them
- as instances/ structs increase (eg: hydrogen fuel, cng etc..) we shd repeat the logic
- instead we can define a generic type with the fields and methods

it is called interface

---

# interfaces in go

- They enable polymorphism — the ability to write code that works with different types as long as they implement the same set of methods.
- An interface is a type that specifies a set of method signatures. Any type that implements those methods automatically satisfies the interface — no explicit declaration is needed.


## keywords
- type
- interface

- we can directly put the method definiton inside interface and use it for multiple struct types

eg:

### defining

<pre>
type Shape interface {
    Area() float64
}
</pre>

- structs that implement it

<pre>
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}
</pre>

- using in a function

<pre>
func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
</pre>

<pre>
func main() {
    c := Circle{radius: 5}
    r := Rectangle{width: 4, height: 3}

    printArea(c)  // Area: 78.5
    printArea(r)  // Area: 12
}
</pre>


### interface with multiple methods

<pre>
type Animal interface {
    Speak() string
    Type() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string { return "Woof!" }
func (d Dog) Type() string  { return "Dog" }

func (c Cat) Speak() string { return "Meow!" }
func (c Cat) Type() string  { return "Cat" }

func describe(a Animal) {
    fmt.Printf("%s says %s\n", a.Type(), a.Speak())
}

func main() {
    describe(Dog{}) // Dog says Woof!
    describe(Cat{}) // Cat says Meow!
}

</pre>

### empty interface

<pre>
func printAnything(i interface{}) {
    fmt.Println(i)
}

printAnything(42)
printAnything("hello")
printAnything(true)

</pre>

---

# methods in go

we directly assigned engine object to the function

<pre>
type gasEngine struct{
	milesPerGallon int8
	gallons int8
}
func (e gasEngine) milesLeft() int{
	return int(e.milesPerGallon)* int(e.gallons)
}
func main(){
	engine1 := gasEngine{25, 30}
	fmt.Println(engine1.milesLeft())
}
</pre>

- the method is bound to a reciver object type, like a struct

funciton with extra reviever object - method

func `(e gasEngine)` milesLeft() int{
	return int(e.milesPerGallon)* int(e.gallons)
}


Use a method when:
- The logic is tightly related to a type (struct)

---

# pointers
- they are special types that stores memory adress of variables

- declaring a pointer
<pre>
var p *int32
</pre>

- this will initially point to null
- it doesnt give compile errors
- but whebn ran it gives null pointer exception

<pre>
func main(){

	var i int32
	var p *int32
	fmt.Println(p)
	fmt.Println(*p)

}

o/p:
<nil>
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x100cb4784]

</pre>


- we will initialize a pointer using `new` key word

`var p *int32 = new (int32)`

- using * notation
	- we will use * in 2 places
		- to declare a pointer variable
		- to reference/ access the value stored at the pointers address
	- by default the value stored is the datatype default value, i.e. if we just initialize a pointer variable with new keyowrd, it stores the default value of that datatype to which it is intialized
	- we can point the pointer to address of another variable using & notation, eg. p= &i i.e. it points to the adress of value of i.
		- p & i refers to same adress location
		- ### if we change the value of p using *p the value of i also changes

summary:
- *int - declares a pointer variable of type int
- *p - gets the value at adress stored in p
- &i - gets the memory location of i


<pre>

func main(){

	var i int32
	var p *int32 = new (int32)
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)
	i=45;
	p = &i
	fmt.Println(p)
	fmt.Println(*p)

}
</pre>

o/p:
<pre>
0
0x14000104020
0
0x1400010400c
45
</pre>


## callby value & call by reference


### normal variables and pointers
- we can point the pointer to address of another variable using & notation, eg. p= &i i.e. it points to the adress of value of i.
		- p & i refers to same adress location
		- ### if we change the value of p using *p the value of i also changes

<pre>
func main(){

	var i int32
	var p *int32 = new (int32)
	i = 10
	p = &i
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&i)
	*p = 15
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(&i)
}

</pre>

o/p:
<pre>
10
10
0x1400000e09c
0x1400000e09c
15
15
0x1400000e09c
0x1400000e09c
</pre>


### slice variables and pointers

---

# Go Routines

- launch multiple functions and execute them concurrently
- jumping back and forth b/w tasks


- go keyword
- wait groups
	- .Add()
	- .Wait()
	- .Done()
- mutex
	- .Lock()
	- .Unlock()
- read-write mutex
