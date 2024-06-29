package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
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

	csvFile, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	header := []string{"date", "value"}
	if err := writer.Write(header); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(data.SalesAndTrafficByDate); i++ {
		fmt.Printf("%+v\n", data.SalesAndTrafficByDate[i])

		record := data.SalesAndTrafficByDate[i]

		parsedDate, err := time.Parse("1/2/2006", record.Date)
		if err != nil {
			log.Fatalf("Error parsing date: %s", err)
		}
		formattedDate := parsedDate.Format("2006-01-02")

		row := []string{formattedDate, strconv.Itoa(record.SalesByDate.UnitsOrdered)}
		if err := writer.Write(row); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("CSV file created successfully")

}
