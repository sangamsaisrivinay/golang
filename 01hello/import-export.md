Here’s a concise summary of the **key points** about **visibility** and **import statements** in Go:

---

## ✅ Visibility Rules in Go

| Aspect                   | Rule                                                          | Example                                                              |
| ------------------------ | ------------------------------------------------------------- | -------------------------------------------------------------------- |
| **Exported (Public)**    | Names starting with a **capital letter** are **exported**     | `func Greet()`                                                       |
| **Unexported (Private)** | Names starting with a **lowercase letter** are **private**    | `func greet()`                                                       |
| **Scope of Visibility**  | Exported items are visible **outside the package**            | `otherpkg.MyFunc()`                                                  |
| **Unexported Scope**     | Unexported items are visible **only within the same package** | Only within the same `.go` file/folder using the same `package` name |

---

## ✅ Import Statement Rules

| Situation                  | Rule                                                                                            | Example                                               |
| -------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------- |
| **With `go.mod`**          | Import path is relative to the **module name** in `go.mod`                                      | `import "myproject/greet"`                            |
| **Without `go.mod`**       | Use **relative path** to import local packages                                                  | `import "./greet"`                                    |
| **Different package name** | You import by **folder name**, but refer using the **package name** declared inside `.go` files | `banana.SayHello()` if `package banana` in `greet.go` |
| **Same folder rule**       | All `.go` files in the same folder **must use the same package**                                | `package greet` in all files                          |
| **Function access**        | You can only call functions that are **exported (Capitalized)**                                 | `greet.SayHello()` works, not `greet.sayHello()`      |

---

## 🔐 TL;DR

* **Capital letter = Public**
* **Lowercase = Private**
* **Import by folder path**, **refer by package name**
* Prefer `go.mod` for cleaner, maintainable projects

---



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

Would you like this example zipped and downloadable?
