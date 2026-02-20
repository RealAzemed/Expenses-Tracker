package main

import (
	"flag"
	"fmt"
	"os"
)

type Expenses struct {
	amount      int
	description string
}

func main() {
	database := []Expenses{}

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
			fmt.Println("amount: ", *addInt)
			if *addInt == 0 {
				fmt.Println("a amount must be given")
			} else {
				fmt.Println("\ndescription :", *addDes)
				database = append(database, Expenses{*addInt, *addDes})
			}

			fmt.Print(database)
		}
	}
}
