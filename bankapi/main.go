package main

import (
	"bankcore"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*bankcore.Account{}

func main() {
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Balance: 2000,
		Number:  1001,
	}
	accounts[1002] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "Tom",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Balance: 0,
		Number:  1002,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
	}
	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}

}
func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	fromnumberqs := req.URL.Query().Get("fromNumber")
	tonumberqs := req.URL.Query().Get("toNumber")
	amountqs := req.URL.Query().Get("amount")

	if fromnumber, err := strconv.ParseFloat(fromnumberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid from account number!")
	} else if tonumber, err := strconv.ParseFloat(tonumberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid to account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		fromaccount, ok := accounts[fromnumber]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", fromnumber)
		}
		toaccount, ok := accounts[tonumber]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", tonumberqs)
		}

		err := fromaccount.Transfer(toaccount, amount)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			fmt.Fprintf(w, fromaccount.Statement(), toaccount.Statement())
		}
	}
}

type CustomAccount struct {
	*bankcore.Account
}

// Statement ...
func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}
