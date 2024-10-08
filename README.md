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

### Section 13 - Grouping data values - Array & Slice

- Array - a **numbered** sequence of elements of the **same TYPE** - not frequently used due to the limitations of size & Type specification
    - init
        - `x := [3]int{1, 2, 3}`
        - `b := [...]string{"Hello", "World"}`
- Slice - frequently used over Array
    - Built on top of an array
    - Holds VALUES of the **same TYPE**, but **changes in size**
    - init - similar to initialize an array but without specifying the size
        - `xs := []string{"Hello"}`
    - Append
        - `xs = append(xs, "World", "SSSS")`
    - Slicing a slice (Cut parts of the slice) - similar to what it looks like in Python
        
        ```go
        xi := []int{1,2,3,4,5,6,7,8,9}
        // 1ST PARAMETER INCLUSIVE, 2ND PARAMETER EXCLUSIVE
        xi[0:4] -> {1,2,3,4}
        ```
        
    - Delete
        
        ```go
        // Example - Remove 4th item from the list
        xi := []int{1,2,3,4,5,6,7,8,9}
        xi = append(xi[:4], xi[5:]...)
        ```
        
    - Make
        
        ```go
        // an int type slice, nothing stored yet, built default capacity to 10
        xi := make([]int, 0, 10)
        fmt.Println(len(xi)) -> 0
        fmt.Println(cap(xi)) -> 10
        
        xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
        fmt.Println(len(xi)) -> 10
        fmt.Println(cap(xi)) -> 10
        
        xi = append(xi, 11, 12)
        fmt.Println(len(xi)) -> 12
        fmt.Println(cap(xi)) -> 20 (Basically, old slice get copied to a new one with doubled size)
        ```
        
    - Copy - Copy values in slice b to slice a
        - `copy(b, a)`
    - MultiDimension Slice - The size of inner slices can be different
        
        ```go
        	slice := [][]string{}
        	slice = append(slice, []string{"James", "Bond", "Shaken, not stirred"})
        	slice = append(slice, []string{"Miss", "MoneyPenny", "I'm 008", "123123"})
        	fmt.Printf("Slice is %v", slice)
        	-> Slice is [[James Bond Shaken, not stirred] [Miss MoneyPenny I'm 008 123123]]
        ```
        

### Section 14 - Grouping data values - Map

- Specifications
    - key/value storage
    - an **unordered** group of VALUES of one TYPE, called the element type, indexed by a set of unique keys of another type, called the key type
- Init
    
    ```go
    // OPTION 1
    am := map[string]int {
    	"xxx":123
    }
    // OPTION 2 - USE make
    an := make(map[string]int)
    an["xxxx"] = 123
    ```
    
- for range in map
    
    ```go
    for k,v := range(an) {
    	// Code
    }
    
    // GET ONLY VALUES
    for _,v := range(an) {
    	// Code
    }
    
    // GET ONLY KEYS
    for k := range(an) {
    	// Code
    }
    ```
    
- Delete Element
    
    ```go
    // delete is a built-in function
    delete(an, "xxx")
    
    // This still works
    delete(an, "xxx")
    fmt.Println(an["xxx"]) // -> prints "default value" 0
    ```
    
- comma ok idiom - If you look up a non-existent key, the zero value will be returned as the value associated with that non-existent key. **You should use comma ok idiom to check whether the key exists.**
    
    ```go
    v, ok := an["xxx"]
    if ok {
    	fmt.Println("Key doesn't exist.")
    }
    else {
    	fmt.Println(v)
    }
    
    // Shorter Codes
    if v, ok := an["xxx"]; !ok {
    	fmt.Println("Key doesn't exist.")
    }
    else {
    	fmt.Println(v)
    }
    ```
    

### Section 15 - Grouping data values - Struct

- Specifications
    - a data structure
    - a composite/aggregate
    - allows us to **collect values of different types together**
- Example
    
    ```go
    type xxx struct {
    	field1 string
    	field2 int
    }
    
    x1 := xxx {
    	field1: "1st Field"
    	field2: 123
    }
    
    // Embeded Struct
    type xxx2 struct {
    	xxx
    	ltk bool
    }
    
    x2 := xxx2 {
    	xxx : {
    			field1: "1st Field"
    			field2: 123
    	}
    	ltk : true
    }
    ```
    
- Anonymous Structs - For one time use?
    
    ```go
    p1 := struct {
    	first string
    	last string
    	age int
    } {
    	first: "Elton"
    	last: "Li"
    	age:18
    }
    ```
    
- Composition - a.k.a. aggregate data types, aka complex data types
    - a way of structuring and building complex types by combining multiple simpler types
    - In go, you don’t create classes, you create types - This is great cuz this is exactly the format of payload used within APIs
    
    ```go
    
    ```
    

### Section 19 - Functions in the go programming language

- Syntax
    - `func (receiver) identifier(parameters) (returns) {code}`
        - define with **parameters**
        - pass in **arguments**
            
            ```go
            // No Parameters
            func foo() {
            	xxxx
            }
            
            // With Parameters
            func foo1(s string) {
            	xxxx
            }
            
            // With string type return
            func foo2(s string) string {
            	xxx
            	return s
            }
            
            // With multiple parameters & returns
            func foo3(name string, age int) (string, int) {
            	str = ""
            	xx = 222
            	return (str, x)
            }
            ```
            
    - Variadic Parameters - meaning the function takes an **unlimited number of arguments**
        - Example
        
        ```go
        func sum(ii ...int) {
        	fmt.Println(ii)
        	fmt.Printf("%T\n", ii)
        }
        
        func main() {
        		sum(1,2,3,4,5,6,7,8,9) -> [1,2,3,4,5,6,7,8,9]
        													->slice type
        		// CAN ALSO CALL LIKE THIS
        		i := []int{1,2,3,4,5,6,7,8,9}
        		sum(i...)
        		// ... helps map the elements within the array out as function arguments
        }
        ```
        
    - Defer - invokes a function whose execution is deferred to the moment the surrounding function returns, either because
        - the surrounding function executed a return statement,
        - reached the end of its function body,
        - or because the corresponding goroutine is panicking
        
        ```go
        f, error := os.Open(filename)
        if (error != nil) {
        	return "", error
        }
        defer f.Close()  // f.Close will run when the function is returned
        ```
        
- Methods - using receiver
    
    ```go
    type person struct {
    	firstName
    }
    
    func **(p person)** speak() {
    	fmt.Printf("This is %s", p.firstName)
    }
    
    func main() {
    	p1 := person{
    		firstName:"James"
    	}
    	p2 := person{
    		firstName:"Danny"
    	}
    	p1.speak() -> "This is James"
    	p2.speak() -> "This is Danny"
    }
    ```
    
- Interfaces & Polymorphism
    - Interface - defines a set of method signatures
    - Polymorphism - the ability of a VALUE of a certain TYPE to also be of another TYPE
        
        ```go
        type person struct {
        	firstName
        }
        
        type secretAgent struct {
        	person
        	ltk bool
        }
        
        func (p person) speak() {
        	fmt.Printf("This is %s", p.firstName)
        }
        
        func (sa secreatAgent) speak() {
        		fmt.Printf("This is Agent %s", p.firstName)
        }
        
        func saySomething(h human) {
        	h.speak()
        }
        
        // MEANING ANY TYPE HAS THIS METHOD IS CONSIDERED AS THIS TYPE
        type human interface {
        	speak()
        }
        
        func main() {
        	sa1 := secretAgent {
        		p1 := person{
        			firstName:"James"
        		}
        		ltk: true
        	}
        
        	p2 := person{
        		firstName:"Danny"
        	}
        	// sa1.speak() -> "This is James"
        	sa1.saySomething()
        	// p2.speak() -> "This is Danny"
        	p2.saySomething()
        }
        ```
        
    - Example
        
        ```go
        type person struct {
        	first string
        }
        
        // PARAMETER IS A WRITER INTERFACE
        func (p person) writeOut(w io.Writer) {
        	w.Write([]byte(p.first))
        }
        
        func main() {
        	p := person {
        		first: "Jeremy"
        	}
        	// BackGround -> File Write
        	f, err := os.Create("output.txt")
        	if err != nil {
        		log.Fatalf("error %s", err)
        	}
        	defer f.Close()
        	 
        	var b bytes.Buffer
        	p.writeOut(f)
        	p.writeOut(&b)
        	fmt.Println(b.String())
        	
        	// s := []byte("hello world")
        	// _, err = f.Write(s)
        	// if (err != nil) {
        	// 	log.Fatalf("error %s", err)
        	// }
        	
        	// BackGround -> Buffer Write
        	// b := bytes.NewBufferString("Hello ")
        	// fmt.Println(b.String()) -> Hello
        	// ADD GOPHERS INTO THE STRING BUFFER
        	// b.WriteString("Gophers")
        	// fmt.Println(b.String()) -> Hello Gophers
        	
        	// b.Reset()
        	// b.Write([]byte("Happy xxx"))
        	// fmt.Println(b.String()) -> Happy xxx
        }
        
        ```
        
- Anonymous functions
    
    ```go
    func main() {
    	// NO ARGUMENT
    	func(){
    		fmt.Println("Test Anoymous func ran")
    	}()
    	
    	// WITH ARGUMENTS
    	func(x int) {
    		for i:= range(x) {
    			fmt.Println("asdasd")
    		}
    	}(2)
    }
    ```
    
- function type (func expression)
    - Assign a function to a variable
        - execute it
    - an expression is something that evaluates to a value
    - first class citizen - status of certain entities, such as values, types, and functions, that are treated equally and have the same capabilities as other entities in the language
    
    ```go
    x := func(){
    	fmt.Println("Test Anoymous func ran")
    }
    //Execute it
    x() 
    
    y := 	func(x int) {
    	for i:= range(x) {
    		fmt.Println("asdasd")
    	}
    }
    //Execute it
    y(2)
    
    ```
    
    - Return a function
        
        ```go
        func main() {
        	x := foo()
        	value := x()
        }
        
        func foo() int {
        	return 42
        }
        
        func bar() func() int {
        	return func(){return 42}
        }
        
        fmt.Println("%T", foo) -> func() int
        fmt.Println("%T", bar) -> func() func() int
        fmt.Println("%T", y) -> func() int
        ```
        
- Callback function - Passing the function as an argument
    
    ```go
    func main() {
    	x := doMath(1,2,add)
    	fmt.Println(x) -> 3
    	
    	x := doMath(1,2,substract)
    	fmt.Println(x) -> -1
    }
    
    func doMath(a int, b int, f func(int, int) int) int {
    	return f(a, b)
    }
    
    func add(a int, b int) int {
    	return a + b
    }
    
    func substract(a int, b int) int {
    	return a - b
    }
    ```
    
- Closure - a function value that references variables from outside its body
    
    ```go
    func main() {
    	f := incrementor()
    	fmt.Println(f()) // 1
    	fmt.Println(f()) // 2
    	fmt.Println(f()) // 3
    	fmt.Println(f()) // 4
    	fmt.Println(f()) // 5
    	fmt.Println(f()) // 6
    	// WHY ? BECAUSE WHEN RUN THE FUNCTION, ONLY RUN THE INNER PART, X IS OUT OF THE SCOPE
    
    	g := incrementor()
    	fmt.Println(g()) // 1
    	fmt.Println(g()) // 2
    	fmt.Println(g()) // 3
    	fmt.Println(g()) // 4
    	fmt.Println(g()) // 5
    	fmt.Println(g()) // 6
    	// WHY ? BECAUSE WHEN INCREMENTOR() IS CALLED, X IS RE-INITIALIZED
    }
    
    func incrementor() func() int {
    	x := 0
    	return func() int {
    		x ++
    		return x
    	}
    }
    ```
    
- Recursion
    - Base Case
    - Recursion
    
    ```go
    func main() {
    	x := factorial(10)
    }
    
    func factorial(n int) int {
    	if n == 0 {
    		return 1
    	}
    	return n * factorial(n - 1)
    }
    ```
    
- Wrapper Function - Basically, add another layer in between of the original function and your code of using it.

### Section 20 - Tests in Go

- Test Files (Under the same package) - Naming - **filename_test.go**
- Test Functions - Naming - `func TestXxx(*testing.T)`
- Example
    - Original File
    
    ```go
    package main
    func Add(a int, b int) int {
    	return a + b
    }
    ```
    
    - Test File
    
    ```go
    package main
    import "testing"
    func TestAdd(t *testing.T) {
    	total := Add(5,5)
    	if total != 10 {
    		t.Errorf("Sum was incorrect, got %d, want: %d.", total, 10)
    	}
    }
    ```
    
- Run test in Go
    
    ```go
    go test
    ```
    
- Documenting code with comments
    - Within terminal
        - `go doc -cmd -all` shows all comments within the go code
        - `go doc -cmd -u`