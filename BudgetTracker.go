package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	Type   string
	Name   string
	Amount float64
}

var transactions []Transaction

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Budget Tracker ===")
		fmt.Println("1. Add Income")
		fmt.Println("2. Add Expense")
		fmt.Println("3. View Transactions")
		fmt.Println("4. View Balance")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTransaction(reader, "Income")
		case "2":
			addTransaction(reader, "Expense")
		case "3":
			viewTransactions()
		case "4":
			viewBalance()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func addTransaction(reader *bufio.Reader, tType string) {
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter amount: ")
	amountText, _ := reader.ReadString('\n')
	amountText = strings.TrimSpace(amountText)

	amount, err := strconv.ParseFloat(amountText, 64)
	if err != nil {
		fmt.Println("Invalid amount.")
		return
	}

	transaction := Transaction{
		Type:   tType,
		Name:   name,
		Amount: amount,
	}

	transactions = append(transactions, transaction)
	fmt.Println(tType, "added successfully.")
}

func viewTransactions() {
	if len(transactions) == 0 {
		fmt.Println("No transactions yet.")
		return
	}

	fmt.Println("\nTransactions:")
	for i, t := range transactions {
		fmt.Printf("%d. %s - %s: $%.2f\n", i+1, t.Type, t.Name, t.Amount)
	}
}

func viewBalance() {
	var income, expenses float64

	for _, t := range transactions {
		if t.Type == "Income" {
			income += t.Amount
		} else {
			expenses += t.Amount
		}
	}

	balance := income - expenses

	fmt.Printf("\nTotal Income: $%.2f\n", income)
	fmt.Printf("Total Expenses: $%.2f\n", expenses)
	fmt.Printf("Balance: $%.2f\n", balance)
}
