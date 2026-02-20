package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Expenses struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Id          int    `json:"id"`
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
				fmt.Println("Expense added succesfully")
				database = append(database, Expenses{*addInt, *addDes, len(database) + 1})
			}
		case "list":
			for i := range database {
				fmt.Print(database[i].Amount)
				fmt.Printf(" %v \n", database[i].Description)
			}
		}
	}
	dbjs, _ := json.Marshal(database)
	os.WriteFile("expenses.json", dbjs, 0644)

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
