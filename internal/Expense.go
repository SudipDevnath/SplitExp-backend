package Expense

type Expense struct {
	Date           string
	ExpenseType    string
	Title          string
	PaidBy         string
	BorrowedOrLent string
	Amount         string
}

func NewExpense(date, expenseType, title, paidBy, borrowedOrLent, amount string) Expense {
	return Expense{Date: date, ExpenseType: expenseType, Title: title, PaidBy: paidBy, BorrowedOrLent: borrowedOrLent, Amount: amount}
}
