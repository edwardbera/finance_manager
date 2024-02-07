package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type income struct {
	date        time.Time
	title       string
	description string
	amount      float64
}

func addIncome() *income {
	var myincome = income{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the title :")
	title, err := reader.ReadString('\n')
	myincome.title = title

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the description :")
	description, err2 := reader.ReadString('\n')
	myincome.description = description

	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the amount :")
	amount, err3 := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)
	myincome.amount, err = strconv.ParseFloat(amount, 64)
	if err3 != nil {
		log.Fatal(err)
	}

	myincome.date = time.Now()
	insertData(&myincome)
	return &myincome
}

func addExpense() *income {
	var myincome = income{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Expense :")
	title, err := reader.ReadString('\n')
	myincome.title = title

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the description :")
	description, err2 := reader.ReadString('\n')
	myincome.description = description

	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the amount :")
	amount, err3 := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)
	myincome.amount, err = strconv.ParseFloat(amount, 64)
	if err3 != nil {
		log.Fatal(err)
	}

	myincome.date = time.Now()
	insertData(&myincome)
	return &myincome
}

func createDB() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sts := `
	CREATE TABLE IF NOT EXISTS income(id INTEGER PRIMARY KEY, title TEXT, description TEXT, date TEXT, amount Decimal);
	`
	sts2 := `
	CREATE TABLE IF NOT EXISTS expense(id INTEGER PRIMARY KEY, title TEXT, description TEXT, date TEXT, amount Decimal);
	`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}
	_, err2 := db.Exec(sts2)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Database created.")
}

func insertData(data *income) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}

	sts := "INSERT INTO income(title, description, date, amount) VALUES ('" + data.title + "', '" + data.description + "', '" + data.date.String() + "'," + strconv.FormatFloat(data.amount, 'f', 2, 64) + " )"

	_, err = db.Exec(sts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	createDB()
	//addIncome()
}
