Here's a concise overview of **maps in Go**, including how to declare them, use `make()`, create them using different syntaxes, loop over them, and delete entries.

---

### ✅ **1. Declaring a Map using `make()`**

```go
m := make(map[string]int)
m["apple"] = 5
m["banana"] = 3
```

---

### ✅ **2. Different Syntaxes to Create Maps**

#### a. Using `make()` (as above)

```go
m := make(map[string]int)
```

#### b. Map literal with values

```go
m := map[string]int{
    "apple":  5,
    "banana": 3,
}
```

#### c. Declare first, assign later

```go
var m map[string]int // nil map, need to assign before use
m = map[string]int{
    "apple": 5,
}
```

---

### ✅ **3. Looping Over a Map (`for` loop)**

```go
m := map[string]int{
    "apple":  5,
    "banana": 3,
}

for key, value := range m {
    fmt.Println(key, value)
}
```

* Order of iteration is **random** in Go maps.

---

### ✅ **4. Deleting an Entry from a Map**

```go
delete(m, "apple")
```

* Safe even if the key doesn't exist.

---

### 🔍 Extra: Check if a Key Exists

```go
value, ok := m["apple"]
if ok {
    fmt.Println("Exists:", value)
} else {
    fmt.Println("Does not exist")
}
```

---

Let me know if you'd like to see practical examples or performance tips with maps.
