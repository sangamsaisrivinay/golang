In Go, **slices** are a powerful, flexible abstraction over arrays. They're used frequently in Go programming due to their dynamic size and ease of use.

---

### 🔹 What is a Slice?

A **slice** is a **descriptor** of an array segment. It includes:

* A pointer to the underlying array
* Length (number of elements it contains)
* Capacity (max number of elements before realloc is needed)

Unlike arrays, slices are dynamic and can grow with `append()`.

---

### 🔸 Slice Initialization Syntax

Here are several ways to initialize slices:

#### 1. **Literal initialization**

```go
nums := []int{1, 2, 3}
```

#### 2. **Slicing an array**

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4] // elements at index 1, 2, 3 => [20 30 40]
```

#### 3. **Using `make()`**

```go
s := make([]int, 3) // len=3, cap=3 => [0 0 0]
t := make([]int, 3, 5) // len=3, cap=5
```

---

### 🔸 `append()` Function

`append()` adds elements to a slice. If capacity is exceeded, a new underlying array is allocated.

```go
a := []int{1, 2}
a = append(a, 3, 4) // [1 2 3 4]
```

You can also append another slice:

```go
b := []int{5, 6}
a = append(a, b...) // [1 2 3 4 5 6]
```

---

### 🔸 `sort` Package with Slices

The `sort` package provides utilities for sorting slices.

```go
import "sort"

s := []int{5, 2, 6, 3}
sort.Ints(s) // modifies s in-place => [2 3 5 6]
```

For custom sorting:

```go
people := []string{"Charlie", "Alice", "Bob"}
sort.Slice(people, func(i, j int) bool {
    return people[i] < people[j]
})
```

---

### 🔸 `:` Operator (Slicing)

Used to create sub-slices:

```go
a := []int{10, 20, 30, 40, 50}
b := a[1:4] // [20 30 40], includes index 1 to 3 (exclusive of 4)
```

Syntax:

```go
slice[low:high]      // from index low to high-1
slice[low:]          // from index low to end
slice[:high]         // from start to high-1
slice[:]             // full slice copy
```

---

### ✅ Summary

| Feature               | Description                                 |
| --------------------- | ------------------------------------------- |
| `[]T{}`               | Literal slice                               |
| `make([]T, len, cap)` | Create slice with given length and capacity |
| `append()`            | Add elements, may allocate                  |
| `sort` package        | Sort built-in or custom slices              |
| `:` operator          | Create sub-slices                           |
