# Expense Tracker In Go 

A CLI app made in Go to Track your Expenses. This Project is a part of Practice projects for Beginners on [roadmap.sh](https://roadmap.sh/projects/expense-tracker).

## Features

- Add expenses with a description and amount
- Delete expenses by ID
- List all expenses
- Show a summary of total Expenses
- Show summary of total Expenses by Month

  ## Getting Started

### Prerequisites

- Go (version 1.25 or later)

### Installation

1. Clone the repository:

   ```bash
   https://github.com/RealAzemed/Expenses-Tracker

2. Build it:
    ```bash
    go build main.go
    
3. Run:
    ```bash
    ./main

### Usage
You can use the following commands to manage your expenses:

* Add an expense:
```bash
main.go add --description "Lunch" --amount 20
```
* List all expenses:
```bash
main.go list
```
* Show total expenses:
```bash
main.go summary
```
* Show total expenses by month:
```bash
main.go summary --month 8
```
* Delete an expense by ID:
```bash
go run cmd/main.go delete --id 1
```
