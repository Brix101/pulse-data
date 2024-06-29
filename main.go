package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type MoneyType struct {
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

type SalesByDate struct {
	UnitsOrdered        int       `json:"unitsOrdered"`
	OrderedProductSales MoneyType `json:"orderedProductSales"`
}

type SalesAndTrafficByDate struct {
	Date        string      `json:"date"`
	SalesByDate SalesByDate `json:"salesByDate"`
}

type Data struct {
	SalesAndTrafficByDate []SalesAndTrafficByDate `json:"salesAndTrafficByDate"`
}

func main() {
	fmt.Println("Hello World")
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var data Data
	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

	for i := 0; i < len(data.SalesAndTrafficByDate); i++ {
		fmt.Printf("%+v\n", data.SalesAndTrafficByDate[i])
	}

}
