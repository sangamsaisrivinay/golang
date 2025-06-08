Sure! Below is an updated version of the `README.md` that shows all date and time examples in **IST (Indian Standard Time)** format (UTC+05:30). This version ensures time is displayed in the IST timezone using `.In(loc)` where appropriate.

---

# 🇮🇳 Working with Date and Time in Go (Golang) — IST Formats

Go's `time` package supports full time and date manipulation, including time zones like IST (Indian Standard Time: UTC+05:30). This guide shows how to format, parse, and manipulate time values using the IST timezone.

---

## 📦 Import Required Packages

```go
import (
    "fmt"
    "time"
)
```

---

## 🕒 Getting the Current Time in IST

```go
loc, _ := time.LoadLocation("Asia/Kolkata")
now := time.Now().In(loc)
fmt.Println("Current Time in IST:", now)
```

**Example Output:**

```
Current Time in IST: 2025-06-08 23:23:45.123456789 +0530 IST
```

---

## 🗓️ Formatting Dates and Times in IST

```go
loc, _ := time.LoadLocation("Asia/Kolkata")
now := time.Now().In(loc)

fmt.Println("Default:", now.String())
fmt.Println("RFC3339:", now.Format(time.RFC3339))
fmt.Println("YYYY-MM-DD HH:MM:SS:", now.Format("2006-01-02 15:04:05"))
fmt.Println("12-Hour Format:", now.Format("03:04 PM"))
fmt.Println("Day, Month Date:", now.Format("Monday, January 2, 2006"))
```

**Output:**

```
Default: 2025-06-08 23:23:45.123456789 +0530 IST
RFC3339: 2025-06-08T23:23:45+05:30
YYYY-MM-DD HH:MM:SS: 2025-06-08 23:23:45
12-Hour Format: 11:23 PM
Day, Month Date: Sunday, June 8, 2025
```

---

## 📥 Parsing a Date String into IST Timezone

```go
layout := "2006-01-02 15:04:05"
dateStr := "2025-12-31 14:00:00"
loc, _ := time.LoadLocation("Asia/Kolkata")

parsedTime, err := time.ParseInLocation(layout, dateStr, loc)
if err != nil {
    fmt.Println("Parse error:", err)
} else {
    fmt.Println("Parsed Time in IST:", parsedTime)
}
```

**Output:**

```
Parsed Time in IST: 2025-12-31 14:00:00 +0530 IST
```

---

## 🧱 Constructing a Specific Date-Time in IST

```go
loc, _ := time.LoadLocation("Asia/Kolkata")
customTime := time.Date(2025, time.March, 15, 10, 30, 0, 0, loc)
fmt.Println("Custom Time in IST:", customTime)
```

**Output:**

```
Custom Time in IST: 2025-03-15 10:30:00 +0530 IST
```

---

## ➕ Adding and Subtracting Time in IST

```go
loc, _ := time.LoadLocation("Asia/Kolkata")
now := time.Now().In(loc)

future := now.Add(48 * time.Hour)
past := now.Add(-24 * time.Hour)

fmt.Println("Now (IST):", now)
fmt.Println("48 Hours Later (IST):", future)
fmt.Println("24 Hours Ago (IST):", past)
```

---

## ⏳ Measuring Duration

```go
start := time.Now()
time.Sleep(2 * time.Second)
end := time.Now()

duration := end.Sub(start)
fmt.Println("Duration:", duration)
```

**Output:**

```
Duration: 2.000000000s
```

---

## 🌐 Converting UTC to IST

```go
utc := time.Now().UTC()
loc, _ := time.LoadLocation("Asia/Kolkata")
ist := utc.In(loc)

fmt.Println("UTC:", utc.Format("2006-01-02 15:04:05 MST"))
fmt.Println("IST:", ist.Format("2006-01-02 15:04:05 MST"))
```

**Output:**

```
UTC: 2025-06-08 17:53:45 UTC
IST: 2025-06-08 23:23:45 IST
```

---

## 📚 Resources

* [Go `time` package documentation](https://pkg.go.dev/time)
* [List of IANA Time Zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)

---

## 🛠️ Run This

Save your code as `main.go` and run:

```bash
go run main.go
```

---

Let me know if you want this output exported into a `.md` file or want to include timers, tickers, or scheduling examples in IST.
