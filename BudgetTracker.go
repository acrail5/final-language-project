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