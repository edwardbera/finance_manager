package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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
	myincome.amount, err = strconv.ParseFloat(amount, 32)

	if err3 != nil {
		log.Fatal(err)
	}

	return &myincome
}

func main() {

}
