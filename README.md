## 🧮 CLI Base Converter (Go)

A simple command-line base converter built with Go. It allows users to convert numbers between binary, decimal, and hexadecimal formats using an interactive command-based interface.

---

## ✨ Features

🔢 Convert **binary → decimal**
🔡 Convert **hexadecimal → decimal**
🔄 Convert **decimal → binary & hexadecimal**
🚫 Input validation (handles invalid numbers gracefully)
💻 Interactive CLI with command-style input (`convert <value> <base>`)

---

## 🚀 Getting Started

### Prerequisites

* Install Go (version 1.18 or later recommended)

### Run the program

```bash
go run main.go convert.go
```

---

## 📌 Usage

Enter commands in the format:

```bash
convert <value> <base>
```

### Examples:

```bash
> convert 1E hex
✦ Decimal: 30

> convert 1010 bin
✦ Decimal: 10

> convert 15 dec
✦ Binary: 1111
✦ Hex: F
```

To exit:

```bash
> quit
```

---

## ⚠️ Error Handling

This program is designed to handle common user mistakes:

**Invalid base**

```bash
> convert 123 abc
Error: base must be hex, bin, or dec
```

**Invalid number input**

```bash
> convert 2G hex
Error: invalid hexadecimal number!
```

The program avoids crashes and provides clear feedback for incorrect input.

---

## 🧠 Project Structure

```
base-converter/
├── main.go        # Entry point and CLI logic
├── convert.go     # Conversion logic (hex, bin, dec)
├── README.md      # Project documentation
```

---

## 📖 Example Function

```go
func Convert(value string, base string) {
	switch strings.ToLower(base) {
	case "hex":
		val, _ := strconv.ParseInt(value, 16, 64)
		fmt.Println("✦ Decimal:", val)
	}
}
```

---

## 🎯 Future Improvements

* Support conversions **between all bases** (e.g. hex → bin directly)
* Add **octal** support
* Accept **flags/arguments** (`--from`, `--to`)
* Build as a **global CLI tool**
* Add unit tests

---

## 📄 License

This project is open-source and free to use.

---
