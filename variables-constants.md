
---

# 📘 Golang Variables: Concepts & Examples

---

## 🟦 Keywords in Variables

**Key Points:**

* Go uses the `var`, `const`, `type`, and `func` keywords.
* `var` is used to declare variables explicitly.

```go
var name string = "John"
```

---

## 🔤 Syntax to Declare Variables

**Key Points:**

* `var <name> <type> = <value>`
* Type or value can be omitted in some cases.

```go
var age int = 25
var city = "New York"     // Type inferred
var country string         // Defaults to ""
```

---

## 🔢 Data Types in Variables

### **Primitive Data Types**

| Type      | Default Value | Example Values  |
| --------- | ------------- | --------------- |
| `int`     | `0`           | `-1`, `0`, `42` |
| `float64` | `0.0`         | `3.14`, `0.0`   |
| `string`  | `""`          | `"Hello"`, `""` |
| `bool`    | `false`       | `true`, `false` |

```go
var score int
var pi float64 = 3.14
var isValid bool = true
var greeting string = "Hello"
```

---

### **Subtypes in `uint` and `float`**

**Unsigned Integers:**

* `uint`, `uint8`, `uint16`, `uint32`, `uint64`

```go
var b uint8 = 255 // Max for uint8
```

**Floating-Point Numbers:**

* `float32`, `float64` (default)

```go
var smallPi float32 = 3.14
```

---

## 🔁 Aliases and Default Values for Datatypes

**Key Points:**

* Aliases help improve code readability and abstraction.
* Example aliases:

  * `byte` is alias for `uint8`
  * `rune` is alias for `int32`

```go
var b byte = 100      // same as uint8
var r rune = 'A'      // same as int32
```

---

## 🧠 Implicit Types & Lexer Behavior

**Key Points:**

* Go lexer infers the type based on the assigned value.
* The Go compiler performs static type checking.

```go
var x = 10        // int
var y = 10.5      // float64
var z = "Go"      // string
```

---

## ❓ Declaring Variables Without Type

**Key Points:**

* Type is inferred automatically from the assigned value.

```go
var count = 10           // inferred as int
var message = "Welcome"  // inferred as string
```

---

## ⚡ Declaring Variables Without `var` Keyword

**Key Points:**

* Use **short declaration** syntax with `:=` (walrus operator).
* Only allowed **inside functions**.

```go
func main() {
    name := "Alice"   // same as var name string = "Alice"
}
```

---

## 🐘 Walrus (`:=`) Operator

**Key Points:**

* Shortcut for declaring and initializing variables.
* Not allowed at the package level (outside functions).

```go
// Correct usage
func example() {
    sum := 5 + 3
}
```

```go
// ❌ Incorrect: outside function
// name := "Bob" // invalid outside function
```

---

## 🔒 Declaring Constants

**Key Points:**

* Use `const` keyword
* Must be assigned at declaration
* Cannot use `:=` with constants

```go
const Pi = 3.14
const Language string = "Go"
```

---

## 🌐 Declaring Variables/Constants at Global Scope

**Key Points:**

* Global declarations are done at the package level (outside functions).
* Can use `var` or `const`, **not** `:=`.

```go
package main

import "fmt"

var site = "openai.com"    // global variable
const version = "1.0"      // global constant

func main() {
    fmt.Println(site, version)
}
```

---

## ✅ Summary

| Concept                    | Syntax/Notes                            |
| -------------------------- | --------------------------------------- |
| Variable declaration       | `var name type = value`                 |
| Short declaration          | `name := value` (inside functions only) |
| Constants                  | `const name = value`                    |
| Default values             | `0`, `0.0`, `""`, `false`, `nil`        |
| Type inference             | Go infers from value                    |
| Global variables/constants | Use `var` and `const` only              |
| Aliases                    | `byte = uint8`, `rune = int32`          |

---

Let me know if you'd like this as a downloadable `.md` file or want to expand with advanced examples (e.g., pointers, structs, etc.)!
