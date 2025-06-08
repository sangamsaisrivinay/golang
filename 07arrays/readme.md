In Go, **arrays** are fixed-length sequences of elements of the same type. They are value types, meaning when you assign or pass them, a copy is made.

---

### 📦 Array Basics

```go
var a [5]int // Declares an array of 5 integers, default initialized to 0
```

* Arrays in Go have a **fixed size** defined at compile time.
* The length is part of the type: `[5]int` is different from `[10]int`.

---

### 🧬 Types of Array Initialization

There are several ways to initialize arrays:

#### 1. **Zero Initialization (Default)**

```go
var a [3]int
fmt.Println(a) // Output: [0 0 0]
```

All elements are initialized to the zero value of the element type.

#### 2. **Explicit Initialization**

```go
a := [3]int{1, 2, 3}
fmt.Println(a) // Output: [1 2 3]
```

#### 3. **Partial Initialization**

```go
a := [5]int{1, 2}
fmt.Println(a) // Output: [1 2 0 0 0]
```

Uninitialized elements get the zero value.

#### 4. **Index-Specific Initialization**

```go
a := [5]int{2: 10, 4: 20}
fmt.Println(a) // Output: [0 0 10 0 20]
```

You can set values at specific indices.

#### 5. **Inferred Length**

```go
a := [...]int{1, 2, 3}
fmt.Println(a) // Output: [1 2 3]
```

The compiler counts the number of elements to determine the length.

---

### ⚙️ Default Functions and Behavior

Go does not provide many "default functions" for arrays like other high-level languages (e.g., Python). But here’s what you can do:

#### 1. **Length and Capacity**

```go
len(a) // Gets the length of the array
cap(a) // Same as len() for arrays (more useful for slices)
```

#### 2. **Looping**

```go
for i, v := range a {
    fmt.Println(i, v)
}
```

#### 3. **Copying (By Value)**

Arrays are copied when assigned:

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 100
fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

#### 4. **Passing to Functions**

Passed by value unless you use a pointer:

```go
func modify(a [3]int) {
    a[0] = 99
}

func modifyPtr(a *[3]int) {
    a[0] = 99
}
```

---

### 🆚 Arrays vs Slices (Heads-Up)

* Arrays: fixed size, value types.
* Slices: dynamic size, reference types — more commonly used in idiomatic Go.
