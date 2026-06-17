

Final Project - Written Report
CS 420


Ammon Crail · Peter Kialo
Go Programming Language
Introduction
- History, who created it, why?
Go, also called Golang, is a programming language created at Google in September 2007. It was designed by Robert Griesemer, Rob Pike, and Ken Thompson. It was designed to solve software engineering and scalability challenges that Google had. It was also designed to help programmers to work better with modern computers, especially programs that need concurrency and networking. 
They also created Go out of frustration with the existing languages Google used, which were C++ and Java. They felt that these languages forced developers to choose between execution speed and development efficiency. 

- Uses, what is your language good at.
Go is useful for backend servers, web services, command-line tools, cloud programs, networking software, and infrastructure tools. A lot of Go’s design focuses on making programs simple, fast, and reliable. Go has a built-in standard library, garbage collection, strong typing, and support for concurrency through goroutines. 

- Advantages/Different/Why would you choose this language?
We chose Go because it is easier to read than languages like C or C++, but it still gives good performance because it is compiled. Go also has simple tools, such as go run, go build, and go test, which these tools makes it easy for us to write, run, and test programs. Another advantage is that Go does not try to include too many complicated features because it focuses more on being simple and practical. 
Some advantages:
Simple syntax: Easy to read, it avoids complicated features found in other languages
Compiled language: Its compiled, which helps them run efficiently.
Built-in tools: useful tools like go run, go build, go test, and gofmt.
Strong standard library: packages for input/output, strings, math, networking, web servers, and more.
Good for concurrency: Goroutines and channels, which helps with programs that need to handle many tasks.
Good for command-line programs: Easy to build simple terminal programs.
In our sample program, we made a Budget Tracker. This program stores income and expenses, calculates total income, total expenses, and shows the final balance. This project shows Go features such as structs, methods, slices, loops, conditionals, and formatted output. 
Information about what we learned in the class
- Compiled or Interpreted or both?
Go is mainly a compiled language. Go program is usually compiled into machine code before it runs. This makes Go different from other languages that are mainly interpreted, like Dr. Racket.
 
- What are the primitives?
Go has several basic built-in types. Some Primitives:
bool
string
int
int8, int 16, int 32, int 64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128
byte
rune
Go also has built-in constants like true, false, and nil.
These primitive types are the basic values that Go programs can use. In our Budget Tracker, we will be using string for the transaction names, float64 for the money amounts, and int. 
my program:
type Transaction struct {
    Type   string
    Name   string
    Amount float64
}

- What are the abstraction mechanism?
Go has several abstraction mechanism that allows programmers to organize code and hide details.
Important abstraction:
Functions: A function groups code into a reusable block.
Methods: A Method is a function attached to a type.
Structs: A Struct groups related data together.
Interfaces: An Interface describes the behavior by listing methods that a type must have. 
Packages: A package helps organize code into separate files and libraries.
In our Budget Tracker, we used a Transaction struct and a Budget struct. This allows us to model the program using real ideas as transactions and a budget.
Example:
type Transaction struct {
    Type   string
    Name   string
    Amount float64
}

- What are the means of combination?
Means of combination are ways the language lets programmers combine simple pieces into larger programs.
Some combinations:
Function calls
Structs
Arrays
Slices
Maps
Control structures
Packages
Example: Slice can store many values together.
In our Budget Tracker we used slice of transaction: Transactions []Transaction
Example:
transactions = append(transactions, transaction)
This means the budget can hold many income and expense items. 

- Variable Declarations?
Go has a few ways to declare variables:
Using var:
Var income float64
Using short declaration:
balance := income - expenses
The short declaration := lets Go infer the type automatically. This means Go can figure out the type based on the value. 
In my Budget Tracker we used: var income, expenses float64 which creates two variables, income and expenses, both with the type float64.

- Methods?  How are parameters passed, what are the different options?
Go has methods, but it does not have classes. A method is a function that has a receiver.
Example: 
Func (p Parent) Speak() {
fmt.Println(“Parent Speak:”, p.Name)
}
The (p Parent) part is the receiver. This means the Speak method belongs to the Parent type.
Go passes parameters by value because functions receive a copy of the value. However, Go can also use Pointers which allows the function to change the original value.
In Budget Tracker, we used a pointer: bufio.Reader
Func addTransaction(reader *bufio.Reader, tType string) {
The reader *bufio.Reader parameter is the pointer which lets the function use the same reader object for user input.
Go also has variadic parameters. That means a function can take a variable number of arguments.
Example:
func addNumbers(nums ...int) int {
    return 0
}

- Imperative Features Lecture 13?
Imperative programming means the program gives commands that changes the state of the program.
Go includes:
Variable assignment
Loops
If statements
Switch statements
Mutable Variables
Input and Output
Functions with side effects
Our Budget Tracker Project uses:
income += t.Amount
expenses  += t.Amount
We also used a for loop to keep the program running until the user exits. This program uses a switch statement to choose what to do based on the user's menu choice. 

- Are functions first class?
Functions are First class in Go. This means that the functions can be stored in variables, passed as arguments, and returned from other functions
Example:
func add(a int, b int) int  {
	return a + b
}
operation := add
result := operation(3, 4)
In this example, the function add is stored in the variable operation.
Our Budget Tracker we used regular functions such as:
addTransaction(reader, “Income”)
viewTransaction()
viewBalance()

- Type Checking? Strong/Weak Static/Dyanmic?
Go is statically typed, This means type checking happens before the program runs. If we try to use a value with the wrong type, Go usually catches the error when compiling.
Go is also strongly typed. This means Go does not freely mix compatible types without the programmer converting them.
Example in Budget Tracker:
amount, err := strconv.ParseFloat(amountText, 64)
This is needed because user input starts as text, but the program needs a float64 to do the math.
Go is:
Statically typed
Strongly typed
Go also has type inference, because the programmer can use := and let Go figure out the type.

- Object Oriented - Is it object oriented? does it have Structs/Records
Go has object-oriented features, but it is not object-oriented in the same way as Java or C++. Go does not have classes. Instead, Go uses structs, methods, and interfaces.
Yes, Go has structs. A struct groups field together.
In our Budget Tracker uses this struct:
type Transaction struct {
	Type string
	Name string
	Amount float64
}
This struct is like a record because it stores related data together.

- Single vs. Multiple Inheritance
Go does not have normal class inheritance. This means Go does not have single inheritance or multiple inheritance like some object-oriented languages.
Go uses composition and embedding.
Example:
type Child struct {
	Parent
	Name string
}
This means Child embeds Parent, but it is not the same as traditional inheritance.

- One root object that everything inherits from?
Go does not have one root that all types inherit from. This is different from Java, where all classes inherit from Onject.
In Go, types do not automatically belong to one big class hierarchy.

- Do you have interfaces/protocols?
Go has interfaces. An interface lists methods that a type must have.
Example:
Type Speaker interface {
	Speak()
}
A type does not need to say it implements an interface. If the type has the required methods, then it satisfies the interface automatically. This is important in Go because it supports polymorphism without requiring class inheritance.

- Do you have mix-ins/extensions?
Go does not have min-ins in the same way as other languages. It also does not have extension methods like C#. It uses embedding instead. 
You NEED to create and run sample programs to determine the following properties (see the object oriented lecture).
- Dynamic variable inheritance/Static variable inheritance.  (Include sample program you used to determine this in the appendix)
Go does not have regular class inheritance, I tested this using struct embedding and interfaces.
In my test program, we created a Parent struct and a Child struct. Both had a Name field. The program showed that field lookup is static based on what we wrote in the code.
Example:
child.Name
child.Parent.Name
child.Name gives the child’s name.
child.Parent.Name gives the parent’s name.
This shows that Go does not use dynamic variable inheritance like some object-oriented languages. Field access is chosen by the expression the programmer writes.

- Dynamic method dispatch/Static method dispatch. (Include sample program you used to determine this in the appendix)
For method dispatch, we used an interface:
type Speaker interface {
    Speak()
}
When we pass a Child into a function that expects a Speaker, Go calls the Child version of Speak.
So for Go:
Variable inheritance is static.
Method dispatch can be dynamic when using interfaces.
Go uses composition and interfaces instead of class inheritance.

Sources
Go — How It All Began. A Look Back at the Beginning of Go… | by Bijesh O S | Geek Culture | Medium, medium.com/geekculture/learn-go-part-1-the-beginning-723746f2e8b0. Accessed 17 June 2026.
Rob Pike                  Google. “Go at Google: Language Design in the Service of Software Engineering.” Go, go.dev/talks/2012/splash.article. Accessed 16 June 2026.
There was some AI usage in the Report and Code

