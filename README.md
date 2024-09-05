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
    
- Constants
    - character, string, boolean, or numeric values
    - Constants cannot be declared using the syntax `:=`
- Bitwise Ops
    - `<<` Left Shift
    - `>>` Right Shift
- Iota
    - Used to efficiently sign values to constant - incremental values
	
### Section 6 - Fundamentals for Programming Beginners - Skipped

### Section 7 - Environment Setup - Skipped

### Section 8 - Go mod and dependency management

- Spaghetti Code vs Modular code/Structure Code
    - Spaghetti Code - Codes in a long file
- Go Mod File - Under the project folder
    - Init - go mod init <some-module-name>
    - Example
        
        ```go
        module xxxx
        
        go.1.20
        
        require xxxx <version>
        ```
        
    - Help manages the dependencies for current project
    - `go get [github.com/xxx/xxx@latest](http://github.com/xxx/xxx@latest)` Add dependency to the mode file
    - `go mod tidy` helps update dependency and add that to your mod file
- visible/not visible; exported/not exported
    - If 1st letter of function name or variable is capitalized, then it’s visible by other go files when those go files try to import current go file.
- Example:
    - Puppy depends on Dog, Main depends on Puppy
    
    ```go
    module xxxx
    
    go.1.20
    
    require xx/xx/puppy <version>
    require xx/xx/dog <version>           // indirect 
    ```
    

### Section 9 - Hands-on Exercises

- Build - `go build xxx.go` Generates a executable files
    - Windows - `GOOS=windows go build`
    - Max - `GOOS=darwin go build`
    - Linux - `GOOS=linux go build`
- Install - `go install xxx` Helps generate the executable file under directory $GOPATH

### Section 10 - Encryption, etc

- Keys
    - Synchronized Key
    - Asynchronized Kye - others w public keys, I own private key
- Communication
    - simplex
        - Information flows in one direction
    - half-duplex
        - Information can flow in both directions, but not at the same time
    - full-duplex
        - Information can flow in both directions simultaneously
### Section 11 - Control Flow

- Go runtime system
    - Goroutine management
    - Garbage Collection
    - Memory Management
    - Channel Management
        - Communications between goroutines
    - Stack Management
- func init() {} - Run before anything within the go file
    - niladic - no argument needed function
- IF Statement
    
    ```go
    if (condition) { }
    ```
    
    - Comparison Operators
        
        ```go
        ==, !=, <, <=, >, >=
        ```
        
    - Logic Operators
        
        ```go
        &&, ||, !
        ```
        
    - Statement statement Idiom
        
        ```go
        // Variable z is only available within the condition scope
        if z:= f(); z < y {
        	return z
        }
        else {
        	return y
        }
        ```
        
        - ok idiom - try to find out whether the assigning statement status is ok
            
            ```go
            if x, ok := m["Q"]; ok {
            	fmt.Printf("Value for key Q is %d.\n", x)
            } else {
            	fmt.Printf("Value for key Q does not exist.")
            }
            ```
            
- Switch - same as java, but no need of `break` anymore
    - keyword fallthrough - tells the program to consider other switch cases
- Select - For channels and concurrency
    - Concurrency vs parallelism
        - Concurrency - code that is written in a concurrent design pattern - meaning that the code has the potential ability to execute multiple tasks simultaneously, where each task may make progress independently of the others - **using goroutines, lightweight threads of execution that are managed by the Go runtime**
        - Parallelism - the ability of a program to execute multiple tasks simultaneously by utilizing multiple CPUs or cores
    - Example
        
        ```go
        ch1 := make(chan int)
        ch2 := make(chan int)
        
        d1 := time.Duration(rand.Int63n(250))
        d2 := time.Duration(rand.Int63n(250))
        
        go func() {
        	time.Sleep(d1 * time.Millisecond)
        	ch1 <- 41
        }()
        
        go func() {
        	time.Sleep(d2 * time.Millisecond)
        	ch2 <- 41
        }()
        
        select {
        case v1 := <- ch1:
        	fmt.Println("Value from channel 1", v1)
        case v2 := <- ch2:
        	fmt.Println("Value from channel 2", v2)
        }
        ```
        
- For Loops
    - **for init; condition; post {}** - for i := 0; i < 5; i ++
    - **for condition {}**
    - **for {}**
- For range loops - Loops through a data structure
    - Array
        
        ```go
        xi := []int{1, 2, 3, 4, 5}
        // PRINT OUT INDEX AND THE VALUE
        for i, v := range xi {
        	fmt.Println("rangeing over a slice", i, v)
        }
        ```
        
    - Map
        
        ```go
        m := map[string]int {
        	"james":42,
        	"MoneyPenny":32,
        }
        // PRINT OUT KEY VALUE PAIRS
        for k,v := range m {
        	fmt.Println("ranging over a map", k, v)
        }
        ```
        

### Section 12 - Control Flow Exercises - Easy & Skipped