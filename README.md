# Golang

## Golang Programming Language Learning Process
1. Learn How To Code: Google's Go (golang) Programming Language - Todd Mcleod (https://www.udemy.com/course/learn-how-to-code/learn/lecture/37482440)
### Section 2 - Intro

- Timeline
    - 2005 - First Duel core processors
    - 2006 - Go Development Started (**Meaning that this language takes advantages on multiple cores construction**)
- Goal
    - Efficient compilation
    - Efficient execution
    - Ease of programming
- Basic Syntax
    - main Function - Under Package main
    - Comment - // or /* */
    - Function Library - fmt
- Back Space (”Enter” on Keyboard) in String
    - Helps start a new line

### Section 4 - Fundamentals

- Decalcification of a value
    
    ```go
    // OPTION 1 - a VARIABLE to hold a VALUE with a certain TYPE
    var h int = 44
    var a string = "Elton"
    
    // OPTION 2 - Short Declaration Operator (:=) Go figures out the type of the variable
    // This option is not available outside of a function
    a := 42
    
    // OPTION 3 - MULTILPLE ASSIGNMENT (Best Practice - SAME TYPE VALUE ASSIGNMENT)
    b, c, d, _, f := 1, 2, 3, 4, "Happiness"
    fmt.Println(b,c,d);
    ```
    
    - Why using underline - Cuz Go helps check whether all parameters are used. If any of them is not used, Go throws an exception - Thus, need underline as a placeholder
    - Default Values
        - false - boolean
        - 0 - int
        - 0.0 - float
        - “” - string
        - nil
            - pointers
            - functions
            - interfaces
            - slices
            - channels
            - maps
        
        ```go
        // 0 for int
        var g int
        ```
        
- Decimal, Binary, & Hexadecimal
    - %b - binary
    - %x - Hexadecimal
- Conversion
    
    ```go
    z:= 42.0
    var m float32 = 32.742
    // This doesn't work
    z = m
    // Conversion needed
    z = float64(m)
    
    // More examples from Go Document
    uint(iota)               // iota value of type uint
    float32(2.718281828)     // 2.718281828 of type float32
    complex128(1)            // 1.0 + 0.0i of type complex128
    float32(0.49999999)      // 0.5 of type float32
    float64(-1e-1000)        // 0.0 of type float64
    string('x')              // "x" of type string
    string(0x266c)           // "♬" of type string
    myString("foo" + "bar")  // "foobar" of type myString
    string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
    (*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
    int(1.2)                 // illegal: 1.2 cannot be represented as an int
    string(65.0)             // illegal: 65.0 is not an integer constant
    ```
    
- Scope
- Built-in Data Types
    - [type ComplexType,](https://pkg.go.dev/builtin#ComplexType) [type FloatType,](https://pkg.go.dev/builtin#FloatType) [type IntegerType,](https://pkg.go.dev/builtin#IntegerType) [type Type,](https://pkg.go.dev/builtin#Type) [type Type1,](https://pkg.go.dev/builtin#Type1) [type any,](https://pkg.go.dev/builtin#any) [type bool,](https://pkg.go.dev/builtin#bool) [type byte,](https://pkg.go.dev/builtin#byte) [type comparable,](https://pkg.go.dev/builtin#comparable) [type complex128,](https://pkg.go.dev/builtin#complex128) [type complex64,](https://pkg.go.dev/builtin#complex64) [type error,](https://pkg.go.dev/builtin#error) [type float32](https://pkg.go.dev/builtin#float32)[type float64,](https://pkg.go.dev/builtin#float64) [type int,](https://pkg.go.dev/builtin#int) [type int16,](https://pkg.go.dev/builtin#int16) [type int32,](https://pkg.go.dev/builtin#int32) [type int64,](https://pkg.go.dev/builtin#int64) [type int8,](https://pkg.go.dev/builtin#int8) [type rune,](https://pkg.go.dev/builtin#rune) [type string,](https://pkg.go.dev/builtin#string) [type uint,](https://pkg.go.dev/builtin#uint) [type uint16,](https://pkg.go.dev/builtin#uint16) [type uint32,](https://pkg.go.dev/builtin#uint32) [type uint64,](https://pkg.go.dev/builtin#uint64) [type uint8,](https://pkg.go.dev/builtin#uint8) [type uintptr](https://pkg.go.dev/builtin#uintptr)
- Aggregate Types
    - Many values together
        - array, slice, struct
    - Composite, Compound, Structure, Struct Type
        - struct
    - The act of constructing a STRUCT which is a composite type is know as composition
        - many different types into one structure which composes all of those different types together into that one data structure - it’s a data STRUCTURE which holds VALUES of many different TYPES
        - combining multiple smaller types into a larger type
        - embedding types and inner-type  promotion
            - (inheriting fields and methos of embedded type)

### Section 5 - Hands-on Exercises

- Function Definition
    
    ```go
    func add(x int, y int) int {
    	return x + y
    }
    
    // multiple results
    func swap(x, y string) (string) {
    	return y, x
    }
    
    // named return values
    // automatically returns x and y
    func split(sum int) (x, y int) {
    	x = sum * 4 / 9
    	y = sum - x
    	return
    }
    ```