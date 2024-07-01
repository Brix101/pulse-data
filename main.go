package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
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

	ch := make(chan []string, len(data.SalesAndTrafficByDate))

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range ch {
				if err := writer.Write(row); err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	for _, record := range data.SalesAndTrafficByDate {
		ch <- []string{record.Date, fmt.Sprintf("%f", record.SalesByDate.OrderedProductSales.Amount)}
	}

	close(ch)
	wg.Wait()
}
