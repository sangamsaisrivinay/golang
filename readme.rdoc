- `Go` is a compiled lamguage
- Go is not a pure OOP
    - yes
        - class based
        - interfaces support
        - inheritance
        - encapsulation
        - composition
    - no
        - no support for overload/ override

- we can build exe files using `go build <filename.go>`
- we can export the exes without giving code


- missing concepts
    - try/ catch (handled by `Lexer`)


- `main()` is the entry point for a golang file
- find path for go installation
    `go env GOPATH`

- all the main packages are present in

<pre>
```
TEMP=$(go env GOPATH)
echo $TEMP
cd $TEMP/pkg/mod
ls -l
```
</pre>

<pre>
```
cache
github.com
golang.org
honnef.co
mvdan.cc
```
</pre>

<br>

##importing packages and using functions

Here is the full content of the `README.md` file:

````markdown
# 📘 Go: Public & Private Functions + Cross-Package Imports

This guide covers how to:

- Create **public vs private functions** in Go
- Organize code across multiple **files** and **folders (packages)**
- Import code **with and without `go.mod`**

---

## 📂 Project Structure Example

```bash
myproject/
├── go.mod         # only if using modules
├── main/
│   └── main.go
└── greet/
    └── greet.go
````

---

## 🧱 1. Public vs Private Functions in Go

| Function Name | Scope          | Usage                                   |
| ------------- | -------------- | --------------------------------------- |
| `Hello()`     | **Exported**   | Accessible from other packages          |
| `hello()`     | **Unexported** | Accessible only inside the same package |

### 🧪 Example: `greet/greet.go`

```go
package greet

import "fmt"

func SayHello() {
    fmt.Println("Hello from greet") // ✅ Public
}

func whisperHello() {
    fmt.Println("...hello...")      // ❌ Private outside package
}
```

---

## 🔗 2. Accessing Functions from Another Folder (Package)

### ✅ A. If `go.mod` **is present** (Go Modules)

#### 🧰 Setup

```bash
go mod init myproject
```

#### 📄 main/main.go

```go
package main

import (
    "myproject/greet" // Folder path based on go.mod name
)

func main() {
    greet.SayHello() // ✅ OK
}
```

#### 🏃 Run:

```bash
go run main/main.go
```

---

### ❌ B. If `go.mod` is **absent**

* You must run with all relevant files together
* Import **local folders** without using module paths

#### 📄 main.go

```go
package main

import (
    "./greet" // ✅ Only works without modules
)

func main() {
    greet.SayHello()
}
```

#### 🏃 Run:

```bash
go run main.go greet/greet.go
```

> ⚠️ **Not recommended** for long-term use — prefer `go.mod`.

---

## 🛠 3. Key Rules and Tips

* ✅ **Import by folder path**, not package name
* ✅ **Function/type must be Capitalized** to be exported (public)
* ❌ **All files in the same folder must use the same `package` name**
* ✅ Folder name and `package` name **can differ**, but it’s **best practice to match them**

---

## 📌 Summary

| Scenario                             | Import Path                                  | Reference Name      | Example                    |
| ------------------------------------ | -------------------------------------------- | ------------------- | -------------------------- |
| `go.mod` present                     | Folder path relative to module name          | `greet.SayHello()`  | `import "myproject/greet"` |
| `go.mod` absent                      | Relative folder path                         | `greet.SayHello()`  | `import "./greet"`         |
| Different package name inside folder | Must reference using declared `package` name | `banana.SayHello()` | —                          |

---

## 🧪 Example Repo

```bash
myproject/
├── go.mod               # module myproject
├── greet/
│   └── greet.go         # package greet
├── utils/
│   └── math.go          # package utils
└── main/
    └── main.go          # package main
```

Use:

```go
import (
    "myproject/greet"
    "myproject/utils"
)
```

```
```





Here’s a **fully working Go example** with folder structure, code files, and instructions that demonstrate:

* Public vs private functions
* How to use `go.mod`
* How to import and use code from another package (folder)

---

## 📁 Folder Structure

```bash
myproject/
├── go.mod
├── main/
│   └── main.go
└── greet/
    └── greet.go
```

---

## 🔧 Step 1: Create the Project

1. Open terminal and run:

```bash
mkdir -p myproject/{main,greet}
cd myproject
go mod init myproject
```

---

## 📄 `go.mod`

Created automatically by `go mod init myproject`:

```go
module myproject

go 1.21
```

---

## 📄 `greet/greet.go`

```go
package greet

import "fmt"

// Exported function (public)
func SayHello() {
	fmt.Println("Hello from greet!")
}

// Unexported function (private)
func whisperHello() {
	fmt.Println("...psst... hello quietly")
}
```

---

## 📄 `main/main.go`

```go
package main

import (
	"myproject/greet"
)

func main() {
	greet.SayHello()       // ✅ works: SayHello is exported
	// greet.whisperHello() // ❌ error: whisperHello is unexported
}
```

---

## ▶️ Step 2: Run the Program

From the `myproject` root folder:

```bash
go run main/main.go
```

You’ll see:

```
Hello from greet!
```

---

## ❌ Trying to Call Private Function?

Uncommenting `greet.whisperHello()` will cause this error:

```
cannot refer to unexported name greet.whisperHello
```

This confirms Go’s **export rules**: only capitalized names are accessible across packages.

---

## ✅ Summary

| What You Did                     | Result                      |
| -------------------------------- | --------------------------- |
| Used `go.mod`                    | Enables clean module import |
| Created `SayHello()` (capital)   | Accessible from `main`      |
| Created `whisperHello()` (lower) | Not accessible from `main`  |
| Imported using folder path       | Matches Go convention       |

---




