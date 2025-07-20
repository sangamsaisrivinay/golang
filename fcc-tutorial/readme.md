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

#### structure of a go code
- packages - collection of modules
- modules - collection of go files

---

### initializing go project
- initializing new project means we initialize a new module

    `go mod init <module_name>`


    eg: here we are creating eg, first-project
    - we will have go.mod file which contains the module name and the installed dependencies along with their versions

---

## functions in go
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
- it is identified by typing the package name at the top of code in the go file

eg: `package main`


# package main
- for executable files, it require the main package
- it looks for main() function as the entry point in the main package
- main function doesnt take any parameters
- no need to create such functions for other packages

- other packages can serve as reusable libraries in go

---

### points to remember: rel b/w folders and packages in go
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

## datatypes

- int (int32)
- int8
- int16
- uint
- float32
- float64
- string

### rules for operations on datatypes
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

##### declaring strings
- "" for single line string
- `` for multi line strings

### special case for string
- use the package `unicode/utf8` which provide the function `RuneCountInString` to get the exact string length

<pre>
import ("fmt"
"unicode/utf8"
);

func main(){
	fmt.Println("c"+"d")
	fmt.Println(utf8.RuneCountInString("α"))
}

gives the length as 1
</pre>


- booleans

---

## constants, variables

- unused variables are not allowed in go file

eg:
<pre>
import "fmt"
func abc(){
    var num int
    fmt.Print(num)
}
</pre>