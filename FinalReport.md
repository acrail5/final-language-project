Final Project - Written Report
CS 420
Ammon Crail · Peter Kialo
Language: Go Programming Language
Repository Link: acrail5/final-language-project 

Introduction
History, Who Created It, and Why (Author: Ammon Crail)
Go, also called Golang, is a programming language created at Google. It was designed by Robert Griesemer, Rob Pike, and Ken Thompson. According to the official Go FAQ, they began sketching the goals for Go on September 21, 2007 [1]. Go was created to help solve software engineering problems at Google, especially problems with large code bases, slow builds, and the difficulty of balancing performance with programmer productivity [2].
Go was publicly released as an open-source language in 2009 [1]. The language was designed to be simple, readable, efficient, and useful for modern software development. It also includes built-in support for concurrency, which helps programs handle multiple tasks at the same time [1].
Uses: What Go Is Good At (Author: Ammon Crail)
Go is commonly used for backend servers, web services, command-line tools, cloud software, networking programs, and infrastructure tools [3]. It is useful in these areas because it is compiled, has a strong standard library, and includes tools for building and testing programs [4].
For our project, we used Go to create an interactive Budget Tracker. The program lets the user add income, add expenses, view transactions, and view the final balance. This kind of command-line program is a good fit for Go because Go is simple to run, easy to organize, and good for practical tools.
Advantages: Why We Chose Go (Author: Peter Kialo)
We chose Go because it is simpler to read than languages like C or C++, but it still has good performance because it is compiled. Go also has built-in tools such as go run, go build, go test, and gofmt, which help programmers run, build, test, and format code [4].
Advantage
Explanation
Simple syntax
Go avoids many complicated language features, so the code is usually easy to read.
Compiled language
Go programs are compiled, which helps them run efficiently [5].
Built-in tools
The go command supports building, running, formatting, and testing Go programs [4].
Strong standard library
Go includes packages for input/output, strings, math, networking, web servers, and more [3].
Concurrency support
Go supports goroutines and channels for concurrent programming [1].
Good for command-line programs
Go is practical for tools like our Budget Tracker.


Information About What We Learned in Class
Summary Table (Author Peter Kialo)
Class Topic
Go Answer
Compiled or interpreted?
Go is mainly compiled.
Primitive types
bool, string, integers, unsigned integers, floats, complex numbers, byte, and rune [6].
Abstraction mechanisms
Functions, structs, methods, interfaces, and packages [6].
Means of combination
Function calls, structs, slices, arrays, maps, packages, and control structures [6].
Variable declarations
var, const, and short declaration := [6].
Parameter passing
Go passes parameters by value, but pointers can be used to modify original data [6].
Imperative features
Assignment, loops, conditionals, switch statements, mutable variables, and input/output.
First-class functions
Yes. Functions are values in Go and can be assigned to variables or passed around [6].
Type checking
Statically typed and strongly typed [6].
Object-oriented features
Go has structs, methods, and interfaces, but no classes [1].
Inheritance
Go does not have traditional class inheritance; it uses composition and embedding [7].
Root object
Go does not have one root object that all types inherit from [1].
Interfaces/protocols
Go has interfaces [6].
Mix-ins/extensions
Go does not have traditional mix-ins or extension methods; it uses embedding instead [7].


Compiled or Interpreted? (Author: Ammon Crail)
Go is mainly a compiled language. This means Go source code is translated into machine code before it runs. The Go documentation explains that go build compiles packages and dependencies into an executable program [5].
Even though the command below feels like running a script, Go is still compiling the program first:
go run BudgetTracker.go

Primitives (Author: Ammon Crail)
Go has several built-in basic types. These include booleans, strings, signed integers, unsigned integers, floating-point numbers, complex numbers, byte, and rune [6].
Go Type
Example Use
bool
true or false
string
Names and text
int
Whole numbers
float64
Decimal numbers
byte
Alias for uint8
rune
Alias for int32, often used for characters

In our Budget Tracker, we used string for the transaction type and name, and float64 for the money amount.
type Transaction struct {
    Type   string
    Name   string
    Amount float64
}

Abstraction Mechanisms (Author: Peter Kialo)
Abstraction mechanisms help programmers organize code and hide unnecessary details. Go supports abstraction through functions, structs, methods, interfaces, and packages [6].
Abstraction Mechanism
Meaning in Go
Function
A reusable block of code.
Struct
Groups related data together.
Method
A function attached to a type.
Interface
Describes behavior through method names.
Package
Organizes code into reusable units.

In our Budget Tracker, the Transaction struct is an abstraction because it groups the transaction type, name, and amount into one organized data type.
Means of Combination (Author: Peter Kialo)
Means of combination are ways a language combines smaller pieces into larger programs. Go uses function calls, structs, slices, arrays, maps, packages, loops, if statements, and switch statements [6].
In our program, we used a slice to store many transactions:
var transactions []Transaction

When the user adds an income or expense, the program uses append to add the new transaction:
transactions = append(transactions, transaction)

Variable Declarations (Author: Ammon Crail)
Go has several ways to declare variables. The Go specification includes variable declarations with var, constants with const, and short variable declarations using := [6].
Declaration Style
Example
var declaration
var income float64
Multiple variables
var income, expenses float64
Short declaration
balance := income - expenses
Constant
const taxRate = 0.05

In our Budget Tracker, we used:
var income, expenses float64

We also used short declaration:
balance := income - expenses

Methods and Parameter Passing (Author: Ammon Crail)
Go has methods, but it does not have classes. A method is a function with a receiver [6].
Example from our method dispatch test:
func (p Parent) Speak() {
    fmt.Println("Parent Speak:", p.Name)
}

The (p Parent) part is the receiver. It means the method belongs to the Parent type.
Go passes parameters by value [6]. This means a function receives a copy of the value. However, Go can also use pointers. A pointer allows a function to work with the original data instead of only a copy.
In our Budget Tracker, this function uses a pointer to a bufio.Reader:
func addTransaction(reader *bufio.Reader, tType string) {

The reader *bufio.Reader parameter lets the function use the same reader object for user input.
Go does not support named parameters. Function arguments are passed by position. For example:
addTransaction(reader, "Income")

Go does not allow a call like this:
addTransaction(tType: "Income", reader: reader)

Imperative Features (Author Ammon Crail)
Go has many imperative programming features. Imperative programming means the program gives commands that change state.
Imperative Feature
Example from Our Program
Assignment
balance := income - expenses
Updating values
income += t.Amount
Looping
for { ... }
Conditional logic
if len(transactions) == 0 { ... }
Switch statement
switch choice { ... }
Input/output
fmt.Println, reader.ReadString

Our Budget Tracker uses a for loop to keep the menu running until the user chooses to exit. It also uses a switch statement to decide what action to take based on the user’s choice.
Are Functions First Class? (Author: Peter Kialo)
Yes, functions are first-class in Go. The Go specification treats functions as values with function types [6]. This means functions can be assigned to variables, passed as arguments, and returned from other functions.
Example:
func add(a int, b int) int {
    return a + b
}

operation := add
result := operation(3, 4)

In our Budget Tracker, we mainly used normal functions such as:
addTransaction(reader, "Income")
viewTransactions()
viewBalance()

Type Checking: Strong or Weak? Static or Dynamic?(Author: Peter Kialo)
Go is statically typed. This means type checking happens before the program runs. The Go specification defines Go’s types and rules for assignability, conversions, variables, constants, and expressions [6].
Go is also strongly typed because it does not freely mix incompatible types without requiring the programmer to handle the conversion. For example, our program must convert the user’s amount from a string into a float64 before doing math:
amount, err := strconv.ParseFloat(amountText, 64)

Type Topic
Go Answer
Static or dynamic?
Static type checking
Strong or weak?
Strong typing
Type inference?
Yes, using := in many cases
Example
balance := income - expenses


Object-Oriented Features in Go
Is Go Object-Oriented? (Author: Peter Kialo)
Go has object-oriented features, but it is not object-oriented in the same way as Java or C++. The Go FAQ explains that Go has types and methods and allows an object-oriented style, but it does not have classes or a traditional type hierarchy [1].
Go uses:
Structs
Methods
Interfaces
Composition
Embedding
Structs or Records (Author: Ammon Crail)
Go has structs. A struct groups fields together [6]. This is similar to a record.
Our Budget Tracker uses this struct:
type Transaction struct {
    Type   string
    Name   string
    Amount float64
}

This lets the program store one transaction as a single organized value.
Single vs. Multiple Inheritance (Author: Ammon Crail)
Go does not have traditional class inheritance [1]. This means it does not use single inheritance or multiple inheritance like Java or C++. Instead, Go uses composition and embedding [7].
Example from our test program:
type Child struct {
    Parent
    Name string
}

This means Child embeds Parent, but it is not the same as traditional class inheritance.
One Root Object (Author: Ammon Crail)
Go does not have one root object that everything inherits from. This is different from Java, where all classes inherit from Object. Go does not require all types to belong to one class hierarchy [1].
Interfaces or Protocols (Author: Ammon Crail)
Go has interfaces. An interface lists method requirements [6]. A type satisfies an interface automatically if it has the required methods.
Example from our test program:
type Speaker interface {
    Speak()
}

A type does not need to explicitly say it implements the interface. If it has a Speak() method, then it satisfies Speaker.
Mix-ins or Extensions (Author: Ammon Crail)
Go does not have traditional mix-ins or extension methods. Instead, Go uses embedding. Effective Go explains that embedding lets a type include another type and promote its methods [7]. This can help reuse behavior, but it is still different from class inheritance.

Required Test Programs
Dynamic vs. Static Variable Inheritance (Author: Peter Kialo)
The assignment asked us to create and run a sample program to test variable inheritance. Since Go does not have normal class inheritance, we tested this using struct embedding.
In the test program, both Parent and Child have a field named Name.
fmt.Println("child.Name =", child.Name)
fmt.Println("child.Parent.Name =", child.Parent.Name)

The output shows that field access depends on the expression written by the programmer.
Expression
Meaning
child.Name
Accesses the Name field in Child.
child.Parent.Name
Accesses the Name field inside the embedded Parent.

Conclusion: Go uses static field selection in this test. It does not use dynamic variable inheritance like some object-oriented languages.
Dynamic vs. Static Method Dispatch (Author: Peter Kialo)
We also tested method dispatch using an interface.
type Speaker interface {
    Speak()
}

When a Child value is passed to a function expecting a Speaker, Go calls the Child version of Speak.
func callSpeaker(s Speaker) {
    s.Speak()
}

Conclusion: Go can use dynamic method dispatch through interfaces. The method called depends on the actual concrete type stored in the interface value.
Feature Tested
Result in Go
Variable inheritance
Static field selection
Method dispatch
Dynamic through interfaces
Traditional class inheritance
Not supported
Main design style
Composition and interfaces



Sample Program Explanation (Author: Peter Kialo)
Our main sample program is an interactive Budget Tracker. The program displays a menu with five options:
Add Income
Add Expense
View Transactions
View Balance
Exit
The program uses bufio.NewReader(os.Stdin) to read input from the keyboard. It uses strings.TrimSpace to remove extra spaces and new lines. It uses strconv.ParseFloat to convert the amount from text into a number.
Each income or expense is stored as a Transaction struct. All transactions are stored in the transactions slice.
This program demonstrates:
Structs
Slices
Functions
Loops
Switch statements
If statements
User input
Error checking
String processing
Formatted output
This program is more useful than a “Hello World” program because it allows the user to interact with it and store multiple transactions while the program is running.

Conclusion (Author: Ammon Crail)
Go is a simple, practical, and efficient programming language. It was created at Google to help with real software engineering problems. Go is compiled, statically typed, and strongly typed. It includes useful tools, a strong standard library, structs, functions, methods, interfaces, and concurrency support.
Go is different from many object-oriented languages because it does not have classes or traditional inheritance. Instead, it uses structs, composition, embedding, and interfaces. This makes Go simpler while still allowing programmers to organize code and use polymorphism.
For this project, we created an interactive Budget Tracker program and a separate program to test variable inheritance and method dispatch. The Budget Tracker helped us practice Go’s syntax and features, while the test program helped us understand Go’s object-oriented design choices.

Appendix A: Hello World Program
The first Go program created during the tutorial phase was a simple Hello World program. This was used to make sure Go was installed correctly and that the program could run from the terminal.
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}


Appendix B: Main Sample Program — BudgetTracker.go (Author: Peter Kialo)
The main sample program for this project is an interactive Budget Tracker. The first version allowed the user to add income, add expenses, view transactions, and view the balance. In Phase 2, we extended the program to make it more useful and more complete.
The updated Budget Tracker includes:
Transaction IDs
Income and expense tracking
Categories for transactions
A formatted transaction table
A balance summary table
Spending by category
Delete transaction feature
Monthly expense limit
Progress bar for budget usage
Better input validation
Structs and methods
This makes the program more useful than a basic tutorial program because the user can organize transactions, track spending by category, and manage a monthly expense limit.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	ID       int
	Type     string
	Name     string
	Category string
	Amount   float64
}

type Budget struct {
	Transactions []Transaction
	NextID       int
	ExpenseLimit float64
}

func NewBudget() Budget {
	return Budget{
		Transactions: []Transaction{},
		NextID:       1,
		ExpenseLimit: 0,
	}
}

func (b *Budget) AddTransaction(tType string, name string, category string, amount float64) {
	transaction := Transaction{
		ID:       b.NextID,
		Type:     tType,
		Name:     name,
		Category: category,
		Amount:   amount,
	}

	b.Transactions = append(b.Transactions, transaction)
	b.NextID++
}

func (b *Budget) DeleteTransaction(id int) bool {
	for i, transaction := range b.Transactions {
		if transaction.ID == id {
			b.Transactions = append(b.Transactions[:i], b.Transactions[i+1:]...)
			return true
		}
	}
	return false
}

func (b Budget) TotalIncome() float64 {
	total := 0.0

	for _, transaction := range b.Transactions {
		if transaction.Type == "Income" {
			total += transaction.Amount
		}
	}

	return total
}

func (b Budget) TotalExpenses() float64 {
	total := 0.0

	for _, transaction := range b.Transactions {
		if transaction.Type == "Expense" {
			total += transaction.Amount
		}
	}

	return total
}

func (b Budget) Balance() float64 {
	return b.TotalIncome() - b.TotalExpenses()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	budget := NewBudget()

	for {
		fmt.Println("\n=== Budget Tracker ===")
		fmt.Println("1. Add Income")
		fmt.Println("2. Add Expense")
		fmt.Println("3. View Transactions")
		fmt.Println("4. View Balance Summary")
		fmt.Println("5. View Spending by Category")
		fmt.Println("6. Delete Transaction")
		fmt.Println("7. Set Monthly Expense Limit")
		fmt.Println("8. Exit")
		fmt.Print("Choose option: ")

		choice := readLine(reader)

		switch choice {
		case "1":
			addTransaction(reader, &budget, "Income")
		case "2":
			addTransaction(reader, &budget, "Expense")
		case "3":
			viewTransactions(budget)
		case "4":
			viewBalanceSummary(budget)
		case "5":
			viewSpendingByCategory(budget)
		case "6":
			deleteTransaction(reader, &budget)
		case "7":
			setExpenseLimit(reader, &budget)
		case "8":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose 1-8.")
		}
	}
}

func addTransaction(reader *bufio.Reader, budget *Budget, tType string) {
	fmt.Println("\n--- Add " + tType + " ---")

	fmt.Print("Enter name: ")
	name := readLine(reader)

	fmt.Print("Enter category: ")
	category := readLine(reader)

	amount := readAmount(reader, "Enter amount: ")

	budget.AddTransaction(tType, name, category, amount)
	fmt.Println(tType, "added successfully.")
}

func viewTransactions(budget Budget) {
	if len(budget.Transactions) == 0 {
		fmt.Println("\nNo transactions yet.")
		return
	}

	fmt.Println("\n=== Transactions ===")
	printTransactionHeader()

	for _, transaction := range budget.Transactions {
		printTransactionRow(transaction)
	}

	printTransactionLine()
}

func viewBalanceSummary(budget Budget) {
	income := budget.TotalIncome()
	expenses := budget.TotalExpenses()
	balance := budget.Balance()

	fmt.Println("\n=== Balance Summary ===")
	fmt.Println("+----------------+------------+")
	fmt.Printf("| %-14s | %10s |\n", "Total Income", money(income))
	fmt.Printf("| %-14s | %10s |\n", "Total Expenses", money(expenses))
	fmt.Printf("| %-14s | %10s |\n", "Balance", money(balance))
	fmt.Println("+----------------+------------+")

	if balance >= 0 {
		fmt.Println("Status: You are within your budget.")
	} else {
		fmt.Println("Status: You are over budget.")
	}

	if budget.ExpenseLimit > 0 {
		percentUsed := (expenses / budget.ExpenseLimit) * 100

		fmt.Println("\nMonthly Expense Limit:")
		fmt.Println("Limit:       ", money(budget.ExpenseLimit))
		fmt.Println("Used:        ", money(expenses))
		fmt.Printf("Percent Used: %.1f%%\n", percentUsed)
		fmt.Println("Progress:    ", progressBar(percentUsed))

		if expenses > budget.ExpenseLimit {
			fmt.Println("Warning: You went over your monthly expense limit.")
		}
	}
}

func viewSpendingByCategory(budget Budget) {
	categoryTotals := make(map[string]float64)

	for _, transaction := range budget.Transactions {
		if transaction.Type == "Expense" {
			categoryTotals[transaction.Category] += transaction.Amount
		}
	}

	if len(categoryTotals) == 0 {
		fmt.Println("\nNo expenses to show yet.")
		return
	}

	fmt.Println("\n=== Spending by Category ===")
	fmt.Println("+----------------------+------------+")
	fmt.Printf("| %-20s | %10s |\n", "Category", "Amount")
	fmt.Println("+----------------------+------------+")

	for category, total := range categoryTotals {
		fmt.Printf("| %-20s | %10s |\n", category, money(total))
	}

	fmt.Println("+----------------------+------------+")
}

func deleteTransaction(reader *bufio.Reader, budget *Budget) {
	if len(budget.Transactions) == 0 {
		fmt.Println("\nNo transactions to delete.")
		return
	}

	viewTransactions(*budget)

	fmt.Print("\nEnter transaction ID to delete: ")
	idText := readLine(reader)

	id, err := strconv.Atoi(idText)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	deleted := budget.DeleteTransaction(id)
	if deleted {
		fmt.Println("Transaction deleted successfully.")
	} else {
		fmt.Println("Transaction ID not found.")
	}
}

func setExpenseLimit(reader *bufio.Reader, budget *Budget) {
	fmt.Println("\n--- Set Monthly Expense Limit ---")
	limit := readAmount(reader, "Enter monthly expense limit: ")

	budget.ExpenseLimit = limit
	fmt.Println("Monthly expense limit set to", money(limit))
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readAmount(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		amountText := readLine(reader)

		amount, err := strconv.ParseFloat(amountText, 64)
		if err != nil || amount < 0 {
			fmt.Println("Invalid amount. Please enter a positive number.")
			continue
		}

		return amount
	}
}

func money(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

func printTransactionLine() {
	fmt.Println("+----+----------+----------------------+----------------+------------+")
}

func printTransactionHeader() {
	printTransactionLine()
	fmt.Printf("| %-2s | %-8s | %-20s | %-14s | %10s |\n",
		"ID", "Type", "Name", "Category", "Amount")
	printTransactionLine()
}

func printTransactionRow(transaction Transaction) {
	fmt.Printf("| %-2d | %-8s | %-20s | %-14s | %10s |\n",
		transaction.ID,
		transaction.Type,
		transaction.Name,
		transaction.Category,
		money(transaction.Amount))
}

func progressBar(percent float64) string {
	totalBlocks := 20
	filledBlocks := int((percent / 100) * float64(totalBlocks))

	if filledBlocks > totalBlocks {
		filledBlocks = totalBlocks
	}

	emptyBlocks := totalBlocks - filledBlocks

	return "[" + strings.Repeat("#", filledBlocks) + strings.Repeat("-", emptyBlocks) + "]"
}


Appendix C: Example Run of BudgetTracker.go  (Author: Ammon Crail)
The following is an example of what the program looks like when it runs.
=== Budget Tracker ===
1. Add Income
2. Add Expense
3. View Transactions
4. View Balance Summary
5. View Spending by Category
6. Delete Transaction
7. Set Monthly Expense Limit
8. Exit
Choose option: 1

--- Add Income ---
Enter name: Paycheck
Enter category: Work
Enter amount: 1200
Income added successfully.

=== Budget Tracker ===
1. Add Income
2. Add Expense
3. View Transactions
4. View Balance Summary
5. View Spending by Category
6. Delete Transaction
7. Set Monthly Expense Limit
8. Exit
Choose option: 2

--- Add Expense ---
Enter name: Groceries
Enter category: Food
Enter amount: 150.50
Expense added successfully.

=== Budget Tracker ===
1. Add Income
2. Add Expense
3. View Transactions
4. View Balance Summary
5. View Spending by Category
6. Delete Transaction
7. Set Monthly Expense Limit
8. Exit
Choose option: 3

=== Transactions ===
+----+----------+----------------------+----------------+------------+
| ID | Type     | Name                 | Category       |     Amount |
+----+----------+----------------------+----------------+------------+
| 1  | Income   | Paycheck             | Work           |   $1200.00 |
| 2  | Expense  | Groceries            | Food           |    $150.50 |
+----+----------+----------------------+----------------+------------+

=== Budget Tracker ===
1. Add Income
2. Add Expense
3. View Transactions
4. View Balance Summary
5. View Spending by Category
6. Delete Transaction
7. Set Monthly Expense Limit
8. Exit
Choose option: 4

=== Balance Summary ===
+----------------+------------+
| Total Income   |   $1200.00 |
| Total Expenses |    $150.50 |
| Balance        |   $1049.50 |
+----------------+------------+
Status: You are within your budget.


Appendix D: Inheritance and Method Dispatch Test Program — inheritance_dispatch.go (Author: Peter Kialo)
The project also required a program to test variable inheritance and method dispatch. The file name is inheritance_dispatch.go.
This program is based on the ideas from Lecture 32. Since Go does not have traditional class inheritance, the program uses struct embedding and interfaces to test similar ideas.
package main

import "fmt"

// Lecture 32 Example 3 idea:
// Tests field lookup / variable inheritance using struct embedding.

type C1 struct {
	X int
	Y int
}

func (c *C1) Initialize() {
	// This method does not need to do anything for this test.
}

func (c *C1) SetX1(v int) {
	c.X = v
}

func (c *C1) SetY1(v int) {
	c.Y = v
}

func (c *C1) GetX1() int {
	return c.X
}

func (c *C1) GetY1() int {
	return c.Y
}

type C2 struct {
	C1
	Y int
}

func (c *C2) SetY2(v int) {
	c.Y = v
}

func (c *C2) GetX2() int {
	return c.X
}

func (c *C2) GetY2() int {
	return c.Y
}

func runExample3() {
	o2 := C2{}

	o2.SetX1(101)
	o2.SetY1(102)
	o2.SetY2(999)

	fmt.Println("Lecture 32 Example 3")
	fmt.Println("Variable inheritance / field lookup test:")
	fmt.Printf("Result: (%d %d %d %d)\n",
		o2.GetX1(),
		o2.GetY1(),
		o2.GetX2(),
		o2.GetY2())

	fmt.Println("Explanation:")
	fmt.Println("GetY1 uses the Y field from C1.")
	fmt.Println("GetY2 uses the Y field from C2.")
	fmt.Println("This shows Go field selection is static based on the code written.")
}

// Lecture 32 Example 5 idea:
// Tests method dispatch and a super-style method call.
// Go does not have super, so this program simulates it by explicitly
// calling the parent behavior.

type HasM2 interface {
	M2() int
}

type C1Methods struct{}

func (c C1Methods) M1(self HasM2) int {
	return self.M2()
}

func (c C1Methods) M2() int {
	return 13
}

type C2Methods struct {
	C1Methods
}

func (c C2Methods) M1() int {
	return 22
}

func (c C2Methods) M2() int {
	return 23
}

func (c C2Methods) SuperM1(self HasM2) int {
	return c.C1Methods.M1(self)
}

type C3Methods struct {
	C2Methods
}

func (c C3Methods) M1() int {
	return 32
}

func (c C3Methods) M2() int {
	return 33
}

func (c C3Methods) M3() int {
	// This simulates: super m1()
	// It starts with the parent version of m1,
	// but self is still the C3 object.
	return c.C2Methods.SuperM1(c)
}

func runExample5() {
	o3 := C3Methods{}

	fmt.Println("\nLecture 32 Example 5")
	fmt.Println("Method dispatch / super-style call test:")
	fmt.Println("Result:", o3.M3())

	fmt.Println("Explanation:")
	fmt.Println("M3 calls the parent version of M1.")
	fmt.Println("That parent M1 calls self.M2().")
	fmt.Println("Since self is still the C3 object, C3's M2 runs.")
	fmt.Println("Therefore, the result is 33.")
}

func main() {
	runExample3()
	fmt.Println()
	runExample5()
}


Appendix E: Expected Output of inheritance_dispatch.go (Author: Ammon Crail)
Lecture 32 Example 3
Variable inheritance / field lookup test:
Result: (101 102 101 999)
Explanation:
GetY1 uses the Y field from C1.
GetY2 uses the Y field from C2.
This shows Go field selection is static based on the code written.

Lecture 32 Example 5
Method dispatch / super-style call test:
Result: 33
Explanation:
M3 calls the parent version of M1.
That parent M1 calls self.M2().
Since self is still the C3 object, C3's M2 runs.
Therefore, the result is 33.


Appendix F: Results Summary (Author: Peter Kialo)
Program
File Name
Purpose
Hello World
shown in Appendix A
Verified Go installation
Budget Tracker
BudgetTracker.go
Main sample program
Inheritance and Dispatch Test
inheritance_dispatch.go
Tests Lecture 32 variable inheritance and method dispatch


Feature Tested
Result in Go
Variable inheritance / field lookup
Static field selection
Method dispatch
Dynamic behavior can be shown through interfaces
Traditional class inheritance
Not supported in Go
Super method call
Not built into Go; simulated by explicitly calling embedded parent behavior


AI Use Disclosure (Author: Ammon Crail)
AI was used to help organize the report, improve wording, suggest source locations, and improve grammar. AI was also used to help format the appendix and explain the results of the method dispatch test.
AI was used to help refine the Budget Tracker program, including formatting, input handling, transaction viewing, balance calculation, and readability.
All AI-generated suggestions were reviewed, tested, modified when necessary, and approved before being included in the final submission.

References (Author: Ammon Crail)
[1] Go Authors. “Frequently Asked Questions.” Official Go Documentation.
https://go.dev/doc/faq
[2] Pike, Rob. “Go at Google: Language Design in the Service of Software Engineering.” Official Go Talks.
https://go.dev/talks/2012/splash.article
[3] Go Authors. “The Go Programming Language.” Official Go Website.
https://go.dev/
[4] Go Authors. “Command Documentation.” Official Go Documentation.
https://go.dev/doc/cmd
[5] Go Authors. “Compile and install the application.” Official Go Tutorial.
https://go.dev/doc/tutorial/compile-install
[6] Go Authors. “The Go Programming Language Specification.” Official Go Documentation.
https://go.dev/ref/spec
[7] Go Authors. “Effective Go.” Official Go Documentation.
https://go.dev/doc/effective_go


