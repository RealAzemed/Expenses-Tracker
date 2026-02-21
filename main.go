package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/maniartech/gotime/v2"
)

type Expenses struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	Time        string `json:"time"`
	Month       string `json:"month"`
}

func main() {

	err := checkFile("expenses.json")

	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	file, err := os.ReadFile("expenses.json")

	if err != nil {
		fmt.Println("Error Occurred", err)
	}

	database := []Expenses{}

	json.Unmarshal(file, &database)

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addInt := addCmd.Int("amount", 0, "amount")
	addDes := addCmd.String("description", "", "description")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", 0, "id")

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	summaryMonth := summaryCmd.Int("month", 0, "month")

	if len(os.Args) < 2 {
		fmt.Println("Usage: <command> <args>")
		return
	}

	if len(os.Args) >= 1 {
		switch os.Args[1] {
		case "add":
			addCmd.Parse(os.Args[2:])
			if *addInt == 0 {
				fmt.Println("a amount must be given")
			} else {
				maxId := 0
				for _, d := range database {
					if d.Id > maxId {
						maxId = d.Id
					}
				}
				tempid := maxId + 1
				database = append(database, Expenses{*addInt, *addDes, tempid, gotime.Format(time.Now(), "yyyy-mm-dd"), gotime.Format(time.Now(), "mm")})
				fmt.Printf("Expense added successfully (ID: %v)", tempid)
			}
		case "list":
			fmt.Printf("%-10s %-20s %-25s  %s\n", "ID", "Time", "Amount", "Description")
			for _, d := range database {
				fmt.Printf("%-10v %-20v $%-25v %v\n", d.Id, d.Time, d.Amount, d.Description)
			}
		case "summary":
			if len(os.Args) == 2 {

				var sum int
				for _, d := range database {
					sum += d.Amount
				}
				fmt.Printf("Total Expenses: $%v", sum)
			} else {
				summaryCmd.Parse(os.Args[2:])
				if *summaryMonth <= 0 {
					fmt.Println("Please enter a valid month number")
				} else {
					var sum int
					for _, d := range database {
						month, _ := strconv.Atoi(d.Month)
						if month == *summaryMonth {
							sum += d.Amount
						}

					}
					fmt.Printf("Total Expenses for Selected month %v", sum)
				}
			}
		case "delete":
			deleteCmd.Parse(os.Args[2:])
			if *deleteId == 0 {
				fmt.Println("Please Enter a id number")
			} else {
				for i, d := range database {
					if d.Id == *deleteId {
						database = slices.Delete(database, i, i+1)
						saveData(&database)
						return
					}
				}
				fmt.Printf("Expense not found")
			}
		}
	}
	saveData(&database)

}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveData(database *[]Expenses) {
	dbjs, _ := json.Marshal(database)
	os.WriteFile("expenses.json", dbjs, 0644)
}
